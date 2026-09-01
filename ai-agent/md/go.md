# Go 语言面试笔记

> 项目背景：机器人管控平台基于 goframe，调度系统基于 gin。Go 是服务端主力语言，面试核心围绕并发模型、内存管理、运行时和微服务实践。

---

## 一、Goroutine 与 GMP 调度模型

### 1. 什么是 Goroutine？

**一句话**：Go 里的"轻量级线程"，由 Go 运行时自己调度，启动成本极低（约 2KB 栈）。

```go
package main

import (
    "fmt"
    "time"
)

func say(msg string) {
    for i := 0; i < 3; i++ {
        fmt.Println(msg, i)
        time.Sleep(100 * time.Millisecond)
    }
}

func main() {
    go say("A")   // 启动一个 goroutine
    go say("B")   // 再启动一个
    time.Sleep(1 * time.Second) // 等它们跑完
    fmt.Println("main end")
}
```

**运行结果**：
```
A 0
B 0
A 1
B 1
A 2
B 2
main end
```

### 2. GMP 模型

```
        ┌─────────────┐
        │   全局队列    │
        │ [G1, G2, ...] │
        └──────┬──────┘
               │
    ┌──────────┼──────────┐
    │          │          │
┌───┴───┐  ┌───┴───┐  ┌───┴───┐
│   M   │  │   M   │  │   M   │   M：操作系统线程
│ ┌───┐ │  │ ┌───┐ │  │ ┌───┐ │
│ │ P │ │  │ │ P │ │  │ │ P │ │   P：逻辑处理器
│ │[G]│ │  │ │[G]│ │  │ │[G]│ │   G：Goroutine
│ │[G]│ │  │ │[G]│ │  │ │[G]│ │
│ │[G]│ │  │ │[G]│ │  │ │[G]│ │
│ └───┘ │  │ └───┘ │  │ └───┘ │
└───────┘  └───────┘  └───────┘
```

| 字母 | 含义 | 作用 |
|---|---|---|
| **G** | Goroutine | 待执行的任务 |
| **M** | Machine | 操作系统线程，真正执行代码 |
| **P** | Processor | 逻辑处理器，维护本地 goroutine 队列 |

**调度流程**：
1. 新 goroutine 优先放入当前 P 的本地队列
2. P 的本地队列满了，放全局队列
3. M 找 P 要 G 执行；P 空了从其他 P 偷一半 G（work stealing）
4. M 阻塞时（如系统调用），P 会和 M 解绑，换到其他 M 继续执行

**面试一句话**：GMP 让 Go 可以用少量 OS 线程调度大量 goroutine，实现高并发。

---

## 二、GC（垃圾回收）

### 1. 什么是三色标记法？

```
初始：所有对象都是白色

第 1 步：从根对象出发，把能直接到达的标记为灰色
┌───────┐     ┌───────┐     ┌───────┐
│ 根对象 │────→│ 灰色  │     │ 白色  │
└───────┘     │ 对象A │     │ 对象C │
              └───┬───┘     └───────┘
                  │
              ┌───┴───┐
              │ 白色  │
              │ 对象B │
              └───────┘

第 2 步：扫描灰色对象 A，把 A 引用的 B 变灰，A 变黑
┌───────┐     ┌───────┐     ┌───────┐
│ 根对象 │────→│ 黑色  │────→│ 灰色  │
└───────┘     │ 对象A │     │ 对象B │
              └───────┘     └───┬───┘
                                │
                            ┌───┴───┐
                            │ 白色  │
                            │ 对象C │
                            └───────┘

第 3 步：扫描灰色对象 B，没有引用新对象，B 变黑
          剩下白色的 C 就是垃圾，可回收
```

### 2. Go GC 发展

| 版本 | 特点 |
|---|---|
| Go 1.3 | 标记-清除，STW 明显 |
| Go 1.5 | 三色标记，并发标记 |
| Go 1.8 | 混合写屏障，STW 降到亚毫秒 |

### 3. 写屏障（Write Barrier）

**一句话**：在 GC 标记期间，修改对象引用时做记录，防止"黑色对象指向白色对象"漏标。

```go
// 比如黑色对象 A 新引用了白色对象 C
A.next = C
// 写屏障会把 C 重新标灰，避免 C 被误回收
```

### 4. 调优参数

```bash
# 设置 GC 目标：堆内存增长到原来的 2 倍时触发下一次 GC
GOGC=100    # 默认
GOGC=off    # 关闭 GC（测试/特殊场景）
GOMEMLIMIT=1GiB  # 软内存上限
```

### 5. 怎么排查 GC 问题？

```bash
go tool pprof http://localhost:6060/debug/pprof/heap
go tool pprof http://localhost:6060/debug/pprof/allocs
go tool trace trace.out
```

---

## 三、切片（Slice）底层与扩容

### 1. 切片结构

```go
type slice struct {
    array unsafe.Pointer  // 指向底层数组的指针
    len   int             // 长度
    cap   int             // 容量
}
```

```go
arr := [5]int{1, 2, 3, 4, 5}
s := arr[1:4]  // 切片 [2, 3, 4]
// len=3, cap=4（从 arr[1] 到 arr 末尾还有 4 个位置）
fmt.Println(len(s), cap(s)) // 3 4
```

### 2. 切片共享底层数组

