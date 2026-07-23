# MQTT 面试笔记

> 项目背景：机器人管控平台通过 MQTT 与机器人通信，包括订阅话题（电量、状态、心跳）和控制话题（行走、表情、动作）。MQTT 是 IoT 场景最常用的轻量级消息协议。

---

## 一、MQTT 是什么？

**一句话**：MQTT（Message Queuing Telemetry Transport）是一种基于**发布/订阅**模式的轻量级消息协议，专为低带宽、不稳定网络设计。

**特点**：
- 轻量：头部最小 2 字节
- 低功耗：适合 IoT 设备
- 发布/订阅：解耦生产者和消费者
- 三种 QoS：满足不同可靠性需求
- 支持遗嘱消息、 retained 消息

---

## 二、架构图

```
        ┌─────────────────┐
        │   MQTT Broker   │  ← 消息中转站
        │   （服务器）     │
        └────────┬────────┘
                 │
    ┌────────────┼────────────┐
    │            │            │
    ↓            ↓            ↓
┌───────┐   ┌───────┐   ┌───────┐
│Publisher│   │Publisher│   │Subscriber│
│ 传感器  │   │  手机   │   │  后台系统 │
└───────┘   └───────┘   └───────┘
    ↑                        ↑
    │      发布 / 订阅        │
    │   topic: robot/1001/battery   │
    └────────────────────────┘
```

**核心角色**：
| 角色 | 作用 |
|---|---|
| **Publisher** | 发布消息到某个 topic |
| **Subscriber** | 订阅某个 topic，接收消息 |
| **Broker** | 消息中转服务器，负责转发 |
| **Topic** | 消息主题，用 `/` 分层，如 `home/living/temperature` |

---

## 三、协议原理

### 1. 连接过程

```
客户端 ───── CONNECT ─────→ Broker
        client_id, username, password, keepalive, will

客户端 ←──── CONNACK ────── Broker
        返回连接结果（0 成功，其他错误码）
```

### 2. 消息流程

```
发布：
客户端 A ───── PUBLISH(topic, payload, qos) ─────→ Broker
Broker 根据 topic 匹配订阅者，转发消息
Broker ───── PUBLISH ─────→ 客户端 B、C

订阅：
客户端 B ───── SUBSCRIBE(topic, qos) ─────→ Broker
Broker ←──── SUBACK ─────── 客户端 B
```

### 3. 主题通配符

| 通配符 | 含义 | 示例 |
|---|---|---|
| `+` | 单层匹配 | `home/+/temperature` 匹配 `home/living/temperature` |
| `#` | 多层匹配 | `home/#` 匹配 `home/living/temperature`、`home/garage/door` |

> ⚠️ 发布消息时不能用通配符，订阅时才能用。

---

## 四、QoS（服务质量等级）

| QoS | 名称 | 说明 | 场景 |
|---|---|---|---|
| **0** | 最多一次 | 发一次，不确认，可能丢失 | 高频 Telemetry，允许丢包 |
| **1** | 至少一次 | 确保送达，可能重复 | 一般控制命令 |
| **2** |  exactly一次 | 四次握手，确保只收一次 | 关键指令，如支付、门锁 |

### QoS 1 流程

```
客户端 ───── PUBLISH ─────→ Broker
客户端 ←──── PUBACK ─────── Broker
```

### QoS 2 流程

```
客户端 ───── PUBLISH ─────→ Broker
客户端 ←──── PUBREC ─────── Broker
客户端 ───── PUBREL ─────→ Broker
客户端 ←──── PUBCOMP ─────── Broker
```

---

## 五、消息幂等性

### 1. 什么是幂等？

**一句话**：同样的消息收多次，结果和收一次一样。

### 2. MQTT 幂等问题

- **QoS 1**：保证"至少一次"，可能重复发送
- **QoS 2**：保证"exactly once"，但性能开销大

### 3. 业务层幂等设计

```go
// 消息带唯一 messageId
const msgKey = "mqtt:handled:" + msg.MessageID

ok, _ := rdb.SetNX(ctx, msgKey, "1", 24*time.Hour).Result()
if !ok {
    return // 已经处理过，直接丢弃
}
process(msg.Payload)
```

**常见做法**：
- 消息带唯一 ID
- 消费端用 Redis/DB 做去重表
- 关键业务用数据库唯一索引兜底

