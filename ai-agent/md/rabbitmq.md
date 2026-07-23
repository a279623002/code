# RabbitMQ 面试笔记

> RabbitMQ 是经典的企业级消息队列，支持多种工作模式，常用于异步任务、削峰填谷、服务解耦。调度系统中可用于训练任务状态通知、算力告警等场景。

---

## 一、RabbitMQ 是什么？

**一句话**：RabbitMQ 是一个开源的、基于 AMQP 协议的**消息中间件**，支持发布/订阅、路由、主题等多种消息模式。

**核心特点**：
- 基于 Erlang 编写，天生高并发
- 支持消息持久化、确认机制
- 支持路由、主题、RPC 等多种模式
- 管理界面友好
- 适合企业级异步任务

---

## 二、架构图

```
┌─────────────────────────────────────────────────────────┐
│                      RabbitMQ Server                     │
│  ┌─────────────────────────────────────────────────────┐ │
│  │                   Exchange（交换机）                  │ │
│  │  ┌────────┐  ┌────────┐  ┌────────┐  ┌────────┐   │ │
│  │  │ direct │  │ fanout │  │ topic  │  │headers │   │ │
│  │  └────┬───┘  └────┬───┘  └────┬───┘  └────┬───┘   │ │
│  │       │           │           │           │        │ │
│  │       └───────────┴─────┬─────┴───────────┘        │ │
│  │                         │                          │ │
│  │                   Binding（绑定）                    │ │
│  │                         │                          │ │
│  │       ┌─────────────────┼─────────────────┐        │ │
│  │       ↓                 ↓                 ↓        │ │
│  │  ┌─────────┐      ┌─────────┐      ┌─────────┐    │ │
│  │  │ Queue A │      │ Queue B │      │ Queue C │    │ │
│  │  │ [msg]   │      │ [msg]   │      │ [msg]   │    │ │
│  │  └────┬────┘      └────┬────┘      └────┬────┘    │ │
│  │       │                │                │         │ │
│  └───────┼────────────────┼────────────────┼─────────┘ │
└──────────┼────────────────┼────────────────┼───────────┘
           │                │                │
           ↓                ↓                ↓
      Consumer 1      Consumer 2      Consumer 3
```

---

## 三、核心概念

| 概念 | 说明 |
|---|---|
| **Producer** | 消息生产者 |
| **Consumer** | 消息消费者 |
| **Queue** | 消息队列，存储消息 |
| **Exchange** | 交换机，接收生产者消息并路由到 Queue |
| **Binding** | Exchange 和 Queue 的绑定关系 |
| **Routing Key** | 路由键，Exchange 根据它决定消息去向 |
| **Channel** | 轻量级连接，建立在真实 TCP 连接上 |
| **VHost** | 虚拟主机，逻辑隔离 |

---

## 四、Exchange 类型

| 类型 | 路由规则 | 场景 |
|---|---|---|
| **direct** | routing key 完全匹配 binding key | 点对点精确路由 |
| **fanout** | 广播到所有绑定队列 | 公告、广播 |
| **topic** | routing key 通配符匹配 `*`、`#` | 日志分级、事件总线 |
| **headers** | 根据消息 headers 匹配 | 复杂条件路由 |

### topic 通配符

| 符号 | 含义 |
|---|---|
| `*` | 匹配一个单词 |
| `#` | 匹配零个或多个单词 |

```
路由键：log.order.error
绑定键：log.order.*   → 匹配
绑定键：log.#         → 匹配
绑定键：log.*.*        → 匹配
绑定键：log.user.*     → 不匹配
```

---

## 五、工作模式

### 1. 简单模式（Hello World）

```
Producer → Queue → Consumer
```

### 2. Work Queues（工作队列）

```
Producer → Queue → Consumer 1
               → Consumer 2
               → Consumer 3
```

- 多个消费者竞争消费
- 默认轮询，可设置公平分发 `basicQos(1)`

### 3. Publish/Subscribe（发布订阅）

```
Producer → Fanout Exchange → Queue A → Consumer A
                           → Queue B → Consumer B
```

### 4. Routing（路由模式）

```
Producer → Direct Exchange ──[error]──→ Queue A
                         ──[info]───→ Queue B
```

### 5. Topics（主题模式）

```
Producer → Topic Exchange ──[log.*.error]──→ Queue A
                         ──[log.#]────────→ Queue B
```

### 6. RPC（远程调用）

通过两个 Queue 实现请求/响应：
- `rpc_queue`：发送请求
- 临时 Queue：接收响应

---

## 六、消息幂等性

### 1. 为什么需要幂等？

RabbitMQ 支持消息确认和重试，网络抖动时可能重复投递。

### 2. 业务层幂等

