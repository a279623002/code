# Python 面试必问知识点

> 项目背景：边端智能助手（langchain==0.2.16 + faiss），python 是主力语言。异步、并发、生成器、装饰器等高频考点几乎都直接对应项目里某段代码。

## 一、异步编程（asyncio）

### 1. 事件循环（Event Loop）核心原理

**答**：
事件循环是 asyncio 的调度中枢，本质是一个**单线程死循环**，不断从"就绪队列"里取出任务执行。一个协程遇到 `await` 时会**挂起**（保存栈帧到 task 对象），把控制权交回事件循环；事件循环继续调度其他就绪协程；当 `await` 的 Future 完成时，回调把对应 task 重新放回就绪队列，等待下一轮被调度。

关键点：
- **单线程**：asyncio 自身是单线程的，所谓的"并发"是**协作式多任务**（靠 await 主动让出），不是抢占式
- **协程切换不涉及线程切换**：上下文就是 `PyFrameObject`，切换成本极低（us 级），远低于线程切换（ms 级）
- **事件循环与主线程绑定**：`asyncio.run()` 会自动创建/绑定/关闭循环；不要在已有循环的线程里再 `asyncio.run()`

```python
import asyncio

async def fetch(name, delay):
    await asyncio.sleep(delay)  # 模拟 IO，让出控制权
    return f"{name} done"

async def main():
    # 创建任务后立即塞入就绪队列，事件循环调度
    t1 = asyncio.create_task(fetch("a", 1))
    t2 = asyncio.create_task(fetch("b", 2))
    # await 触发事件循环"转起来"
    r1 = await t1
    r2 = await t2
    print(r1, r2)

asyncio.run(main())
```

### 2. gather 并发执行

**答**：
`asyncio.gather(*aws, return_exceptions=False)` 把多个协程/任务**并发**调度，**一起等所有结果**。底层就是循环里 `await` 每个 task，遇到 IO 就让出，整体执行时间 = 最慢的那个任务。

> ⚠️ **gather 不等于并行**：在单线程事件循环里只是"交替执行 IO 等待"，没有多核加速。真正的并行要靠 `loop.run_in_executor` 扔到线程/进程池。

```python
async def main():
    # 并发 3 个请求，总耗时 ≈ 2s（不是 6s）
    results = await asyncio.gather(
        fetch("a", 2),
        fetch("b", 2),
        fetch("c", 2),
    )
    return results

# return_exceptions=True 让一个失败不影响其他
async def safe():
    results = await asyncio.gather(
        may_fail(), may_ok(), may_fail2(),
        return_exceptions=True
    )
    for r in results:
        if isinstance(r, Exception):
            print("skip", r)
```

**gather vs wait vs as_completed**：
| 场景 | 用法 |
|---|---|
| 全部结果要齐 | `gather` |
| 部分完成即可 | `asyncio.wait(tasks, return_when=FIRST_COMPLETED)` |
| 流式拿到结果顺序 | `asyncio.as_completed(tasks)` |

### 3. async/await 关键字

**答**：
- `async def` 定义**协程函数**，调用返回**协程对象**（不会自动执行！）
- `await` 表达式：只能在 `async def` 内使用；后面跟 awaitable（协程/Future/Task）
- **`await` 的本质是 `yield from` 的语法糖**（CPython 3.5+ 源码可见）

```python
async def foo():
    return 1  # 普通协程

# ❌ 错：调用没 await，得到的是 coroutine 对象，不会执行
result = foo()  # <coroutine object foo at 0x...>

# ✅ 对
result = await foo()  # 1
```

### 4. 项目里的异步应用

| 场景 | 实现 |
|---|---|
| 知识库切片（CPU 密集） | `await loop.run_in_executor(None, blocking_splitter)` |
| embedding 批量调用（IO 密集 + 限流） | `asyncio.Semaphore(8)` + `gather` |
| LLM 流式输出 | `async for chunk in llm.astream(prompt): yield chunk` |
| 会议纪要实时转写 | 异步生成器 + `async for evt in asr_stream(): ...` |

---

