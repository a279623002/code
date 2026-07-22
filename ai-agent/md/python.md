# Python 面试必问知识点

> 项目背景：边端智能助手（langchain==0.2.16 + faiss），python 是主力语言。异步、并发、生成器、装饰器等高频考点几乎都直接对应项目里某段代码。

---

## 一、异步编程（asyncio）

### 1. 事件循环（Event Loop）核心原理

**一句话**：单线程里的"任务调度员"，哪个协程在 `await` 等 IO，就去执行另一个协程；IO 完成后再回来继续。

**生活例子**：饭店只有 1 个服务员（单线程）。客人 A 点菜后去等后厨（await），服务员立刻去招待客人 B，而不是傻等 A。

```python
import asyncio

async def cook(name, sec):
    print(f"[{name}] 开始等后厨")
    await asyncio.sleep(sec)   # 模拟等 IO（等后厨），主动让出
    print(f"[{name}] 菜好了")
    return f"{name} 的菜"

async def main():
    # create_task：告诉调度员"这两件事同时盯着"
    t1 = asyncio.create_task(cook("番茄炒蛋", 1))
    t2 = asyncio.create_task(cook("红烧肉", 2))
    # await 等结果
    r1 = await t1
    r2 = await t2
    print(r1, r2)

asyncio.run(main())
```

**运行结果**：
```
[番茄炒蛋] 开始等后厨
[红烧肉] 开始等后厨
[番茄炒蛋] 菜好了
[红烧肉] 菜好了
番茄炒蛋 的菜 红烧肉 的菜
```
> 总耗时 ≈ 2s（不是 1+2=3s），因为两道菜是"同时等"的。

**面试考点**：
- 协程切换不创建线程，只是保存/恢复一个栈帧，成本极低
- 不要在 asyncio 里放阻塞操作（如 `time.sleep`、`requests.get`），会卡住整个循环
- **time.sleep**：
    - 同步阻塞休眠，死死占用当前线程，不会主动让出事件循环。
    - 执行它时，整个事件循环被冻结，所有其他协程全部暂停排队
    - 三个任务用 gather + time.sleep (2)，总耗时 = 6s，退化成串行。
    - await asyncio.sleep(2)：只是向事件循环注册一个定时事件，立刻挂起协程、释放线程，循环可以调度别的任务，等待时间重叠。
- **requests.get**：
    - requests 是纯同步网络库，底层是阻塞式 socket
    - 发起 HTTP 请求 → 线程卡死等待服务器返回数据，全程不会释放线程。
    - 可使用aiohttp（asyncio 生态标配异步网络库），全程非阻塞、配合 await，gather 完美并发。
    - 老旧业务只能用 requests，不能改 aiohttp，用 asyncio.to_thread() 将同步阻塞函数放到系统线程运行，主线程事件循环不会被阻塞。


### 2. gather 并发执行

**一句话**：把多个任务"打包"，一起启动、一起等，总耗时 = 最慢的那个。

```python
async def main():
    # 同时发起 3 个 IO 请求
    results = await asyncio.gather(
        cook("A", 1),
        cook("B", 2),
        cook("C", 3),
    )
    print(results)

asyncio.run(main())
```

**运行结果**：
```
[A] 开始等后厨
[B] 开始等后厨
[C] 开始等后厨
[A] 菜好了
[B] 菜好了
[C] 菜好了
['A 的菜', 'B 的菜', 'C 的菜']
```

**处理单个失败**：
```python
async def may_fail():
    raise ValueError("出错了")

async def ok():
    return "ok"

async def main():
    results = await asyncio.gather(
        may_fail(), ok(),
        return_exceptions=True  # 失败返回异常对象，不炸整体
    )
    for r in results:
        if isinstance(r, Exception):
            print("失败:", r)
        else:
            print("成功:", r)

asyncio.run(main())
```

**运行结果**：
```
失败: 出错了
成功: ok
```

### 3. async/await 关键字

- `async def`：定义协程函数
- `await`：等一个协程完成，并**让出 CPU**
- 调用协程函数不会执行，只会返回一个"协程对象"

```python
async def hello():
    return "hello"

# ❌ 错：没 await，得到的是协程对象，不会执行
cor = hello()
print(cor)  # <coroutine object hello at 0x...>

# ✅ 对
result = await hello()
print(result)  # hello
```

