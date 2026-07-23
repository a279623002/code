# 微服务架构面试笔记

> 项目背景：调度系统拆分为 API 网关、算力资源服务、任务调度服务，通过 gRPC/HTTP 通信，etcd 做服务发现，k8s 做部署。微服务架构是面试核心考点。

---

## 一、CAP 定理

```
         ┌─────────────────┐
         │    Consistency  │ 一致性：所有节点同时看到相同数据
         │    一致性        │
         └────────┬────────┘
                  │
    ┌─────────────┼─────────────┐
    │             │             │
    ↓             ↓             ↓
Partition   Availability   只能三选二
Tolerance   可用性         （网络分区时）
分区容错性
```

| 组合 | 代表 | 特点 |
|---|---|---|
| **CP** | ZooKeeper、etcd、HBase | 保证一致性，牺牲部分可用性 |
| **AP** | Eureka、Cassandra | 保证可用性，牺牲强一致性 |
| **CA** | 单机 MySQL | 不考虑分区，理论上存在 |

**面试重点**：网络分区不可避免，所以 **P 必须满足**，实际只有 CP 和 AP 两种选择。

---

## 二、protobuf

### 1. 是什么？

**一句话**：Protocol Buffers 是 Google 出的二进制序列化协议，用于结构化数据的高效传输和存储。

### 2. 为什么快？

| 特性 | 说明 |
|---|---|
| 二进制 | 比 JSON/XML 体积小 3-10 倍 |
| 强类型 | 编译期检查 |
| 向前/向后兼容 | 可加字段不影响旧版本 |
| 解析快 | 无需像 JSON 一样字符串解析 |

### 3. 示例

```protobuf
syntax = "proto3";

message User {
    int64 id = 1;
    string name = 2;
    int32 age = 3;
}

service UserService {
    rpc GetUser (GetUserReq) returns (User);
}

message GetUserReq {
    int64 id = 1;
}
```

### 4. proto 编译

```bash
protoc --go_out=. --go-grpc_out=. user.proto
```

---

## 三、gRPC vs HTTP/RESTful

| 特性 | gRPC | HTTP/RESTful |
|---|---|---|
| 协议 | HTTP/2 + protobuf | HTTP/1.1 + JSON |
| 性能 | 高（二进制、多路复用） | 低（文本、头部大） |
| 接口定义 | proto 文件严格约束 | 口头/文档约定 |
| 自动生成 | 客户端/服务端代码自动生成 | 需手写 |
| 流式 | 支持双向流 | 需用 WebSocket/SSE |
| 调试 | 需工具（grpcurl） | 浏览器/curl 直接调 |
| 浏览器支持 | 需 grpc-web 转 | 原生支持 |

**选型**：
- 内部微服务通信：gRPC
- 对外 Open API：RESTful
- 需要浏览器直接访问：RESTful 或 grpc-web

---

## 四、gRPC 通信方式

### 1. 四种模式

| 模式 | 说明 | 场景 |
|---|---|---|
| **Unary RPC** | 一元：请求 → 响应 | 普通接口调用 |
| **Server Streaming** | 服务端流：一个请求，多个响应 | 大文件下载、实时推送 |
| **Client Streaming** | 客户端流：多个请求，一个响应 | 大文件上传、批量提交 |
| **Bidirectional Streaming** | 双向流：两端可同时发 | 实时聊天、在线游戏 |

### 2. 示例

```protobuf
service ChatService {
    // 一元
    rpc SendMessage (Message) returns (Message);

    // 服务端流
    rpc GetHistory (HistoryReq) returns (stream Message);

    // 客户端流
    rpc UploadFiles (stream FileChunk) returns (UploadResult);

    // 双向流
    rpc ChatStream (stream Message) returns (stream Message);
}
```

---

## 五、HTTP/2 多路复用

### 1. HTTP/1.1 的问题

```
HTTP/1.1：
    请求 1 ──→ 响应 1
    请求 2 ──→ 响应 2
    请求 3 ──→ 响应 3
    队头阻塞：一个请求阻塞，后续都要等

HTTP/2：
    一个 TCP 连接上并行多个 Stream
    Stream 1 ──┐
    Stream 2 ──┼──→ 帧（Frame）交织发送
    Stream 3 ──┘
```

### 2. HTTP/2 核心特性

- **二进制分帧**：数据分成小帧传输
- **多路复用**：一个连接多个 Stream，互不阻塞
- **头部压缩（HPACK）**：减少重复头部
- **服务器推送**：服务端主动推送资源

---