```go
a := []int{1, 2, 3, 4, 5}
b := a[1:3]
b[0] = 100
fmt.Println(a) // [1, 100, 3, 4, 5]
```

> ⚠️ 切片修改会影响原数组，这是面试常考陷阱。

### 3. append 扩容规则

**Go 1.18 之前**：
- 容量 < 1024：翻倍
- 容量 ≥ 1024：增长 1.25 倍

**Go 1.18 之后**：
- 引入了更平滑的扩容曲线，小切片翻倍，大切片约 1.25 倍

```go
s := make([]int, 0, 2)
s = append(s, 1, 2)
fmt.Println(len(s), cap(s)) // 2 2

s = append(s, 3)            // 触发扩容
fmt.Println(len(s), cap(s)) // 3 4
```

### 4. 常见坑

```go
func foo() []int {
    a := []int{1, 2, 3}
    b := append(a[:1], a[2:]...)  // [1, 3]，但改的是 a 的底层数组
    return b
}
```

---

## 四、Map 底层与扩容

### 1. 底层结构

Go map 是 **哈希表**，底层结构 `hmap`：
- `buckets`：桶数组，每个桶存 8 个 key-value
- 溢出桶：一个桶满了用溢出桶链接

```
┌────────────────────────────────────┐
│               hmap                 │
│  buckets ──→ [bucket0][bucket1]... │
│  count    overflow                 │
└────────────────────────────────────┘

bucket:
┌─────────┬─────────┬─────────┐
│ key0-v0 │ key1-v1 │ ... x8  │
└─────────┴─────────┴─────────┘
```

### 2. 扩容机制

**触发条件**：
- 负载因子 > 6.5（平均每个桶超过 6.5 个元素）
- 溢出桶太多

**扩容方式**：
- **等量扩容**：数据不多但溢出桶多，重新整理，桶数不变
- **翻倍扩容**：数据量大，桶数量翻倍

**渐进式扩容**：
- 不一次性搬完，每次读写 map 时搬一部分
- 访问旧桶时会触发搬迁

### 3. 并发不安全

```go
m := make(map[int]int)

// ❌ 并发读写会 panic
for i := 0; i < 1000; i++ {
    go func(n int) {
        m[n] = n
    }(i)
}
```

**解决方案**：
- `sync.RWMutex` + map
- `sync.Map`（适合读多写少）

---

## 五、Channel

### 1. 基本概念

**一句话**：channel 是 goroutine 之间通信的"管道"，遵循 CSP（Communicating Sequential Processes）模型。

```go
package main

import "fmt"

func main() {
    ch := make(chan int, 2)  // 缓冲区大小为 2

    ch <- 1
    ch <- 2

    fmt.Println(<-ch) // 1
    fmt.Println(<-ch) // 2
}
```

### 2. 有缓冲 vs 无缓冲

| 类型 | 特点 |
|---|---|
| 无缓冲 `make(chan int)` | 发送和接收必须同时准备好，否则阻塞（同步） |
| 有缓冲 `make(chan int, n)` | 缓冲未满可继续发送，未满可继续接收（异步） |

### 3. select 多路复用

```go
ch1 := make(chan string)
ch2 := make(chan string)

select {
case msg1 := <-ch1:
    fmt.Println("ch1:", msg1)
case msg2 := <-ch2:
    fmt.Println("ch2:", msg2)
default:
    fmt.Println("都没有数据")
}
```

### 4. 关闭 channel

```go
close(ch)

// 读取已关闭 channel：返回零值 + ok=false
v, ok := <-ch
if !ok {
    fmt.Println("channel 已关闭")
}
```

> ⚠️ 不能重复关闭 channel，不能向已关闭 channel 发送数据，都会 panic。

### 5. Channel 底层实现原理

**核心结构**（简化版）：

```go
type hchan struct {
    qcount   uint           // 当前队列中元素个数
    dataqsiz uint           // 循环队列大小（缓冲区大小）
    buf      unsafe.Pointer // 指向循环队列
    elemsize uint16         // 每个元素大小
    closed   uint32         // 是否已关闭
    sendx    uint           // 发送位置下标
    recvx    uint           // 接收位置下标
    recvq    waitq          // 等待接收的 goroutine 队列
    sendq    waitq          // 等待发送的 goroutine 队列
    lock     mutex          // 互斥锁，保护 hchan
}
```

**核心原理**：
- channel 本质上是一个**带锁的循环队列 + 两个等待队列**
- 发送时：缓冲区有空位就放到 `buf`，没空位就把自己挂到 `sendq` 等待
- 接收时：缓冲区有数据就取走，没数据就把自己挂到 `recvq` 等待
- 所有对 channel 的操作都受 `lock` 保护，所以 channel 是**并发安全**的

**发送流程**：
```
ch <- v
   │
   ▼
加锁
   │
   ├── 有等待接收者？──→ 直接把值给对方，唤醒对方 goroutine
   │
   ├── 缓冲区没满？──→ 放入 buf[sendx]，sendx++
   │
   └── 缓冲区满了？──→ 当前 goroutine 挂到 sendq，阻塞等待
```

**接收流程**：
```
v := <-ch
   │
   ▼
加锁
   │
   ├── 有等待发送者？──→ 直接拿值，唤醒对方
   │
   ├── 缓冲区有数据？──→ 取 buf[recvx]，recvx++
   │
   └── 缓冲区空了？──→ 当前 goroutine 挂到 recvq，阻塞等待
```

