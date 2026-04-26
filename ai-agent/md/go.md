# Go语言面试八股文

## 基础语法

### 1. Go语言的特点
- 简洁高效，编译型语言
- 天生支持并发（goroutine + channel）
- 内置垃圾回收机制
- 丰富的标准库
- 跨平台编译
- 静态类型，类型安全

### 2. 什么是goroutine？它和线程的区别？
- goroutine是Go语言的轻量级线程
- 栈空间初始很小（2KB），可动态增长
- 由Go运行时调度，不是操作系统调度
- 上下文切换开销小
- 一个程序可以同时运行成千上万个goroutine

### 3. Channel的作用和类型
- 用于goroutine之间的通信和同步
- 无缓冲channel：发送和接收同时准备好才完成
- 有缓冲channel：有容量限制，满了才会阻塞
- 单向channel：只用于发送或只用于接收

### 4. defer的执行顺序
- 多个defer按后进先出（LIFO）顺序执行
- defer在return语句之后、函数返回之前执行
- defer常用于资源清理（关闭文件、释放锁等）

### 5. Go的内存模型
- 栈内存：由编译器自动分配释放
- 堆内存：由垃圾回收器管理
- 逃逸分析：编译器决定变量分配在栈还是堆

## 并发编程

### 6. sync包的常用组件
- sync.Mutex：互斥锁
- sync.RWMutex：读写锁
- sync.WaitGroup：等待一组goroutine完成
- sync.Once：只执行一次
- sync.Cond：条件变量
- sync.Map：并发安全的Map

### 7. 互斥锁和读写锁的区别
- Mutex：同一时间只能有一个goroutine持有锁
- RWMutex：读锁可多个同时持有，写锁只能一个持有，且读写互斥

### 8. select的使用
- 用于同时处理多个channel
- 多个case都准备好时随机选择一个
- 没有准备好时会阻塞
- default分支可避免阻塞

### 9. context的作用和使用场景
- 用于传递请求上下文
- 控制goroutine的取消和超时
- 传递请求范围内的键值对
- 常用方法：WithCancel、WithDeadline、WithTimeout、WithValue

### 10. 如何避免竞态条件？
- 使用channel代替共享内存
- 使用sync包的锁机制
- 使用sync.Map代替普通map
- 使用atomic包进行原子操作

## 项目经验相关

### 11. 在GPU-Mall项目中如何使用Prometheus？
- 集成Prometheus客户端库
- 自定义指标（Counter、Gauge、Histogram等）
- 暴露/metrics接口供Prometheus采集
- 查询PromQL获取GPU训练工况数据
- 实现任务运行状态监控和告警

### 12. 任务编排系统的设计思路
- 优先级队列：根据电价和任务紧急程度排序
- 依赖管理：任务之间的依赖关系图
- 调度算法：最优资源匹配、电价敏感调度
- 状态机：任务生命周期管理（等待、运行、完成、失败）
- 容错机制：任务重试、失败回滚

### 13. 高并发场景下Redis的使用
- 使用Redis+Lua脚本保证原子性
- 分布式锁防止并发问题
- 消息队列实现异步处理
- 缓存热点数据
- 限流器控制访问频率

### 14. Go项目的性能优化经验
- pprof进行性能分析
- 减少内存分配（对象复用、sync.Pool）
- 优化并发控制，避免锁竞争
- 使用缓冲channel提高吞吐量
- 合理设置GOMAXPROCS

### 15. 数据库适配（如MySQL到达梦）的经验
- 使用ORM抽象层（如GORM）
- 统一SQL语法，避免数据库特有语法
- 驱动层适配
- 数据类型映射处理
- 测试用例覆盖验证
