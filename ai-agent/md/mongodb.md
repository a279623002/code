# MongoDB 面试笔记

> MongoDB 是文档型 NoSQL 数据库，核心围绕文档模型、副本集、分片集群、聚合管道、索引。

---

## 一、基础概念

### 1. MongoDB 是什么？核心特点？

- 文档型 NoSQL 数据库，基于 BSON 存储
- 无 Schema 约束，灵活的数据模型
- 支持丰富的查询和聚合能力
- 原生分布式：副本集 + 分片集群
- 默认使用 WiredTiger 存储引擎

### 2. 核心术语对照

| MongoDB | 关系型数据库 |
|---|---|
| Database | Database |
| Collection（集合） | Table（表） |
| Document（文档） | Row（行） |
| Field（字段） | Column（列） |
| _id | 主键（默认 ObjectId） |
| Embedded Document | Join 表 |
| Index | Index |

---

## 二、数据模型

### 1. 嵌入式 vs 引用式

**嵌入式（Embedded）**：
```json
{
  "_id": 1,
  "name": "张三",
  "orders": [
    { "order_id": 101, "amount": 99.9 },
    { "order_id": 102, "amount": 199.9 }
  ]
}
```

**引用式（Reference）**：
```json
// user 集合
{ "_id": 1, "name": "张三" }

// order 集合
{ "_id": 101, "user_id": 1, "amount": 99.9 }
```

### 2. 如何选择嵌入/引用？

| 场景 | 推荐 |
|---|---|
| 一对一 / 一对少 | 嵌入式 |
| 一对多（数据不频繁变化） | 嵌入式 |
| 一对多（数据频繁变化） | 引用式 |
| 多对多 | 引用式 |
| 子文档需要独立查询 | 引用式 |
| 文档大小可能超过 16MB | 引用式 |

---

## 三、索引

### 1. 索引类型

| 索引类型 | 说明 |
|---|---|
| 单字段索引 | 对单个字段创建索引 |
| 复合索引 | 多个字段组合，支持前缀查询 |
| 多键索引 | 数组字段，为每个数组元素创建索引 |
| 文本索引 | 全文搜索，支持分词 |
| 地理空间索引 | 2dsphere（球面）、2d（平面） |
| 哈希索引 | 基于字段哈希值，适合分片键 |
| TTL 索引 | 自动过期删除文档 |
| 唯一索引 | 保证字段值唯一 |
| 部分索引 | 只索引符合条件的文档 |
| 稀疏索引 | 只索引包含该字段的文档 |

### 2. 复合索引 ESR 规则

**E**qual → **S**ort → **R**ange

```javascript
// 查询：status = 'active' AND age > 25 ORDER BY create_time DESC
// 最优索引顺序：
db.collection.createIndex({ status: 1, create_time: -1, age: 1 })
// 等值 → 排序 → 范围
```

### 3. 覆盖索引

查询的所有字段都在索引中，不需要回表查文档数据，性能最高。

```javascript
db.collection.createIndex({ name: 1, age: 1 })
db.collection.find({ name: "张三" }, { name: 1, age: 1, _id: 0 })
// 直接从索引返回，不访问文档
```

### 4. explain() 分析执行计划

```javascript
db.collection.find({ name: "张三" }).explain("executionStats")
```

关键指标：
- **IXSCAN**：索引扫描（好）
- **COLLSCAN**：全表扫描（差）
- **totalDocsExamined**：扫描文档数
- **nReturned**：返回文档数（理想情况两者相等）

---

## 四、副本集

### 1. 副本集架构

```
┌──────────────┐
│   Primary    │ ← 写入
└──────┬───────┘
   ┌───┴───┐
┌──┴──┐ ┌──┴──┐
│Sec1 │ │Sec2 │  ← 读取（可配）
└─────┘ └─────┘
```

- 一个 Primary，多个 Secondary
- 数据通过 oplog 异步同步
- 主节点故障时自动选举新 Primary

### 2. 选举机制（Raft）

- 每个节点有投票权，最多 7 个投票节点
- 获得超过半数投票的节点成为 Primary
- 触发选举条件：主节点心跳超时（默认 10s）
- 选举期间集群不可写入（通常 2-12 秒）

### 3. 写关注（Write Concern）

```javascript
{ writeConcern: { w: "majority", j: true, wtimeout: 5000 } }
```

| 级别 | 说明 |
|---|---|
| w: 0 | 不确认 |
| w: 1 | 主节点确认 |
| w: "majority" | 多数节点确认 |
| w: <number> | 指定数量节点确认 |
| j: true | 写入 journal 日志 |

### 4. 读关注（Read Concern）

| 级别 | 说明 |
|---|---|
| local | 默认，返回节点最新数据 |
| available | 非事务场景可用 |
| majority | 多数节点已确认的数据 |
| linearizable | 线性一致性（最严格） |
| snapshot | 快照读（事务中） |

