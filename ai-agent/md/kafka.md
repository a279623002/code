# Kafka 面试笔记

> Kafka 是高性能分布式流处理平台，常用于日志收集、消息系统、流式数据处理。调度系统中可用于训练任务状态流转、指标采集等场景。

---

## 一、Kafka 是什么？

**一句话**：Kafka 是一个分布式、高吞吐、可持久化的**发布/订阅消息队列**，同时也是一个流处理平台。

**核心特点**：
- 高吞吐：单机每秒百万级消息
- 持久化：消息落盘，可重复消费
- 高可用：多副本机制
- 水平扩展：增加 Broker 即可扩容
- 流处理：Kafka Streams / Kafka Connect

---

## 二、架构图

```
┌───────────────────────────────────────────────┐
│                   Producer                     │
│              （生产者：应用服务）                │
└───────────────────┬───────────────────────────┘
                    │ send()
                    ↓
┌───────────────────────────────────────────────┐
│              Kafka Cluster                     │
│  ┌─────────┐  ┌─────────┐  ┌─────────┐       │
│  │ Broker 1│  │ Broker 2│  │ Broker 3│       │
│  │ [Partition 0]│ [Partition 1]│ [Partition 2]│
│  │ Leader  │  │ Leader  │  │ Leader  │       │
│  │ Replica │  │ Replica │  │ Replica │       │
│  └─────────┘  └─────────┘  └─────────┘       │
│        ↑              ↑             ↑         │
│        └──────────────┼─────────────┘         │
│                       │                        │
│                   ZooKeeper / KRaft            │
└───────────────────────┬───────────────────────┘
                        │ pull()
┌───────────────────────┴───────────────────────┐
│                   Consumer                     │
│              （消费者：应用服务）                │
└───────────────────────────────────────────────┘
```

---

## 三、核心概念

| 概念 | 说明 |
|---|---|
| **Producer** | 消息生产者 |
| **Consumer** | 消息消费者 |
| **Broker** | Kafka 服务器节点 |
| **Topic** | 消息主题，逻辑分类 |
| **Partition** | Topic 的分区，物理存储单位 |
| **Offset** | 消息在分区内的唯一标识 |
| **Replica** | 分区副本，保证高可用 |
| **Leader/Follower** | 主副本负责读写，从副本同步 |
| **Consumer Group** | 消费者组，组内消费者共同消费一个 Topic |
| **Coordinator** | 负责消费者组的分区分配和 Rebalance |

---

## 四、Topic 与 Partition

```
Topic: order
├─ Partition 0 → [msg0, msg1, msg2, ...]  Offset 0,1,2...
├─ Partition 1 → [msg0, msg1, msg2, ...]
└─ Partition 2 → [msg0, msg1, msg2, ...]
```

**为什么分区？**
- 提高并发：多个 Partition 可并行读写
- 提高吞吐：分散到不同 Broker
- 提高存储：单台机器存不下可分多台

**分区策略**：
| 策略 | 说明 |
|---|---|
| RoundRobin | 轮询，均匀分配 |
| Key Hash | 按 key 的 hash 值决定分区，保证同 key 消息进同一分区 |
| 自定义 | 实现 Partitioner 接口 |

---

## 五、工作模式

### 1. 点对点 vs 发布订阅

| 模式 | 特点 |
|---|---|
| **点对点** | 一条消息只被一个消费者消费 |
| **发布/订阅** | 一条消息可被多个消费者/消费者组消费 |

Kafka 通过 **Consumer Group** 实现两种模式：
- 一个组内：点对点（消息被组内某个消费者消费）
- 多个组之间：发布/订阅（每个组都能收到全量消息）

### 2. Consumer Group Rebalance

当消费者组成员变化时（加入、退出、崩溃），Coordinator 会重新分配 Partition。

```
分区分配策略：
- Range（默认）：按 Topic 范围分配
- RoundRobin：轮询分配
- Sticky：尽量保持上次分配，减少变动
```

---

## 六、消息幂等性

### 1. 生产者幂等

```java
props.put("enable.idempotence", "true");  // 开启幂等
props.put("acks", "all");
props.put("retries", Integer.MAX_VALUE);
```

**原理**：
- Producer 初始化时分配 PID（Producer ID）
- 每个消息带 `<PID, Sequence Number, Partition>`
- Broker 去重：同一个 PID + Sequence Number 重复提交只保留一条

> 局限：幂等只保证**单分区、单会话内**的 exactly once。

### 2. 事务实现跨分区幂等

```java
producer.initTransactions();
try {
    producer.beginTransaction();
    producer.send(record1);
    producer.send(record2);
    producer.commitTransaction();
} catch (Exception e) {
    producer.abortTransaction();
}
```

### 3. 消费者幂等

- 业务层去重：消息唯一 ID + Redis/DB 去重表
- 幂等设计：数据库唯一索引、状态机校验