## 二、多线程（threading）

### 1. GIL（全局解释器锁）— 必问

**答**：
GIL 是 CPython 解释器级别的互斥锁，**保证同一时刻只有一个线程执行 Python 字节码**。目的：保护 CPython 内存管理（引用计数）不被多线程破坏。

**结论**：
- CPU 密集型任务 → 多线程**无效**（甚至更慢，因为 GIL 争抢）
- IO 密集型任务 → 多线程**有效**（IO 等待时 GIL 释放，线程切换到其他线程）
- 解决 CPU 密集 → 用**多进程**或 C 扩展（numpy 释放 GIL）

```python
# 验证 GIL：两个 CPU 密集线程，总耗时 ≈ 串行的 2 倍
import threading, time

def cpu():
    s = 0
    for _ in range(10**7):
        s += 1

t = time.time()
t1 = threading.Thread(target=cpu); t1.start()
t2 = threading.Thread(target=cpu); t2.start()
t1.join(); t2.join()
print(time.time() - t)  # 2.x 秒（没有并行加速）
```

### 2. threading vs ThreadPoolExecutor

**答**：
- `threading.Thread` 手动管理线程
- `concurrent.futures.ThreadPoolExecutor` 线程池，**自动复用**线程，`submit` 返回 Future，`map` 批量提交

```python
from concurrent.futures import ThreadPoolExecutor

def io_task(url):
    return requests.get(url).text

with ThreadPoolExecutor(max_workers=10) as pool:
    results = list(pool.map(io_task, urls))
    # 或：
    futs = [pool.submit(io_task, u) for u in urls]
    for f in futs:
        print(f.result())
```

### 3. 线程同步原语

| 原语 | 作用 |
|---|---|
| `Lock` | 互斥锁，最常用 |
| `RLock` | 可重入锁，同一线程可多次 acquire |
| `Semaphore` | 信号量，限制并发数（如数据库连接池） |
| `Event` | 事件标志，wait/clear/set 跨线程通知 |
| `Condition` | 条件变量，配合 Lock 实现生产者-消费者 |
| `Queue` | 线程安全队列，自带锁 |

```python
import threading
total = 0
lock = threading.Lock()

def add():
    global total
    with lock:  # 等价于 acquire/release
        total += 1  # 这步不是原子的，必须加锁
```

### 4. daemon 线程

**答**：
`thread.setDaemon(True)` 设为守护线程。主线程退出时，守护线程**立即被杀死**（不执行 finally）。日志/心跳/监控线程适合设为守护。

> Python 3.10+ 改用 `thread.daemon = True` 或 `threading.Thread(daemon=True)`。

---

## 三、多进程（multiprocessing）

### 1. 适用场景

**答**：
- **CPU 密集**任务（计算、图像处理、模型推理）：多进程绕开 GIL，真正多核并行
- **进程隔离**：一个进程崩溃不影响其他（适合 worker 池容错）

### 2. Process vs Pool

```python
from multiprocessing import Process, Pool

def work(x):
    return x * x

# 单进程
p = Process(target=work, args=(5,))
p.start(); p.join()

# 进程池：自动负载均衡
with Pool(processes=4) as pool:
    results = pool.map(work, range(10))  # [0, 1, 4, 9, ...]
    # 异步版本
    futs = [pool.apply_async(work, (i,)) for i in range(10)]
    results = [f.get() for f in futs]
```

### 3. 进程间通信（IPC）

| 方式 | 特点 |
|---|---|
| `Queue` | 线程/进程安全，跨进程自动 pickle |
| `Pipe` | 双工，速度快，但只两个端点 |
| `Value/Array` | 共享内存，最快但要自己加锁 |
| `Manager` | 代理对象，支持 list/dict 等复杂结构 |
| `multiprocessing.shared_memory` | Python 3.8+，显式共享内存 |

### 4. 进程 vs 线程 vs 协程 — 选择决策树