### 6. Select 底层实现原理

**一句话**：select 会**同时监控多个 channel**，谁先就绪就执行谁；多个同时就绪时**随机选择**。

**底层机制**：
1. 对所有 case 的 channel 按地址**排序**，减少死锁概率
2. 依次尝试每个 case：
   - 如果某个 channel 可以发送/接收，就直接执行
3. 如果没有就绪：
   - 有 `default`：执行 default
   - 无 `default`：把当前 goroutine 挂到所有 channel 的等待队列上，阻塞等待
4. 任意一个 channel 就绪后，唤醒当前 goroutine，再从等待队列移除自己

**为什么多个 case 就绪要随机选？**

防止某个 channel 总是优先被选中，造成"饥饿"。

### 7. Channel 死锁

**一句话**：当所有 goroutine 都因为 channel 阻塞，且没有人能唤醒它们时，就发生死锁。

**常见死锁场景**：

```go
// 场景 1：无缓冲 channel，发送方没有接收方
func main() {
    ch := make(chan int)
    ch <- 1  // 永远阻塞，main 自己也被卡住 → fatal error: all goroutines are asleep
}
```

```go
// 场景 2：通道读写顺序错误
func main() {
    ch := make(chan int)
    <-ch        // 先接收，但没人发送，阻塞
    ch <- 1     // 这行永远执行不到
}
```

```go
// 场景 3：互相等待
func main() {
    ch1 := make(chan int)
    ch2 := make(chan int)

    go func() {
        <-ch1
        ch2 <- 2
    }()

    ch1 <- 1
    <-ch2
    // 上面两行在主 goroutine 里按顺序执行，子 goroutine 还没启动就到 <-ch1？
    // 实际上这里不会死锁，但类似的双向等待模式很容易写出死锁
}
```

**如何避免死锁**：
- 无缓冲 channel 保证**发送和接收配对**
- 有缓冲 channel 注意缓冲区大小，不要写满后没人读
- 用 `select + default` 做非阻塞读写
- 复杂场景用 context 控制超时和取消

---

## 六、并发编程

### 1. sync.WaitGroup

```go
var wg sync.WaitGroup

for i := 0; i < 3; i++ {
    wg.Add(1)
    go func(n int) {
        defer wg.Done()
        fmt.Println(n)
    }(i)
}

wg.Wait() // 等所有 goroutine 完成
```

### 2. sync.Mutex / RWMutex

```go
// 写多读少用（如状态，计数类）
var counter int
var mu sync.Mutex

func add() {
    mu.Lock()
    counter++
    mu.Unlock()
}
```

```go
// 读多写少用（如集群信息、机器信息） RWMutex
var rw sync.RWMutex
var data map[string]int

func read(key string) int {
    rw.RLock()
    defer rw.RUnlock()
    return data[key]
}

func write(key string, val int) {
    rw.Lock()
    defer rw.Unlock()
    data[key] = val
}
```

### 3. sync.Once

```go
var once sync.Once
var instance *singleton

func GetInstance() *singleton {
    once.Do(func() {
        instance = &singleton{}
    })
    return instance
}
```

### 4. context 上下文

**一句话**：用来在 goroutine 之间传递**取消信号、超时、截止时间、请求元数据**。

```go
ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
defer cancel()

select {
case <-time.After(3 * time.Second):
    fmt.Println("任务完成")
case <-ctx.Done():
    fmt.Println("超时取消:", ctx.Err()) // context deadline exceeded
}
```

**常用函数**：

| 函数 | 用途 |
|---|---|
| `context.Background()` | 根 context，最顶层用 |
| `context.TODO()` | 临时占位 |
| `WithCancel` | 手动取消 |
| `WithTimeout` | 超时自动取消 |
| `WithDeadline` | 到某个时间点取消 |
| `WithValue` | 传请求元数据（如 trace_id） |

**项目应用**：
```go
// HTTP 请求链路：控制器 → 服务 → DB，一层层传 ctx
func handler(ctx context.Context, req *Request) {
    ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
    defer cancel()
    service.DoSomething(ctx, req)
}
```

---

## 七、各种锁

| 锁 | 特点 | 场景 |
|---|---|---|
| `sync.Mutex` | 互斥锁，读写都互斥 | 写多读少 |
| `sync.RWMutex` | 读锁共享，写锁互斥 | 读多写少 |
| `sync.Once` | 只执行一次 | 初始化、单例 |
| `atomic` | 原子操作，无锁 | 简单计数器 |

```go
var count int64
atomic.AddInt64(&count, 1)
```

---

## 八、内存逃逸分析

### 1. 什么是逃逸？

**一句话**：本应该在栈上分配的变量，因为某些原因被分配到了堆上。

### 2. 常见逃逸场景

```go
// 1. 返回局部变量指针
func foo() *int {
    a := 1
    return &a  // a 逃逸到堆
}

// 2. 闭包引用局部变量
func bar() func() int {
    a := 1
    return func() int {
        return a  // a 逃逸
    }
}

// 3. 大对象或不确定大小
func big() {
    s := make([]int, 100000)  // 大切片可能逃逸
    _ = s
}

// 4. interface{} 传参
func print(x interface{}) {
    fmt.Println(x)  // fmt.Println 参数是 interface{}，经常导致逃逸
}
```

### 3. 查看逃逸