---

## 七、消息可靠性

### 1. acks 参数

| acks | 含义 | 可靠性 |
|---|---|---|
| `0` | 发送不等确认，可能丢 | 最低 |
| `1` | 等 Leader 确认 | 中等 |
| `all` | 等所有 ISR 副本确认 | 最高 |

### 2. ISR（In-Sync Replicas）

**一句话**：和 Leader 保持同步的副本集合。

```
副本集合：
- AR（Assigned Replicas）：所有副本
- ISR：同步中的副本
- OSR：掉队的副本

只有 ISR 中的副本才有资格成为新 Leader
```

### 3. min.insync.replicas

```
acks=all + min.insync.replicas=2
表示：至少要有 2 个 ISR 副本确认，才认为写入成功
```

### 4. 消费者提交 Offset

| 提交方式 | 特点 |
|---|---|
| 自动提交 | 简单，可能丢消息或重复 |
| 手动同步提交 | 处理完再提交，可靠性高 |
| 手动异步提交 | 性能高，可能丢 offset |

**推荐**：手动提交 + 业务处理幂等。

---

## 八、消息顺序性

### 1. Kafka 天然保证

- **同一个 Partition 内**：消息按发送顺序存储，消费者按顺序读取
- **不同 Partition 之间**：不保证顺序

### 2. 保证全局顺序

**方法一**：一个 Topic 只设 1 个 Partition
- 简单，但牺牲吞吐量

**方法二**：按业务 key 分区
- 同 key 的消息进入同一 Partition
- 例如：同一订单 ID 的所有消息按 `order_id` hash 进同一分区

```java
producer.send(new ProducerRecord<>("order", orderId, msg));
```

### 3. 乱序原因

- 多 Partition 天然乱序
- 生产者重试：消息 A 发送失败重试，消息 B 先发成功，导致 A 在 B 之后
- 多消费者并发处理

**解决**：设置 `max.in.flight.requests.per.connection=1`，避免重试导致乱序。

---

## 九、Kafka 高性能原因

1. **顺序写磁盘**：比随机写快很多
2. **零拷贝（Zero Copy）**：sendfile 系统调用，减少数据拷贝
3. **批量处理**：消息批量压缩发送
4. **分区并行**：多 Partition 多消费者并发
5. **页缓存**：利用 OS 缓存，减少磁盘 IO

---

## 十、常用命令

```bash
# 创建 Topic
kafka-topics.sh --create --topic order --partitions 3 --replication-factor 2 --bootstrap-server localhost:9092

# 查看 Topic 列表
kafka-topics.sh --list --bootstrap-server localhost:9092

# 发送消息
kafka-console-producer.sh --topic order --bootstrap-server localhost:9092

# 消费消息
kafka-console-consumer.sh --topic order --from-beginning --bootstrap-server localhost:9092

# 查看消费者组
kafka-consumer-groups.sh --bootstrap-server localhost:9092 --describe --group g1
```

---

## 十一、面试高频问题

### Q1：Kafka 和 RabbitMQ 的区别？

| 特性 | Kafka | RabbitMQ |
|---|---|---|
| 设计目标 | 高吞吐流处理 | 通用消息队列 |
| 消息持久化 | 默认持久化，可重复读 | 默认内存，可配置持久化 |
| 吞吐量 | 更高（百万级/秒） | 万级/秒 |
| 消费模式 | pull | push |
| 顺序性 | 单分区有序 | 队列内有序 |
| 适用场景 | 日志、大数据、流处理 | 企业级异步任务、RPC |

### Q2：Kafka 为什么快？

**答**：
1. 顺序写磁盘
2. 零拷贝
3. 批量和压缩
4. 分区并行
5. 页缓存

### Q3：怎么保证消息不丢失？

**答**：
- Producer：`acks=all`，开启重试
- Broker：`replication.factor >= 2`，`min.insync.replicas >= 2`
- Consumer：手动提交 offset，处理完再提交

### Q4：怎么保证消息顺序？

**答**：
- 单 Partition 天然有序
- 多 Partition 时按业务 key 分区
- 设置 `max.in.flight.requests.per.connection=1` 避免重试乱序

### Q5：Rebalance 是什么？怎么避免频繁 Rebalance？

**答**：
- 消费者组成员变化时重新分配分区
- 避免：合理设置 session.timeout.ms、heartbeat.interval.ms，避免消费者处理过久

---

## 十二、一句话总结

- **Kafka**：高吞吐分布式消息流平台
- **Topic/Partition**：逻辑主题 + 物理分区，Partition 是并行和顺序的单位
- **幂等**：生产者 `enable.idempotence=true`，消费者业务去重
- **可靠**：`acks=all` + 多副本 + ISR + 手动提交 offset
- **顺序**：单分区有序，多分区按 key 分区
- **高性能**：顺序写 + 零拷贝 + 批量 + 页缓存