```
任务类型？
├─ CPU 密集（计算/模型推理）→ 多进程
├─ IO 密集（网络/磁盘）
│   ├─ 高并发（万级连接）→ asyncio（单线程万级协程）
│   └─ 低并发（几十个）→ 多线程（代码简单）
└─ CPU + IO 混合 → 进程池 + 协程（外层进程隔离，内层 asyncio）
```

---

## 四、生成器与迭代器（核心）

### 1. 三者区别（必问）

| 概念 | 定义 | 关键方法 |
|---|---|---|
| **可迭代对象** (Iterable) | 可用 `for` 循环遍历的对象 | `__iter__` 返回迭代器 |
| **迭代器** (Iterator) | 记住遍历位置的对象 | `__iter__` + `__next__` |
| **生成器** (Generator) | 用 `yield` 简化迭代器写法 | 自动实现 `__iter__` 和 `__next__` |

```python
from collections.abc import Iterable, Iterator

# 可迭代但不是迭代器
l = [1, 2, 3]
isinstance(l, Iterable)  # True
isinstance(l, Iterator)  # False（list 不是迭代器）

# 列表的迭代器
it = iter(l)
next(it)  # 1
next(it)  # 2
next(it)  # StopIteration
```

### 2. 生成器函数

**答**：
`yield` 是生成器函数的关键字。**调用生成器函数不会执行函数体**，而是返回一个生成器对象；每次 `next()` 才执行到下一个 `yield` 暂停。

```python
def gen():
    print("start")
    yield 1
    print("after 1")
    yield 2
    print("end")

g = gen()       # 不打印
next(g)         # start  → 1
next(g)         # after 1 → 2
next(g)         # end → StopIteration
```

**生成器的四个状态**（CPython 内部）：
1. `GEN_CREATED` — 刚创建，未启动
2. `GEN_RUNNING` — 正在执行（解释器内部）
3. `GEN_SUSPENDED` — `yield` 处暂停
4. `GEN_CLOSED` — 执行完毕或 `close()`

### 3. yield 表达式（send/throw/close）

```python
def echo():
    while True:
        val = yield  # 接收 send 进来的值
        print("got:", val)

g = echo()
next(g)              # 启动到第一个 yield
g.send("hello")      # got: hello
g.send("world")      # got: world
g.throw(ValueError)  # 抛异常进生成器（可用于重置）
g.close()            # 关闭，GeneratorExit 抛到生成器内
```

### 4. yield from — 委托生成器

**答**：
`yield from` 把内层生成器**委托**给外层。**3 个作用**：
1. 简化嵌套 `for yield` 代码
2. **自动转发** `send`/`throw`/`close` 到子生成器
3. **返回**子生成器的 return 值（其他写法无法直接拿）

```python
def inner():
    yield 1
    yield 2
    return "done"  # 返回值给 yield from 表达式

def outer():
    result = yield from inner()  # result = "done"
    print(result)

list(outer())  # [1, 2]，然后打印 done
```

### 5. 项目应用：流式输出

```python
async def stream_llm(prompt):
    # langchain 的 astream 是异步生成器
    async for chunk in llm.astream(prompt):
        yield chunk.content  # 每个 token 立即推给前端

# FastAPI 路由
@app.get("/chat/stream")
async def chat(prompt: str):
    return StreamingResponse(stream_llm(prompt), media_type="text/event-stream")
```

---

## 五、装饰器（Decorator）

### 1. 基础

**答**：
装饰器本质是一个**接受函数返回函数的可调用对象**，用来在不修改原函数代码的前提下增加功能（如日志/计时/缓存/权限校验）。

```python
def logger(func):
    def wrapper(*args, **kwargs):
        print(f"call {func.__name__}")
        result = func(*args, **kwargs)
        print(f"end {func.__name__}")
        return result
    return wrapper

@logger
def add(a, b):
    return a + b

add(1, 2)
# call add
# end add
```

### 2. functools.wraps — 保留元信息

```python
from functools import wraps

def logger(func):
    @wraps(func)  # 把 __name__/__doc__ 复制给 wrapper
    def wrapper(*args, **kwargs):
        ...
    return wrapper

# 没有 @wraps，add.__name__ 会变成 'wrapper'，debug/序列化会出问题
```