```bash
go build -gcflags="-m" main.go
```

### 4. 闭包内存逃逸以及闭包变量异常原理

#### 错误示例（经典坑）
```go
package main

import (
	"fmt"
	"time"
)

func main() {
	for i := 0; i < 5; i++ {
		go func() {
			fmt.Println(i) // 闭包直接引用循环变量 i
		}()
	}
	time.Sleep(time.Second)
}
```
**预期输出**：`0 1 2 3 4`
**实际经常输出**：`5 5 5 5 5`，结果还可能随机混乱，并不稳定。

---
##### 核心根本原因
###### 1. for‑loop 的迭代变量**内存复用**
Go 的 `for i:=0;...` / `for range`，**循环变量只会创建一次，全程共用同一个内存地址**，每次迭代只是修改这块内存上的值，并不会每次循环新建一个变量。

```go
for i := 0; i < 5; i++ { 
    // i 的地址 &i 在整个循环全程不变
}
```

###### 2. Go闭包捕获的是**变量引用(内存地址)**，不是瞬间的值快照
匿名闭包函数不会把当前这一刻 `i` 的数值复制保存下来；
闭包里面存的只是 `i` 的**内存地址**，等到 goroutine 真正被 CPU 调度、执行 `fmt.Println(i)` 的时候，才会去这个地址读取 i 的值。

> ⚠️ 关键点：**启动goroutine ≠ 立刻执行goroutine**
> `go func(){...}()` 只是把协程放入运行队列，并不会马上抢占执行；
> 主线程循环速度极快，往往循环全部跑完，`i` 已经自增到 `5`，此时5个协程才开始执行。所有协程读取同一个内存地址，读到的值自然都是5。

###### 3. 调度时序不确定性 → 结果随机
并不是每次一定全部打印 `5`；
如果某一个goroutine刚好在主线程还没修改i之前抢到CPU执行，就会读到当时i的值。
因此输出结果随机混乱，属于典型**并发时序bug，偶现问题，很难复现排查**。

###### 4. 变量逃逸
因为闭包协程引用了循环变量，编译器逃逸分析会把循环变量 `i` 从栈上分配**逃逸到堆内存**，即使主线程循环结束，goroutine依旧可以访问这块堆内存。

---
##### range循环一模一样的坑
```go
items := []string{"A","B","C"}
for _, v := range items {
	go func() {
		fmt.Println(v)
	}()
}
time.Sleep(time.Second)
```
`v` 全程同一个地址，大概率打印三次 `C`。

---
##### 4种正确修复方案
###### 方案1：作为参数传入goroutine（最推荐，值拷贝）
调用的时候，实参 `i` **立刻完成一次值拷贝**，复制一份独立的值给匿名函数的形参，每个协程拥有自己独立副本，和外层循环变量彻底解绑。
```go
for i := 0; i < 5; i++ {
	go func(val int) {
		fmt.Println(val)
	}(i) // 此处立刻拷贝当前i的值
}
```

###### 方案2：循环体内新建局部变量
每次循环迭代，在循环块内部声明新变量，开辟全新内存地址，每个闭包捕获不同的变量。
```go
for i := 0; i < 5; i++ {
	val := i // 每次循环生成新变量，新地址
	go func() {
		fmt.Println(val)
	}()
}
```
> 原理：`val` 在 `{}` 循环块内，每一轮迭代都会创建新变量。

###### 方案3：同步等待，让协程先跑完（不推荐，牺牲并发）
用`sync.WaitGroup`，但仅仅加WaitGroup**不能修复这个闭包问题**！WaitGroup只做等待，不会拷贝变量。
错误示范（依旧会出现5,5,5）：
```go
var wg sync.WaitGroup
for i :=0;i<5;i++{
    wg.Add(1)
    go func(){
        defer wg.Done()
        fmt.Println(i)
    }()
}
wg.Wait()
```
WaitGroup必须配合上面传参/新建局部变量一起使用。

###### 方案4：循环内runtime.Gosched()让出调度（不推荐，不稳定）
主动让渡CPU给goroutine执行，只是概率修复，不能保证每次生效，生产代码禁止使用。

---
##### 对比总结
|写法|行为|结果|
|---|---|---|
|`go func(){fmt.Println(i)}()`|闭包捕获i的引用，共享内存|结果错乱，经常全是5|
|`go func(v int){fmt.Println(v)}(i)`|启动goroutine时拷贝i当前值，每个协程独立副本|正确输出0‑4|
|`val:=i; go func(){fmt.Println(val)}()`|每次循环新建局部变量val，闭包捕获各自val|正确输出0‑4|

##### 避坑口诀
> **协程闭包用循环变量，一定要做值拷贝；不要直接捕获外层循环迭代变量。**

---

## 九、堆栈分配

| 特性 | 栈 | 堆 |
|---|---|---|
| 分配速度 | 极快（移动栈指针） | 慢（需要 GC 管理） |
| 回收 | 函数返回自动回收 | GC 回收 |
| 大小 | 初始 2KB，可自动增长 | 大 |
| 存储内容 | 局部变量、函数参数 | 全局变量、逃逸变量、大对象 |

**Go 栈特点**：
- 每个 goroutine 初始栈 2KB
- 栈空间不足时自动扩容，最大可达 1GB（64 位）
- 栈是连续内存，扩缩容时复制到新的连续空间

---