## 六、容错机制

### 1. 服务雪崩

```
服务 A ──→ 服务 B ──→ 服务 C
              ↓
           C 故障
              ↓
    B 线程全部阻塞等待 C
              ↓
    B 也故障，A 也故障
```

### 2. 常见容错策略

| 策略 | 说明 | 实现 |
|---|---|---|
| **超时（Timeout）** | 调用超时不再等待 | 配置 timeout |
| **重试（Retry）** | 失败重试，带退避 | 固定/指数退避 |
| **熔断（Circuit Breaker）** | 失败率达到阈值直接失败 | go-zero、hystrix |
| **降级（Fallback）** | 返回兜底数据 | 静态值/缓存 |
| **限流（Rate Limit）** | 限制请求速率 | 令牌桶、漏桶 |
| **隔离（Bulkhead）** | 资源隔离，防止互相影响 | 线程池隔离 |

### 3. go-zero 熔断示例

```go
import "github.com/zeromicro/go-zero/core/breaker"

bk := breaker.NewBreaker()
err := bk.Do("getUser", func() error {
    _, err := userRpc.GetUser(ctx, &user.GetUserReq{Id: 1})
    return err
})
```

---

## 七、横向扩展与纵向扩展

| 维度 | 横向扩展（Scale Out） | 纵向扩展（Scale Up） |
|---|---|---|
| 方式 | 加机器 | 升级单机配置 |
| 成本 | 相对低，线性扩展 | 成本高，有上限 |
| 可用性 | 更好，单点故障影响小 | 差，单机故障影响大 |
| 复杂度 | 高（分布式协调） | 低 |
| 适用 | 微服务、大数据 | 单体应用、数据库 |

**微服务天然适合横向扩展**。

---

## 八、服务发现与注册

```
服务启动：
    User-Service ──register──► etcd
                                    │
                                    │ watch
                                    ↓
    API-Gateway ◄────fetch──── etcd
```

**常见注册中心**：
| 组件 | 特点 |
|---|---|
| **etcd** | CP，K8s 原生，gRPC 友好 |
| **Consul** | 支持健康检查、多数据中心 |
| **Eureka** | AP，Netflix 出品 |
| **Nacos** | 阿里开源，注册中心 + 配置中心 |

---

## 九、API 网关

### 1. 作用

```
客户端 ──→ API Gateway ──→ 用户服务
                      ──→ 订单服务
                      ──→ 商品服务
```

**功能**：
- 统一入口
- 路由转发
- 鉴权认证
- 限流熔断
- 日志监控
- 协议转换（HTTP ↔ gRPC）

### 2. 常见网关

| 网关 | 特点 |
|---|---|
| **Kong** | 基于 OpenResty，插件丰富 |
| **Nginx** | 高性能反向代理 |
| **Spring Cloud Gateway** | Java 生态 |
| **go-zero API Gateway** | Go 微服务网关 |

---

## 十、面试高频问题

### Q1：什么是 CAP？微服务怎么选？

**答**：CAP 指一致性、可用性、分区容错性，网络分区时只能满足两个。微服务注册中心：
- 需要强一致性选 CP（etcd、ZooKeeper）
- 需要高可用选 AP（Eureka、Nacos）

### Q2：gRPC 为什么比 RESTful 快？

**答**：
- gRPC 用 protobuf 二进制序列化，体积小、解析快
- 基于 HTTP/2 多路复用，减少连接数
- 强类型接口，减少运行时错误

### Q3：熔断和降级的区别？

**答**：
- 熔断：下游故障时，直接切断调用，快速失败
- 降级：服务压力大或异常时，返回兜底数据，保证可用

### Q4：什么是服务网格（Service Mesh）？

**答**：把服务间通信逻辑（路由、熔断、监控、安全）从业务代码中抽离，交给 Sidecar（如 Istio、Envoy）处理。业务代码只关心业务。

### Q5：微服务拆分原则？

**答**：
- 按业务领域拆分（DDD）
- 高内聚低耦合
- 独立部署、独立扩展
- 避免过细导致调用链复杂

---

## 十一、一句话总结

- **CAP**：网络分区时只能选 CP 或 AP
- **protobuf**：高效二进制序列化协议
- **gRPC**：HTTP/2 + protobuf，适合内部高性能 RPC
- **多路复用**：HTTP/2 一个连接多个 Stream
- **容错**：超时、重试、熔断、降级、限流、隔离
- **扩展**：微服务优先横向扩展
- **服务发现**：etcd/Consul/Nacos 实现服务注册与发现