### 3. 带参装饰器

```python
def retry(times=3):
    def decorator(func):
        @wraps(func)
        def wrapper(*args, **kwargs):
            for i in range(times):
                try:
                    return func(*args, **kwargs)
                except Exception as e:
                    if i == times - 1: raise
        return wrapper
    return decorator

@retry(times=5)
def fetch():
    ...
```

### 4. 类装饰器

```python
class Count:
    def __init__(self, func):
        self.func = func
        self.cnt = 0
    def __call__(self, *args, **kwargs):
        self.cnt += 1
        print(f"call {self.cnt}")
        return self.func(*args, **kwargs)

@Count
def hello(): print("hi")
```

### 5. 项目应用

```python
# 1. embedding 结果缓存（减少 token 费）
from functools import lru_cache

@lru_cache(maxsize=2048)
def embed_cached(text: str):
    return embed_model.embed_query(text)

# 2. 重试 + 限流（llm/asr 经常超时）
@retry(times=3)
@rate_limit(calls=10, period=1)  # 装饰器链
def call_llm(prompt): ...
```

---

## 六、Lambda

**答**：
`lambda args: expr` 是**匿名函数**的语法糖，只能写**单表达式**，自动 return 该表达式结果。

| 场景 | 用法 |
|---|---|
| 简单排序 key | `sorted(data, key=lambda x: x["age"])` |
| map/filter 配合 | `list(map(lambda x: x*2, [1,2,3]))` → `[2,4,6]` |
| 回调函数 | `button.clicked.connect(lambda: print("ok"))` |

**限制**：
- 不能写 `if/else` 语句（但三元表达式可以）
- 不能写 `for/while`
- 没有 `__name__`（`functools.wraps` 也救不了）

**建议**：复杂逻辑别用 lambda，用 `def`（可读性 + 调试栈友好）。

---

## 七、map / filter / reduce

### 1. map

**答**：
`map(func, *iterables)` 把 func 应用到每个元素，**惰性**返回一个迭代器（Python 3）。

```python
list(map(str, [1, 2, 3]))           # ['1','2','3']
list(map(lambda x: x*2, range(3)))  # [0, 2, 4]
list(map(pow, [2,3,4], [3,2,1]))    # [8, 9, 4]（多序列，按最短截断）
```

**对比列表推导式**：
- 列表推导式更快（CPython 优化），map 更省内存（惰性）
- 简单逻辑用推导式，复杂函数用 map

### 2. filter

```python
list(filter(lambda x: x > 0, [-1, 0, 1, 2]))  # [1, 2]
# Python 3 中 filter 返回迭代器，Python 2 返回 list
# 等价的推导式：[x for x in lst if x > 0]
```

### 3. reduce

**答**：
`functools.reduce(func, iterable, initializer=None)` 累计应用 func。**面试考 reduce = 考手写累加**。

```python
from functools import reduce
reduce(lambda acc, x: acc + x, [1, 2, 3, 4])           # 10
reduce(lambda acc, x: acc + x, [1, 2, 3, 4], 100)      # 110（带初值）
# 累乘
reduce(lambda a, b: a * b, [1, 2, 3, 4])                # 24
```

**实现原理**（面试加分项）：
```python
def myreduce(func, seq, init=None):
    it = iter(seq)
    if init is None:
        acc = next(it)
    else:
        acc = init
    for x in it:
        acc = func(acc, x)
    return acc
```

### 4. map/filter/reduce 对比

| 函数 | 作用 | 返回 |
|---|---|---|
| `map` | 一对一变换 | 同样长度的迭代器 |
| `filter` | 过滤 | 长度 ≤ 原来的迭代器 |
| `reduce` | 累计 | 单个值 |

---

## 八、综合面试题

### Q1：GIL 是什么？多线程有用吗？怎么绕开？

**答**：
GIL 保证同一时刻只有一个线程执行 Python 字节码。
- IO 密集 → 多线程有效（IO 等待释放 GIL）
- CPU 密集 → 多线程**无效**，用**多进程**绕开；或用 C 扩展（numpy/pandas 内部释放 GIL）
- 高并发 IO → 用 **asyncio**（单线程万级协程，切换成本 < 1μs）