## 十、值拷贝 vs 引用拷贝

```go
type Person struct {
    Name string
    Age  int
}

func byValue(p Person) {
    p.Age = 100
}

func byRef(p *Person) {
    p.Age = 100
}

func main() {
    a := Person{Name: "A", Age: 20}
    byValue(a)
    fmt.Println(a.Age) // 20（没改）

    byRef(&a)
    fmt.Println(a.Age) // 100（改了）
}
```

| 类型 | 传递方式 | 修改是否影响原值 |
|---|---|---|
| 基本类型、struct、array | 值拷贝 | 不影响 |
| slice、map、channel、指针 | 引用拷贝 | 影响 |

> ⚠️ slice 和 map 本身是引用类型，但内部结构是值拷贝，append 扩容会导致底层数组变化。

---

## 十一、资源泄漏

### 常见类型

| 类型 | 原因 | 解决 |
|---|---|---|
| goroutine 泄漏 | goroutine 永远阻塞，退不出 | 用 context 取消、带缓冲 channel、timeout |
| channel 泄漏 | 发送方没有接收方 | 确保有接收、合理关闭 |
| 数据库连接泄漏 | 连接没关闭 | `defer rows.Close()` |
| 文件句柄泄漏 | 文件没关闭 | `defer f.Close()` |

### goroutine 泄漏示例

```go
// ❌ 泄漏：ch 没有接收方，goroutine 永远阻塞
func leak() {
    ch := make(chan int)
    go func() {
        ch <- 1  // 阻塞
    }()
}

// ✅ 用缓冲 channel 或有接收方
func noLeak() {
    ch := make(chan int, 1)
    go func() {
        ch <- 1
    }()
    <-ch
}
```

---

## 十二、微服务架构

### 1. Go 微服务常用组件

| 领域 | 常用方案 |
|---|---|
| Web 框架 | gin、echo、goframe |
| RPC 框架 | gRPC、go-zero、kitex |
| 配置中心 | etcd、nacos、apollo |
| 服务发现 | etcd、consul、nacos |
| 消息队列 | Kafka、RabbitMQ、RocketMQ |
| 链路追踪 | Jaeger、Zipkin |
| 监控 | Prometheus + Grafana |
| 日志 | zap + ELK / Loki |

### 2. 项目实战：调度系统的微服务拆分

```
┌─────────────┐      ┌─────────────┐      ┌─────────────┐
│   API 网关   │────→│  算力资源服务 │────→│  任务调度服务 │
│   (gin)     │      │  (goframe)  │      │  (go-zero)  │
└─────────────┘      └──────┬──────┘      └──────┬──────┘
                            │                    │
                            ↓                    ↓
                     ┌─────────────┐      ┌─────────────┐
                     │  prometheus │      │     k8s     │
                     │   minio     │      │   算力卡     │
                     └─────────────┘      └─────────────┘
```

**关键技术点**：
- 通过 k8s client 动态创建训练 Job
- 用 etcd 做任务状态同步
- prometheus 采集算力指标
- minio 存模型和数据集

---

## 十三、面试高频问题

### Q1：Goroutine 和线程的区别？

**答**：
- goroutine 由 Go 运行时调度，线程由 OS 调度
- goroutine 栈 2KB 起步，可自动扩缩；线程栈固定 1-8MB
- 切换成本：goroutine 约 200ns，线程约 1-2μs
- 一个 Go 程序可轻松启动百万级 goroutine，线程通常几千就吃力

### Q2：channel 是引用类型还是值类型？

**答**：channel 是引用类型，但函数参数传递的是指针的拷贝，所以可以在函数内部继续操作同一个 channel。

### Q3：select 会随机选择吗？

**答**：当多个 case 同时就绪时，select 会**随机**选择一个执行，避免饥饿。

### Q4：defer 的执行顺序？

```go
func main() {
    defer fmt.Println("1")
    defer fmt.Println("2")
    defer fmt.Println("3")
}
```

**运行结果**：
```
3
2
1
```

> defer 是**栈**（LIFO），先进后出。

### Q5：new 和 make 的区别？

| 函数 | 用途 | 返回 |
|---|---|---|
| `new` | 分配内存 | 返回指针 |
| `make` | 初始化 slice、map、channel | 返回引用类型本身 |

```go
a := new(int)       // *int，值为 0
b := make([]int, 3) // []int
```

### Q6：interface 的底层实现？

**答**：
```go
type iface struct {
    tab  *itab          // 类型信息 + 方法表
    data unsafe.Pointer // 实际数据指针
}
```

- 空接口 `interface{}` 可以存任意类型
- 非空接口包含具体类型和方法集
- 类型断言 `x.(T)` 失败会 panic，可用 `v, ok := x.(T)`

### Q7：Go 是值传递还是引用传递？

**答**：Go 只有**值传递**。slice、map、channel、指针看起来是"引用传递"，实际上传递的是它们内部结构或地址的拷贝。

### Q8：怎么实现一个线程安全的单例？

```go
type singleton struct{}

var (
    instance *singleton
    once     sync.Once
)

func GetInstance() *singleton {
    once.Do(func() {
        instance = &singleton{}
    })
    return instance
}
```

### Q9：GC 触发时机？

**答**：
- 主动触发：`runtime.GC()`
- 堆内存达到阈值：默认是上次 GC 后存活对象的 2 倍（GOGC=100）
- 定时触发：2 分钟未 GC 会强制触发