### 4. 项目里的异步应用

| 场景 | 为什么用异步 | 代码形态 |
|---|---|---|
| embedding 批量调用 | 调用 embedding API 是 IO 等待 | `asyncio.Semaphore(8)` + `gather` |
| LLM 流式输出 | 等模型返回每个 token | `async for chunk in llm.astream(prompt): yield chunk` |
| 会议纪要实时转写 | 等 ASR 流式返回 | `async for evt in asr_stream(): ...` |
| markitdown 切片（CPU 密集） | 不能阻塞事件循环 | `await loop.run_in_executor(pool, split)` |

---

## 二、多线程（threading）

### 1. GIL（全局解释器锁）— 必问

**一句话**：CPython 里任何时刻只有**一个线程**在执行 Python 字节码。

**为什么要有**：保护引用计数，防止多线程同时改对象内存导致崩溃。

**结论**：
- IO 密集型（网络请求、读写文件）→ 多线程**有效**
- CPU 密集型（大量计算）→ 多线程**无效**，甚至更慢
- CPU 密集想真正并行 → 用**多进程**

**验证代码**：
```python
import threading, time

def cpu():
    s = 0
    for _ in range(10_000_000):
        s += 1
    return s

start = time.time()
t1 = threading.Thread(target=cpu)
t2 = threading.Thread(target=cpu)
t1.start(); t2.start()
t1.join();  t2.join()
print(time.time() - start)
```

**运行结果**：约 **2.x 秒**（两个 CPU 密集线程，并没有 2 倍加速，因为 GIL 在抢）

**对比多进程**（下面会讲）：
```python
from multiprocessing import Process
p1 = Process(target=cpu)
p2 = Process(target=cpu)
# 两个进程各跑一个核，总耗时 ≈ 1.x 秒
```

### 2. threading vs ThreadPoolExecutor

```python
import requests
from concurrent.futures import ThreadPoolExecutor

urls = ["https://example.com/1", "https://example.com/2", "https://example.com/3"]

def fetch(url):
    return requests.get(url, timeout=5).status_code

# 线程池：10 个线程复用
with ThreadPoolExecutor(max_workers=10) as pool:
    results = list(pool.map(fetch, urls))
    print(results)  # [200, 200, 200]
```

### 3. 线程同步原语

```python
import threading

total = 0
lock = threading.Lock()

def add():
    global total
    for _ in range(100_000):
        with lock:          # 拿到锁才执行
            total += 1      # 这行不是原子操作

threads = [threading.Thread(target=add) for _ in range(10)]
for t in threads: t.start()
for t in threads: t.join()
print(total)  # 1000000 ✅ 没锁会远小于这个数
```

**常见原语**：
| 原语 | 作用 |
|---|---|
| `Lock` | 互斥锁 |
| `RLock` | 可重入锁 |
| `Semaphore` | 限制并发数（如最多 5 个请求同时发） |
| `Event` | 一个线程通知另一个线程 |
| `Queue` | 线程安全队列 |

---

## 三、多进程（multiprocessing）

### 1. 适用场景

**一句话**：绕开 GIL，真正利用多核 CPU。

适合：
- CPU 密集计算（图片处理、模型推理）
- 需要进程隔离（一个 worker 崩了不影响其他）

### 2. Process vs Pool

```python
from multiprocessing import Process, Pool
import os

def square(x):
    print(f"进程 {os.getpid()} 计算 {x}")
    return x * x

# 单进程方式
p = Process(target=square, args=(5,))
p.start(); p.join()

# 进程池方式（推荐）
with Pool(processes=4) as pool:
    results = pool.map(square, [1, 2, 3, 4, 5])
    print(results)  # [1, 4, 9, 16, 25]
```

### 3. 进程 vs 线程 vs 协程 怎么选？

| 场景 | 选择 | 原因 |
|---|---|---|
| CPU 密集 | 多进程 | 绕开 GIL，多核并行 |
| 高并发 IO（万级） | asyncio | 单线程，切换极快，省资源 |
| 少量 IO（几十个） | 多线程 | 代码简单 |
| CPU + IO 混合 | 进程池 + asyncio | 外层并行，内层高并发 |

**项目例子**：
- 知识库切片：CPU 密集 → `ProcessPoolExecutor`
- embedding API 调用：IO 密集高并发 → `asyncio.gather`