```go
// 消息体带唯一 messageId
type OrderMsg struct {
    MessageID string `json:"message_id"`
    OrderID   string `json:"order_id"`
    Action    string `json:"action"`
}

func handle(msg OrderMsg) {
    // 1. 去重
    ok, _ := rdb.SetNX(ctx, "mq:"+msg.MessageID, "1", 24*time.Hour).Result()
    if !ok {
        return // 已处理
    }
    // 2. 处理业务
    processOrder(msg.OrderID, msg.Action)
}
```

### 3. 数据库唯一索引兜底

```sql
CREATE TABLE order_log (
    message_id VARCHAR(64) PRIMARY KEY,
    order_id   VARCHAR(64),
    action     VARCHAR(32)
);
```

---

## 七、消息可靠性

### 1. 生产者确认（Publisher Confirm）

```go
ch.Confirm(false)  // 开启确认模式
ack, nack := ch.NotifyPublish(make(chan amqp.Confirmation, 1))

err := ch.Publish("", "queue", false, false, msg)
if err != nil {
    // 发送失败处理
}

select {
case confirm := <-ack:
    if confirm.Ack {
        // Broker 已接收
    }
case <-nack:
    // Broker 拒绝，需要重试
}
```

### 2. 消息持久化

```go
// Queue 持久化
q, _ := ch.QueueDeclare("task", true, false, false, false, nil)

// Message 持久化
ch.Publish("", q.Name, false, false, amqp.Publishing{
    DeliveryMode: amqp.Persistent,  // 持久化
    Body:         []byte(body),
})
```

### 3. 消费者确认（Consumer Ack）

```go
// 手动确认
msgs, _ := ch.Consume(q.Name, "", false, false, false, false, nil)
for d := range msgs {
    process(d.Body)
    d.Ack(false)  // 处理完再确认
}
```

> 自动确认模式下，消息一出队列就确认，可能丢失。

---

## 八、消息顺序性

### 1. 单队列单消费者

RabbitMQ 保证**同一个队列、同一个消费者**内消息按发送顺序处理。

### 2. 多消费者乱序

多个消费者同时处理时，完成顺序可能不同。

**解决方法**：
- 一个队列只配一个消费者
- 按业务 key 分多个队列，每个队列一个消费者
- 业务层加 seq 号，消费端排序

### 3. 分区顺序

```
订单 1 的消息 → Queue 1 → Consumer 1
订单 2 的消息 → Queue 2 → Consumer 2
订单 3 的消息 → Queue 3 → Consumer 3
```

---

## 九、常用命令

```bash
# 启动服务
sudo systemctl start rabbitmq-server

# 查看状态
sudo rabbitmqctl status

# 列出队列
sudo rabbitmqctl list_queues

# 列出交换机
sudo rabbitmqctl list_exchanges

# 创建用户
sudo rabbitmqctl add_user admin 123456
sudo rabbitmqctl set_user_tags admin administrator
sudo rabbitmqctl set_permissions -p / admin ".*" ".*" ".*"
```

---

## 十、面试高频问题

### Q1：RabbitMQ 和 Kafka 的区别？

| 特性 | RabbitMQ | Kafka |
|---|---|---|
| 协议 | AMQP | 自定义协议 |
| 设计目标 | 通用消息队列 | 高吞吐流处理 |
| 吞吐量 | 万级/秒 | 百万级/秒 |
| 消费模式 | push | pull |
| 路由能力 | 强（direct/topic/headers） | 弱（按 topic/partition） |
| 消息持久化 | 可选 | 默认 |
| 顺序性 | 单队列有序 | 单分区有序 |

### Q2：Exchange 有哪些类型？

**答**：direct、fanout、topic、headers。

### Q3：怎么保证消息不丢失？

**答**：
- Producer：开启 confirm 机制
- Broker：Queue 和 Message 持久化
- Consumer：手动 ack，处理完再确认

### Q4：怎么保证消息不被重复消费？

**答**：
- 消息带唯一 ID
- 消费端 Redis/DB 去重
- 数据库唯一索引兜底

### Q5：消息积压怎么处理？

**答**：
1. 增加消费者，临时扩容
2. 优化消费者处理逻辑
3. 设置消息 TTL，丢弃过期消息
4. 转移至新的队列异步处理

---

## 十一、一句话总结

- **RabbitMQ**：基于 AMQP 的企业级消息队列
- **核心组件**：Producer、Exchange、Queue、Consumer
- **Exchange 类型**：direct、fanout、topic、headers
- **可靠**：Publisher Confirm + 持久化 + Consumer Ack
- **幂等**：消息唯一 ID + 去重表
- **顺序**：单队列单消费者有序，多消费者需业务层控制