### Q10：怎么做性能优化？

| 场景 | 优化 |
|---|---|
| 高并发 | 用 goroutine + channel，避免锁竞争 |
| 内存高 | 减少逃逸、对象池 sync.Pool、复用切片 |
| GC 频繁 | 增大 GOGC、减少小对象分配 |
| 数据库慢 | 连接池、批量查询、索引 |
| JSON 慢 | 避免 interface{}、用 jsoniter |

---

## 十四、高并发高可用实战示例

### 1. 高并发限流器（令牌桶）

**一句话**：以固定速率往桶里放令牌，请求先拿令牌，拿到才能执行。

```go
package main

import (
    "context"
    "fmt"
    "time"

    "golang.org/x/time/rate"
)

func main() {
    // 每秒产生 10 个令牌，桶容量 20
    limiter := rate.NewLimiter(10, 20)

    for i := 0; i < 30; i++ {
        // Wait 会阻塞等到拿到令牌
        err := limiter.Wait(context.Background())
        if err != nil {
            fmt.Println("error:", err)
            return
        }
        fmt.Println("处理请求", i, time.Now().Format("15:04:05.000"))
    }
}
```

**运行结果**：前 20 个请求瞬间通过（桶里预存），后面每秒通过 10 个。

---

### 2. 高并发爬虫（控制最大并发数）

```go
package main

import (
    "fmt"
    "sync"
    "time"
)

func worker(id int, jobs <-chan int, wg *sync.WaitGroup) {
    defer wg.Done()
    for j := range jobs {
        fmt.Printf("worker %d 处理任务 %d\n", id, j)
        time.Sleep(500 * time.Millisecond) // 模拟耗时
    }
}

func main() {
    jobs := make(chan int, 100)
    var wg sync.WaitGroup

    // 启动 5 个 worker，最多同时处理 5 个任务
    for w := 1; w <= 5; w++ {
        wg.Add(1)
        go worker(w, jobs, &wg)
    }

    for j := 1; j <= 20; j++ {
        jobs <- j
    }
    close(jobs)

    wg.Wait()
    fmt.Println("全部完成")
}
```

**运行结果**：20 个任务分批处理，每批 5 个，总共约 4 批，耗时 2 秒左右。

---

### 3. 高可用 HTTP 服务（graceful shutdown）

```go
package main

import (
    "context"
    "fmt"
    "net/http"
    "os"
    "os/signal"
    "syscall"
    "time"
)

func main() {
    srv := &http.Server{
        Addr:    ":8080",
        Handler: http.HandlerFunc(hello),
    }

    // 启动服务
    go func() {
        if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
            fmt.Println("listen error:", err)
        }
    }()

    // 等待退出信号
    quit := make(chan os.Signal, 1)
    signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
    <-quit

    fmt.Println("shutting down server...")

    // 优雅关闭：给正在处理的请求 5 秒收尾时间
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()

    if err := srv.Shutdown(ctx); err != nil {
        fmt.Println("server forced to shutdown:", err)
    }
    fmt.Println("server exiting")
}

func hello(w http.ResponseWriter, r *http.Request) {
    time.Sleep(2 * time.Second)
    fmt.Fprintln(w, "hello")
}
```

**关键点**：
- `srv.Shutdown(ctx)` 会关闭监听，等待已有请求完成
- 避免直接 kill 导致请求中断

---

### 4. 高并发计数器（atomic 无锁）

```go
package main

import (
    "fmt"
    "sync"
    "sync/atomic"
)

func main() {
    var counter int64
    var wg sync.WaitGroup

    for i := 0; i < 1000; i++ {
        wg.Add(1)
        go func() {
            defer wg.Done()
            atomic.AddInt64(&counter, 1)
        }()
    }

    wg.Wait()
    fmt.Println("counter:", counter) // 1000
}
```

**优势**：atomic 是 CPU 指令级别，比 Mutex 快得多，适合简单计数。

---

### 5. 对象池减少 GC 压力（sync.Pool）

```go
package main

import (
    "bytes"
    "fmt"
    "sync"
)

var pool = sync.Pool{
    New: func() interface{} {
        return new(bytes.Buffer)
    },
}

func main() {
    // 获取对象
    buf := pool.Get().(*bytes.Buffer)
    buf.WriteString("hello world")
    fmt.Println(buf.String())

    // 重置并放回池中复用
    buf.Reset()
    pool.Put(buf)
}
```

**适用场景**：频繁创建销毁的大对象（如 buffer、结构体），复用减少 GC。

---

### 6. 熔断降级示例（go-zero breaker）

```go
package main

import (
    "errors"
    "fmt"
    "github.com/zeromicro/go-zero/core/breaker"
)

func main() {
    b := breaker.NewBreaker()

    for i := 0; i < 100; i++ {
        err := b.Do("request", func() error {
            // 模拟下游故障
            return errors.New("service unavailable")
        })
        if err != nil {
            fmt.Println(i, err)
        }
    }
}
```

**效果**：连续失败后，熔断器打开，后续请求直接返回错误，不再调用下游，保护系统。

---

### 7. 项目实战：调度系统高并发任务下发

```
请求接入层（gin）
    │
    ↓ 把任务写入 Kafka / RabbitMQ
消息队列
    │
    ↓ 多个消费者 goroutine 并发拉取
任务调度服务（go-zero）
    │
    ↓ 调用 k8s client 创建训练 Job
K8s 集群
```