---

## 四、生成器与迭代器（核心）

### 1. 三者区别

| 概念 | 通俗解释 | 关键方法 |
|---|---|---|
| 可迭代对象 | 能 `for` 循环的东西 | `__iter__` |
| 迭代器 | 能记住"现在轮到哪"的东西 | `__iter__`、`__next__` |
| 生成器 | 用 `yield` 写的"懒人版"迭代器 | 自动实现上面两个方法 |

```python
from collections.abc import Iterable, Iterator

lst = [1, 2, 3]
print(isinstance(lst, Iterable))  # True
print(isinstance(lst, Iterator))  # False

it = iter(lst)
print(isinstance(it, Iterator))   # True

print(next(it))  # 1
print(next(it))  # 2
print(next(it))  # 3
print(next(it))  # StopIteration 异常
```

### 2. 生成器函数

**一句话**：用 `yield` 把函数变成"暂停-继续"的机器。

```python
def count_down(n):
    print("开始")
    while n > 0:
        yield n      # 暂停，返回 n
        n -= 1
    print("结束")

for i in count_down(3):
    print(i)
```

**运行结果**：
```
开始
3
2
1
结束
```

**和列表推导的区别**：
```python
# 列表：一次性生成所有，占内存
[x for x in range(1_000_000)]  # 内存爆炸

# 生成器：用一个拿一个
(x for x in range(1_000_000))   # 几乎不占内存
```

### 3. yield 的 send / throw / close

```python
def echo():
    while True:
        val = yield
        print("收到:", val)

g = echo()
next(g)            # 先启动到第一个 yield
g.send("hello")    # 收到: hello
g.send("world")    # 收到: world
g.close()          # 关闭生成器
```

### 4. yield from 委托

**一句话**：把另一个生成器"接"到自己里面。

```python
def inner():
    yield 1
    yield 2
    return "inner 完成"

def outer():
    result = yield from inner()
    print(result)
    yield 3

print(list(outer()))
```

**运行结果**：
```
inner 完成
[1, 2, 3]
```

### 5. 项目应用：LLM 流式输出

```python
async def stream_llm(prompt):
    async for chunk in llm.astream(prompt):
        yield chunk.content  # 每个 token 立即吐出去

# FastAPI 使用
@app.get("/chat/stream")
async def chat(prompt: str):
    return StreamingResponse(stream_llm(prompt), media_type="text/event-stream")
```

**好处**：用户不用等模型全部生成完，看一个字出一个字。

---

## 五、装饰器（Decorator）

### 1. 基础：给函数"穿衣服"

**一句话**：不修改原函数，给它加功能。

```python
def timer(func):
    import time
    def wrapper(*args, **kwargs):
        start = time.time()
        result = func(*args, **kwargs)
        print(f"{func.__name__} 耗时 {time.time() - start:.3f}s")
        return result
    return wrapper

@timer
def slow():
    time.sleep(1)
    return "done"

slow()
```

**运行结果**：
```
slow 耗时 1.002s
```

### 2. 用 functools.wraps 保留名字

```python
from functools import wraps

def timer(func):
    @wraps(func)  # 保留 __name__、__doc__
    def wrapper(*args, **kwargs):
        ...
    return wrapper

# 没 @wraps 的话，slow.__name__ 会变成 'wrapper'
```

### 3. 带参数的装饰器

```python
def retry(times=3):
    def decorator(func):
        @wraps(func)
        def wrapper(*args, **kwargs):
            for i in range(times):
                try:
                    return func(*args, **kwargs)
                except Exception as e:
                    print(f"第 {i+1} 次失败: {e}")
                    if i == times - 1:
                        raise
        return wrapper
    return decorator

@retry(times=3)
def call_api():
    import random
    if random.random() < 0.7:
        raise ValueError("网络抖动")
    return "ok"

call_api()
```

### 4. 项目应用

```python
from functools import lru_cache

# embedding 缓存：相同 query 不用重复调用模型
@lru_cache(maxsize=2048)
def embed_cached(text: str):
    return embed_model.embed_query(text)

# 调用示例
v1 = embed_cached("什么是 RAG")
v2 = embed_cached("什么是 RAG")  # 直接命中缓存，0 耗时
```

---

## 六、Lambda

**一句话**：写一行匿名函数。