### 5. 读偏好（Read Preference）

| 模式 | 说明 |
|---|---|
| primary | 默认，只读主节点 |
| primaryPreferred | 优先主节点，主不可用时读从 |
| secondary | 只读从节点 |
| secondaryPreferred | 优先从节点，无可用从时读主 |
| nearest | 读延迟最低的节点 |

---

## 五、分片集群

### 1. 分片集群架构

```
┌──────────┐  ┌──────────┐
│ Mongos   │  │ Mongos   │  ← 路由（查询路由）
└────┬─────┘  └────┬─────┘
     └──────┬──────┘
    ┌───────┴───────┐
    │ Config Server │  ← 配置服务器（元数据）
    └───────────────┘
    ┌───┐ ┌───┐ ┌───┐
    │S1 │ │S2 │ │S3 │  ← 分片（每个分片是一个副本集）
    └───┘ └───┘ └───┘
```

### 2. 分片键选择

**好的分片键**：
- 高基数（大量不同值）
- 写分布均匀（避免热点）
- 查询定向到单个分片

**常见分片策略**：
- **哈希分片**：分布均匀，范围查询需跨分片
- **范围分片**：范围查询高效，可能产生热点

### 3. Chunk 与均衡

- Chunk 是数据迁移的最小单元（默认 64MB）
- Balancer 自动迁移 Chunk 保持数据均衡
- 避免在业务高峰期运行 Balancer

---

## 六、聚合管道

### 1. 常用聚合阶段

```javascript
db.orders.aggregate([
  { $match: { status: "completed" } },        // 过滤
  { $group: { _id: "$user_id", total: { $sum: "$amount" } } },  // 分组
  { $sort: { total: -1 } },                   // 排序
  { $limit: 10 },                             // 限制
  { $project: { user_id: "$_id", total: 1, _id: 0 } },  // 投影
  { $lookup: {                                // 关联查询
      from: "users",
      localField: "user_id",
      foreignField: "_id",
      as: "user"
  }},
  { $unwind: "$user" },                       // 展开数组
  { $addFields: { discount: 0.1 } }           // 添加字段
])
```

### 2. MapReduce vs 聚合管道

- 聚合管道性能更好，使用 C++ 实现
- MapReduce 用 JavaScript，更灵活但更慢
- 优先使用聚合管道

---

## 七、事务

### 1. MongoDB 事务支持

- 4.0+ 支持副本集多文档事务
- 4.2+ 支持分片集群多文档事务
- 基于 WiredTiger 的 MVCC 实现
- 支持 ACID，但性能有损耗

### 2. 事务使用示例

```javascript
const session = client.startSession();
session.startTransaction();
try {
  db.accounts.updateOne({ _id: 1 }, { $inc: { balance: -100 } }, { session });
  db.accounts.updateOne({ _id: 2 }, { $inc: { balance: 100 } }, { session });
  session.commitTransaction();
} catch (e) {
  session.abortTransaction();
} finally {
  session.endSession();
}
```

---

## 八、性能优化

### 1. 查询优化

- 使用 explain 分析执行计划
- 确保查询走索引，避免 COLLSCAN
- 使用投影减少返回字段
- 使用 limit 限制返回数量
- 避免 $where（JavaScript 表达式，慢）

### 2. 索引优化

- 遵守 ESR 规则创建复合索引
- 删除未使用的索引（减少写入开销）
- 使用覆盖索引
- 避免过多索引（每个索引都会增加写入成本）

### 3. 写入优化

- 批量写入（bulkWrite）
- 合理设置 Write Concern（非关键数据用 w:1）
- 使用 upsert 代替先查后写

---

## 九、常见面试题

**Q1：MongoDB 和 MySQL 如何选择？**
| 场景 | 推荐 |
|---|---|
| 数据结构固定、强事务 | MySQL |
| 数据结构灵活、快速迭代 | MongoDB |
| 复杂 JOIN 查询 | MySQL |
| 高并发读写、日志数据 | MongoDB |
| 需要地理空间查询 | MongoDB |

**Q2：ObjectId 的组成？**
- 12 字节：4 字节时间戳 + 5 字节随机值 + 3 字节计数器
- 保证全局唯一，默认按时间排序

**Q3：MongoDB 的锁机制？**
- WiredTiger 引擎使用文档级并发控制
- 写操作锁文档，读操作不阻塞（MVCC）
- 早期版本使用库级/表级锁，现已优化

**Q4：oplog 是什么？**
- 操作日志，记录主节点的所有写操作
- 副本集从节点通过 tailing oplog 同步数据
- 大小固定（默认磁盘 5%），循环写入

**Q5：MongoDB 内存管理？**
- WiredTiger 默认使用 50% 可用内存 - 1GB 作为缓存
- 数据优先在内存中操作，满后刷盘
- 通过 `storage.wiredTiger.engineConfig.cacheSizeGB` 调整