**高并发手段**：
- 接入层：限流 + 负载均衡
- 任务队列：削峰填谷，避免瞬时打爆 k8s
- 消费者：固定 worker 数量，避免 goroutine 爆炸
- 调用 k8s：加熔断 + 重试 + 指数退避
- 状态同步：etcd watch 实时感知任务状态

---

## 十五、高并发高可用设计原则

| 原则 | 说明 |
|---|---|
| **无状态** | 服务本身不存状态，方便横向扩展 |
| **缓存** | 热点数据放缓存，减少 DB 压力 |
| **异步** | 非核心流程异步化，降低响应时间 |
| **限流** | 保护自己和下游，防止被流量打崩 |
| **熔断降级** | 故障时快速失败，返回兜底数据 |
| **超时重试** | 避免长时间阻塞，重试带退避 |
| **负载均衡** | 多实例分摊流量 |
| **监控告警** | Prometheus + Grafana + 告警 |

---

## 十六、一句话总结

- **GMP**：G 是任务，M 是线程，P 是调度器，work stealing 实现高并发
- **GC**：三色标记 + 写屏障，Go 1.8 后 STW 很短
- **Slice**：底层数组 + len/cap，append 可能扩容，注意共享底层数组
- **Map**：哈希表，负载因子 6.5 或溢出桶多会扩容，并发不安全
- **Channel**：goroutine 通信管道，有缓冲/无缓冲两种
- **Context**：传取消信号、超时、元数据
- **锁**：Mutex、RWMutex、Once、atomic 按场景选
- **高并发**：限流、worker pool、atomic、对象池、熔断
- **高可用**：graceful shutdown、无状态、负载均衡、监控告警
- **微服务**：gin/goframe 做业务，go-zero/gRPC 做 RPC，etcd 做注册发现，k8s 做部署调度

---

## 十七、并发模型与并发模式

### 1. CSP 并发模型

**一句话**：Go 采用 **CSP（Communicating Sequential Processes）** 并发模型，核心思想是"**不要通过共享内存通信，而要通过通信共享内存**"。

```
传统多线程：
线程 A ──→ 共享变量 ←── 线程 B
          ↑ 加锁保护

CSP：
线程 A ──→ channel ──→ 线程 B
            通信传递数据
```

**优势**：
- 数据所有权清晰
- 避免锁竞争
- 更容易推理和避免死锁

---

### 2. 常见并发模式

#### （1）生产者-消费者模式

```go
package main

import (
    "fmt"
    "sync"
    "time"
)

func producer(ch chan<- int, wg *sync.WaitGroup) {
    defer wg.Done()
    for i := 1; i <= 5; i++ {
        ch <- i
        fmt.Println("生产:", i)
        time.Sleep(100 * time.Millisecond)
    }
}

func consumer(ch <-chan int, wg *sync.WaitGroup) {
    defer wg.Done()
    for v := range ch {
        fmt.Println("消费:", v)
        time.Sleep(200 * time.Millisecond)
    }
}

func main() {
    ch := make(chan int, 3)
    var wg sync.WaitGroup

    wg.Add(1)
    go producer(ch, &wg)

    wg.Add(1)
    go consumer(ch, &wg)

    wg.Wait()
    close(ch)
}
```

**场景**：消息队列、任务流水线、日志处理。

---

#### （2）Worker Pool 模式

```go
package main

import (
    "fmt"
    "sync"
    "time"
)

func worker(id int, jobs <-chan int, results chan<- int, wg *sync.WaitGroup) {
    defer wg.Done()
    for j := range jobs {
        fmt.Printf("worker %d 处理任务 %d\n", id, j)
        time.Sleep(time.Second)
        results <- j * 2
    }
}

func main() {
    jobs := make(chan int, 100)
    results := make(chan int, 100)

    var wg sync.WaitGroup
    // 3 个 worker
    for w := 1; w <= 3; w++ {
        wg.Add(1)
        go worker(w, jobs, results, &wg)
    }

    // 9 个任务
    for j := 1; j <= 9; j++ {
        jobs <- j
    }
    close(jobs)

    wg.Wait()
    close(results)

    for r := range results {
        fmt.Println("结果:", r)
    }
}
```

**场景**：控制并发数、批量任务处理、爬虫、图片处理。

---

#### （3）Pipeline 流水线模式

```go
package main

import "fmt"

// 阶段 1：生成数据
func gen(nums ...int) <-chan int {
    out := make(chan int)
    go func() {
        defer close(out)  // goroutine 退出前关闭 channel
        for _, n := range nums {
            out <- n
        }
    }()
    return out
}

// 阶段 2：平方
func sq(in <-chan int) <-chan int {
    out := make(chan int)
    go func() {
        defer close(out)  // goroutine 退出前关闭 channel
        for n := range in {
            out <- n * n
        }
    }()
    return out
}

func main() {
    c := gen(2, 3, 4)
    out := sq(c)

    for v := range out {
        fmt.Println(v) // 4, 9, 16
    }
}
```

**为什么返回 channel 没问题？**