---

## 六、消息可靠性

### 1. Broker 层面

- **持久化会话（Clean Session = false）**：客户端断线后，Broker 保存订阅和未送达消息
- **Retained Message**：最后一条保留消息，新订阅者立即收到
- **Will Message（遗嘱消息）**：客户端异常断线时，Broker 自动发布遗嘱

### 2. 应用层面

- 合理选择 QoS（关键指令用 QoS 1/2）
- 客户端实现断线重连 + 指数退避
- 重要消息落库后再 ACK

```go
// 断线重连示例（paho mqtt）
opts := MQTT.NewClientOptions()
opts.AddBroker("tcp://broker:1883")
opts.SetAutoReconnect(true)           // 自动重连
opts.SetConnectRetry(true)            // 启动时重试
opts.SetClientID("robot-backend-001")
opts.SetWill("status/backend", "offline", 1, true)  // 遗嘱消息
```

---

## 七、消息顺序性

### 1. MQTT 保证什么？

- **同一个 topic、同一个 QoS、同一条连接**上，消息按发送顺序到达
- 不同 topic 之间不保证顺序
- QoS 升级/混用可能导致乱序

### 2. 顺序性保证方法

1. **统一 topic**：同一设备的消息发到同一个 topic
2. **单线程消费**：一个 topic 只用一个消费者处理
3. **业务序号**：消息里带 seq，消费端按 seq 排序或丢弃乱序

```
机器人上传状态：
robot/1001/status → {"seq":1, "battery":80}
robot/1001/status → {"seq":2, "battery":79}
robot/1001/status → {"seq":3, "battery":78}

消费者按 seq 处理，seq=2 没到就缓存 seq=3
```

---

## 八、项目实战：机器人 MQTT 通信

```
控制话题：
robot/{robot_id}/control/move      → {"action":"forward","speed":1.0}
robot/{robot_id}/control/expression → {"name":"happy"}
robot/{robot_id}/control/action     → {"name":"dance"}

订阅话题：
robot/{robot_id}/status            → {"battery":80,"state":"idle"}
robot/{robot_id}/heartbeat         → {"ts":1690000000}
robot/{robot_id}/action/result     → {"action":"dance","result":"ok"}
```

**为什么用 MQTT？**
- 机器人是移动设备，网络不稳定，MQTT 轻量、支持断线重连
- 发布/订阅天然解耦控制和上报
- QoS 1 保证控制指令可靠到达

---

## 九、面试高频问题

### Q1：MQTT 和 HTTP 的区别？

**答**：
- MQTT 是发布/订阅模式，HTTP 是请求/响应模式
- MQTT 长连接、低功耗，HTTP 短连接、开销大
- MQTT 支持实时推送，HTTP 需要轮询
- MQTT 适合 IoT，HTTP 适合 Web/API

### Q2：QoS 0/1/2 怎么选？

**答**：
- 传感器数据：QoS 0（允许丢）
- 普通控制命令：QoS 1（可靠，可能重复）
- 关键控制/支付：QoS 2（exactly once）

### Q3：MQTT Broker 有哪些？

**答**：
- **Eclipse Mosquitto**：轻量，适合嵌入式
- **EMQX**：企业级，支持百万连接、规则引擎
- **HiveMQ**：商业 Broker
- **RabbitMQ**：通过插件支持 MQTT

### Q4：如何处理 MQTT 消息重复？

**答**：
1. 优先用 QoS 2（但性能差）
2. 业务层做幂等：消息唯一 ID + 去重表
3. 关键操作用数据库唯一索引兜底

### Q5：MQTT 消息顺序怎么保证？

**答**：
- 同一 topic、同一连接、同一 QoS 下天然有序
- 跨 topic 不保证
- 需要严格顺序时，业务上加 seq 号，消费端排序或缓存

---

## 十、一句话总结

- **MQTT**：轻量级发布/订阅协议，IoT 首选
- **QoS**：0 最多一次、1 至少一次、2 exactly once
- **幂等**：QoS 1 可能重复，业务层用唯一 ID 去重
- **可靠**：持久化会话 + 遗嘱消息 + 断线重连
- **顺序**：同 topic 同 QoS 同连接有序，否则业务加 seq
