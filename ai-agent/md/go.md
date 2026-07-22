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

### 5. 项目应用

```go
// 控制最大并发数
func worker(jobs <-chan int, wg *sync.WaitGroup) {
    defer wg.Done()
    for j := range jobs {
        fmt.Println("处理任务", j)
    }
}

func main() {
    jobs := make(chan int, 100)
    var wg sync.WaitGroup

    // 启动 10 个 worker
    for i := 0; i < 10; i++ {
        wg.Add(1)
        go worker(jobs, &wg)
    }

    for i := 0; i < 100; i++ {
        jobs <- i
    }
    close(jobs)
    wg.Wait()
}
```

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
var counter int
var mu sync.Mutex

func add() {
    mu.Lock()
    counter++
    mu.Unlock()
}
```

```go
// 读多写少用 RWMutex
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

## 十四、一句话总结

- **GMP**：G 是任务，M 是线程，P 是调度器，work stealing 实现高并发
- **GC**：三色标记 + 写屏障，Go 1.8 后 STW 很短
- **Slice**：底层数组 + len/cap，append 可能扩容，注意共享底层数组
- **Map**：哈希表，负载因子 6.5 或溢出桶多会扩容，并发不安全
- **Channel**：goroutine 通信管道，有缓冲/无缓冲两种
- **Context**：传取消信号、超时、元数据
- **锁**：Mutex、RWMutex、Once、atomic 按场景选
- **微服务**：gin/goframe 做业务，go-zero/gRPC 做 RPC，etcd 做注册发现，k8s 做部署调度

> **面试口诀：GMP 调度万级并发，channel 通信不要共享内存，context 管超时和取消，slice/map 注意底层共享**