### Q2：asyncio 和多线程的区别？

**答**：
- asyncio：**单线程**协作式，靠 await 让出；适合高并发 IO 密集（万级连接）
- 多线程：**多线程**抢占式，靠 OS 调度；适合需要并行阻塞 IO 的场景（几十个）
- asyncio 切换成本是线程的 1/1000，但**不能**做阻塞调用（会卡住整个循环）
- 项目里常用：asyncio + `run_in_executor(线程池)` 混合模式处理"阻塞 + 高并发"

### Q3：生成器和迭代器的区别？自己写一个迭代器？

```python
class MyRange:
    def __init__(self, n): self.n = n; self.i = 0
    def __iter__(self): return self
    def __next__(self):
        if self.i >= self.n: raise StopIteration
        self.i += 1
        return self.i - 1
```

### Q4：装饰器链的执行顺序？

**答**：
```python
@A
@B
@C
def f(): pass
# 等价于 f = A(B(C(f)))
# 调用时：f() → A.wrapper → B.wrapper → C.wrapper → f → 反向返回
```

### Q5：yield 和 return 的区别？

- `return` 终止函数并返回值
- `yield` 暂停函数并返回值，下次从暂停处继续
- 生成器可以有多个 yield；普通函数只能一个 return
- yield 让函数变成"惰性的"，适合大文件流式读取、LLM 流式输出

### Q6：Python 里实现单例的 5 种方式？

1. 模块级变量（最简单）
2. `__new__` 重写
3. 装饰器 + dict
4. metaclass `type` 子类
5. `functools.lru_cache` 装饰器（最 Pythonic）

### Q7：深拷贝和浅拷贝？

- 浅拷贝：拷贝父对象，子对象**共享引用**（`copy.copy`、列表的 `[:]`、`list()`）
- 深拷贝：完全独立副本（`copy.deepcopy`，递归）
- **坑**：嵌套可变对象修改会影响浅拷贝源

### Q8：*args 和 **kwargs？

- `*args` 收集多余**位置参数**为元组
- `**kwargs` 收集多余**关键字参数**为字典
- 调用时：`*lst` 解包列表为位置参数；`**d` 解包字典为关键字参数

### Q9：Python 内存管理？

- 引用计数（主）+ 标记清除（解决循环引用）+ 分代回收（gc 模块）
- 引用计数为 0 立即释放
- `sys.getrefcount(obj)` 看引用次数
- 调试内存泄漏：`objgraph.show_backrefs(obj)` 或 `tracemalloc`

### Q10：项目里 Python 性能优化点？

| 场景 | 优化 |
|---|---|
| faiss embedding 慢 | `embeddings.embed_documents` 批量 + asyncio.Semaphore 限流 |
| markitdown 切片 CPU 密集 | `run_in_executor(ProcessPoolExecutor)` 进程池并行切片 |
| 重复 query 重复 embedding | `functools.lru_cache(maxsize=2048)` 缓存 |
| 大文件加载 | 异步生成器 + 分块 yield，避免一次性读入内存 |
| 频繁 JSON 序列化 | `orjson` 替代 `json`（3-5x 快） |
| token 计数 | `tiktoken` 编码缓存，避免重复 encode |

---

## 九、一句话总结（面试收尾用）

- **异步**：单线程事件循环 + 协作式调度 + await 让出，适合高并发 IO
- **多线程**：受 GIL 限制只对 IO 密集有效，线程安全靠 Lock/Queue
- **多进程**：绕开 GIL 适合 CPU 密集，进程隔离强但通信成本高
- **生成器**：yield 实现的惰性迭代器，省内存 + 支持流式
- **装饰器**：AOP 思想的应用，叠加功能不改源码
- **lambda/map/filter/reduce**：函数式编程三件套，能用但列表推导式通常更快更清晰

> 选型口诀：**IO 密集看并发（asyncio > thread），CPU 密集看并行（process > thread）**