```python
# 排序
students = [{"name": "A", "age": 20}, {"name": "B", "age": 18}]
students.sort(key=lambda x: x["age"])
print(students)
# [{'name': 'B', 'age': 18}, {'name': 'A', 'age': 20}]

# map
print(list(map(lambda x: x * 2, [1, 2, 3])))
# [2, 4, 6]
```

**限制**：只能写单个表达式，复杂逻辑请用 `def`。

---

## 七、map / filter / reduce

### 1. map

```python
nums = [1, 2, 3, 4]
print(list(map(str, nums)))       # ['1', '2', '3', '4']
print(list(map(lambda x: x**2, nums)))  # [1, 4, 9, 16]
```

### 2. filter

```python
nums = [1, -2, 3, -4]
print(list(filter(lambda x: x > 0, nums)))  # [1, 3]
# 等价推导式：[x for x in nums if x > 0]
```

### 3. reduce

```python
from functools import reduce

nums = [1, 2, 3, 4]
print(reduce(lambda a, b: a + b, nums))      # 10
print(reduce(lambda a, b: a + b, nums, 100)) # 110（带初始值）
print(reduce(lambda a, b: a * b, nums))      # 24
```

**手写 reduce**（面试加分）：
```python
def myreduce(func, seq, init=None):
    it = iter(seq)
    acc = init if init is not None else next(it)
    for x in it:
        acc = func(acc, x)
    return acc
```

---

## 八、综合面试题

### Q1：GIL 是什么？多线程有用吗？

**答**：
GIL 保证同一时刻只有一个线程执行 Python 字节码。
- IO 密集：多线程有效，因为等待时 GIL 会释放
- CPU 密集：多线程无效，因为 GIL 在抢
- CPU 密集想并行：用多进程或 C 扩展

### Q2：asyncio 和多线程的区别？

**答**：
- asyncio：单线程协作式，适合高并发 IO（万级连接）
- 多线程：多线程抢占式，适合少量 IO 场景
- asyncio 切换更快，但不能有阻塞调用

### Q3：生成器和迭代器的区别？

**答**：
- 迭代器有 `__next__`，能记住位置
- 生成器是用 `yield` 写的迭代器，更简洁、省内存

### Q4：装饰器链怎么执行？

```python
@A
@B
@C
def f(): pass
# 等价 f = A(B(C(f)))
# 调用时：f() → A → B → C → 原函数 → C → B → A
```

### Q5：yield 和 return 的区别？

- `return`：结束函数，返回值
- `yield`：暂停函数，返回值；下次从暂停处继续

### Q6：深拷贝 vs 浅拷贝？

```python
import copy

a = [[1], [2]]
b = copy.copy(a)    # 浅拷贝
b[0].append(3)
print(a)  # [[1, 3], [2]]  ❌ 源数据变了

c = copy.deepcopy(a)  # 深拷贝
c[0].append(4)
print(a)  # [[1, 3], [2]]  ✅ 源数据没变
```

### Q7：*args 和 **kwargs？

```python
def demo(a, *args, **kwargs):
    print(a)      # 1
    print(args)   # (2, 3)
    print(kwargs) # {'x': 4, 'y': 5}

demo(1, 2, 3, x=4, y=5)
```

### Q8：Python 内存管理？

- 引用计数：为 0 立即释放
- 标记清除：解决循环引用
- 分代回收：gc 模块按代清理

### Q9：项目里怎么做性能优化？

| 场景 | 优化 |
|---|---|
| embedding API 慢 | `asyncio.gather` + `Semaphore` 限流 |
| markitdown 切片慢 | `ProcessPoolExecutor` 多进程切片 |
| 重复 query | `functools.lru_cache` 缓存 embedding |
| 大文件读取 | 生成器逐行 yield |
| JSON 序列化慢 | 用 `orjson` 替代 `json` |

---

## 九、一句话总结

- **异步**：单线程、等 IO 时换人，适合高并发网络请求
- **多线程**：GIL 限制，只适合做 IO 等待
- **多进程**：绕开 GIL，真正并行计算
- **生成器**：用 `yield` 省内存、做流式
- **装饰器**：不改代码给函数加功能
- **lambda/map/filter/reduce**：函数式小工具，简单场景用用就好

> **选型口诀：IO 密集看并发（asyncio > 多线程），CPU 密集看并行（多进程 > 多线程）**