1. **函数立即返回 `out`，但 goroutine 还在运行**，继续往 `out` 里发数据
2. **调用方用 `for range out` 读取**，能正常收到所有数据
3. **goroutine 发完数据后 `close(out)`**，调用方的 `range` 检测到关闭就退出循环
4. **关键规则：谁发送谁关闭**。这里 `gen` 和 `sq` 内部启动的 goroutine 是发送方，所以由它们关闭 channel

> ⚠️ 注意：返回的 channel 不能在外部再关闭，否则会 panic。发送方 goroutine 已经负责关闭了。

**场景**：数据处理流水线、ETL、多阶段计算。

---

#### （4）Fan-Out / Fan-In 模式

**Fan-Out**：一个输入分发到多个 goroutine 并行处理。
**Fan-In**：多个 goroutine 的结果合并到一个 channel。

```go
package main

import (
    "fmt"
    "sync"
)

// Fan-Out: 多个 worker 同时消费
func fanOut(in <-chan int, workerCount int) []<-chan int {
    outs := make([]<-chan int, workerCount)
    for i := 0; i < workerCount; i++ {
        ch := make(chan int)
        outs[i] = ch
        go func(out chan<- int) {
            defer close(out)
            for n := range in {
                out <- n * n
            }
        }(ch)
    }
    return outs
}

// Fan-In: 合并多个 channel
func fanIn(chs ...<-chan int) <-chan int {
    out := make(chan int)
    var wg sync.WaitGroup
    for _, ch := range chs {
        wg.Add(1)
        go func(c <-chan int) {
            defer wg.Done()
            for n := range c {
                out <- n
            }
        }(ch)
    }
    go func() {
        wg.Wait()
        close(out)
    }()
    return out
}

func main() {
    in := make(chan int)
    go func() {
        for i := 1; i <= 9; i++ {
            in <- i
        }
        close(in)
    }()

    outs := fanOut(in, 3)
    merged := fanIn(outs...)

    for v := range merged {
        fmt.Println(v)
    }
}
```

**场景**：并行计算、MapReduce、分布式任务。

---

#### （5）Context 取消传播模式

```go
package main

import (
    "context"
    "fmt"
    "time"
)

func worker(ctx context.Context, id int) {
    for {
        select {
        case <-ctx.Done():
            fmt.Printf("worker %d 退出\n", id)
            return
        default:
            fmt.Printf("worker %d 工作中\n", id)
            time.Sleep(300 * time.Millisecond)
        }
    }
}

func main() {
    ctx, cancel := context.WithCancel(context.Background())

    for i := 1; i <= 3; i++ {
        go worker(ctx, i)
    }

    time.Sleep(1 * time.Second)
    cancel() // 通知所有 worker 退出
    time.Sleep(500 * time.Millisecond)
}
```

**场景**：批量任务取消、超时控制、优雅关闭。

---

#### （6）读写锁分离模式

```go
package main

import (
    "fmt"
    "sync"
)

var (
    cache = make(map[string]string)
    rw    sync.RWMutex
)

func get(key string) string {
    rw.RLock()
    defer rw.RUnlock()
    return cache[key]
}

func set(key, val string) {
    rw.Lock()
    defer rw.Unlock()
    cache[key] = val
}

func main() {
    set("name", "go")
    fmt.Println(get("name"))
}
```

**场景**：配置缓存、热点数据、读多写少的共享资源。

---

### 3. 并发模式选择速查

| 模式 | 核心组件 | 适用场景 |
|---|---|---|
| 生产者-消费者 | channel + goroutine | 任务缓冲、解耦生产消费速度 |
| Worker Pool | 固定数量 worker | 控制并发、批量处理 |
| Pipeline | 多个 channel 串联 | 多阶段数据处理 |
| Fan-Out/Fan-In | 多个 goroutine + 合并 | 并行计算、聚合结果 |
| Context 取消 | context.Context | 超时、取消、优雅退出 |
| 读写锁分离 | sync.RWMutex | 读多写少的热点数据 |

---

## 十八、一句话总结

- **GMP**：G 是任务，M 是线程，P 是调度器，work stealing 实现高并发
- **GC**：三色标记 + 写屏障，Go 1.8 后 STW 很短
- **Slice**：底层数组 + len/cap，append 可能扩容，注意共享底层数组
- **Map**：哈希表，负载因子 6.5 或溢出桶多会扩容，并发不安全
- **Channel**：goroutine 通信管道，有缓冲/无缓冲两种
- **Context**：传取消信号、超时、元数据
- **锁**：Mutex、RWMutex、Once、atomic 按场景选
- **并发模型**：CSP，通过通信共享内存
- **并发模式**：生产者-消费者、Worker Pool、Pipeline、Fan-Out/Fan-In、Context 取消、读写锁分离
- **高并发**：限流、worker pool、atomic、对象池、熔断
- **高可用**：graceful shutdown、无状态、负载均衡、监控告警

> **面试口诀：GMP 调度万级并发，channel 通信不要共享内存，context 管超时和取消，slice/map 注意底层共享**

> **Channel 底层口诀：一把锁、一个环形 buf、两个等待队列，发满挂 sendq，读空挂 recvq**

> **Select 底层口诀：多个 case 先排序，随机选择防饥饿，没 default 就挂多个队列等唤醒**

> **并发模式口诀：生产消费解耦忙，worker pool 控并发，pipeline 分阶段，fan-out/fan-in 并行算，context 一把取消杀**

> **高并发口诀：限流削峰异步化，缓存池化无状态，熔断降级超时控，监控告警保平安**
