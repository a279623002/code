# Elasticsearch 面试笔记

> Elasticsearch 是分布式搜索引擎，基于 Lucene 构建，核心围绕倒排索引、分片、集群、查询 DSL、聚合分析。

---

## 一、基础概念

### 1. ES 是什么？核心特点？

- 基于 Lucene 的分布式全文搜索引擎
- 近实时搜索（1 秒延迟）
- RESTful API（JSON over HTTP）
- 分布式、高可用、水平扩展

### 2. 核心术语对照

| ES | 关系型数据库 |
|---|---|
| Index（索引） | Database |
| Type（已废弃 7.x+） | Table |
| Document（文档） | Row |
| Field（字段） | Column |
| Mapping（映射） | Schema |
| Shard（分片） | 分库分表 |

---

## 二、倒排索引

### 1. 什么是倒排索引？

正排索引：文档 ID → 单词列表
倒排索引：单词 → 文档 ID 列表（Posting List）

**示例**：
```
文档1: "Elasticsearch is fast"
文档2: "Elasticsearch is powerful"
文档3: "Search is important"

倒排索引：
Elasticsearch → [1, 2]
is           → [1, 2, 3]
fast         → [1]
powerful     → [2]
Search       → [3]
important    → [3]
```

### 2. 倒排索引数据结构

- **Term Dictionary**：排序的单词列表，B-Tree 存储，二分查找 O(log n)
- **Posting List**：包含该词的文档 ID 列表
- **Term Index**：Term Dictionary 的索引（FST 有限状态转换器），前缀压缩

### 3. FST (Finite State Transducer) 是什么？

- 有向无环图，共享前缀和后缀，极大压缩存储
- 支持前缀查询、范围查询
- 比 HashMap 节省 90% 内存

---

## 三、写入与搜索流程

### 1. 写入流程

```
1. 请求到达协调节点（Coordinating Node）
2. 路由计算：shard = hash(_id) % 主分片数
3. 转发到主分片（Primary Shard）
4. 主分片写入 translog + 内存 buffer
5. 主分片转发到副本分片（Replica Shard）
6. 副本写入成功后返回协调节点
7. 协调节点返回客户端
```

### 2. refresh 与 flush 区别

| | refresh | flush |
|---|---|---|
| 触发 | 默认 1s（可配） | translog 达到阈值 / 默认 30min |
| 操作 | buffer → segment | segment → 磁盘 + 清空 translog |
| 可见性 | 数据可搜索 | 数据持久化 |
| 性能 | 轻量 | 重量（fsync） |

### 3. 搜索流程（两阶段）

**Query Phase（查询阶段）**：
1. 协调节点广播到所有分片
2. 各分片返回 doc_id + score（默认前 10）
3. 协调节点合并排序

**Fetch Phase（取回阶段）**：
1. 协调节点根据确定的 doc_id 列表
2. 向对应分片获取完整文档内容
3. 返回给客户端

### 4. 为什么 ES 搜索是近实时的？

写入数据先进入内存 buffer，默认 1s refresh 一次生成 segment 才能被搜索到。可调小 `refresh_interval` 但会增加开销。

---

## 四、分片与集群

### 1. 主分片和副本分片

| | 主分片（Primary） | 副本分片（Replica） |
|---|---|---|
| 数量 | 索引创建时确定，不可修改 | 可动态修改 |
| 读写 | 处理读写 | 只处理读，从主分片同步 |
| 默认 | 5.x: 5个；6.x: 5个；7.x+: 1个 | 1个 |

### 2. 分片分配策略

- 主分片和副本分片不在同一节点
- 分片数 = 主分片数 × (1 + 副本数)
- 建议：单分片 10-50GB，堆内存 1GB 对应 20 个分片

### 3. 集群健康状态

- **Green**：所有主分片和副本分片都正常
- **Yellow**：所有主分片正常，部分副本分片未分配
- **Red**：部分主分片未分配，数据不完整

---

## 五、查询 DSL

### 1. 全文查询 vs 精确查询

**全文查询**（会分词）：
```json
{ "match": { "title": "elasticsearch guide" } }
{ "match_phrase": { "title": "elasticsearch guide" } }
{ "multi_match": { "query": "es", "fields": ["title", "content"] } }
```

**精确查询**（不分词）：
```json
{ "term": { "status": "active" } }
{ "terms": { "status": ["active", "pending"] } }
{ "range": { "price": { "gte": 100, "lte": 200 } } }
{ "exists": { "field": "email" } }
```

### 2. Bool 查询（组合查询）

```json
{
  "bool": {
    "must":     [{ "match": { "title": "elasticsearch" } }],   // AND + 算分
    "filter":   [{ "term": { "status": "published" } }],       // AND + 不算分
    "should":   [{ "match": { "content": "tutorial" } }],      // OR
    "must_not": [{ "term": { "deleted": true } }]              // NOT
  }
}
```

### 3. filter 和 must 区别

| | must | filter |
|---|---|---|
| 算分 | 会影响相关性评分 | 不影响评分 |
| 缓存 | 不缓存 | 自动缓存结果 |
| 用途 | 需要相关性排序 | 精确过滤/范围过滤 |

---

## 六、聚合分析

### 1. 聚合三大类

| 类型 | 说明 | 示例 |
|---|---|---|
| Bucket | 分桶分组 | terms、range、date_histogram |
| Metric | 指标计算 | avg、sum、max、min、stats |
| Pipeline | 对聚合结果再聚合 | moving_avg、derivative |

### 2. 聚合示例

```json
{
  "aggs": {
    "by_category": {
      "terms": { "field": "category", "size": 10 },
      "aggs": {
        "avg_price": { "avg": { "field": "price" } }
      }
    }
  }
}
```

---

## 七、性能优化

### 1. 写入优化

- 批量写入（bulk API），建议 5-15MB 一批
- 写入前关闭 refresh：`refresh_interval: -1`
- 写入前关闭副本：`number_of_replicas: 0`，写完后恢复
- 使用 SSD 硬盘
- 增大 translog flush 阈值

### 2. 查询优化

- **filter 代替 must**：利用缓存
- **合理设置分片数**：避免过多小分片
- **使用 routing**：减少查询的分片数
- **限制字段返回**：`_source: ["title", "price"]`
- **避免深度分页**：使用 search_after 代替 from+size
- **force merge 只读索引**：减少 segment 数量

### 3. 深度分页问题

`from + size` 超过 10000 需要修改 `max_result_window`，性能差。替代方案：

- **search_after**：基于上一页排序值查询下一页（推荐）
- **scroll**：生成快照，适合全量导出（不推荐实时查询）

---

## 八、Mapping 映射

### 1. 动态映射 vs 显式映射

- **动态映射**：ES 自动推断字段类型
- **显式映射**：手动定义字段类型和分词器

### 2. 常见字段类型

| 类型 | 说明 |
|---|---|
| text | 全文搜索，会被分词 |
| keyword | 精确匹配，不分词，适合聚合/排序 |
| long / integer / short | 整数 |
| float / double | 浮点数 |
| date | 日期 |
| boolean | 布尔 |
| geo_point | 地理位置 |
| nested | 嵌套对象（独立索引） |
| join | 父子关系 |

### 3. text 和 keyword 区别

| | text | keyword |
|---|---|---|
| 分词 | 会被分词器分词 | 不分词，完整存储 |
| 用途 | 全文搜索 | 精确匹配、聚合、排序 |
| 示例 | 文章内容、描述 | 状态、标签、ID |

---

## 九、分词器

### 1. 常见分词器

| 分词器 | 说明 |
|---|---|
| standard | 默认，按词边界分割 |
| ik_smart | IK 中文粗粒度分词 |
| ik_max_word | IK 中文细粒度分词 |
| pinyin | 拼音分词 |
| whitespace | 按空格分割 |
| keyword | 不分词，整体作为一个词 |

### 2. IK 分词器

```
输入："中华人民共和国国歌"

ik_smart：中华人民共和国 | 国歌
ik_max_word：中华人民共和国 | 中华人民 | 中华 | 华人 | 人民共和国 | 人民 | 共和国 | 共和 | 国歌
```

---

## 十、常见面试题

**Q1：ES 如何实现高可用？**
- 主分片 + 副本分片机制，副本分片分布在不同节点
- 主分片故障时，自动从副本分片选举新主分片
- 最少需要 2 个节点（主+副本）

**Q2：ES 的数据一致性模型？**
- 最终一致性，写入主分片后异步同步副本
- 可通过 `wait_for_active_shards` 参数控制写入一致性级别

**Q3：为什么 ES 比 MySQL 搜索快？**
- 倒排索引 vs B+Tree 索引
- 分布式并行搜索，多个分片同时查询
- 相关性评分算法（TF-IDF/BM25）

**Q4：ES 和 Solr 的区别？**
- ES 分布式自带，Solr 基于 Zookeeper
- ES 近实时，Solr 历史更久
- ES RESTful 更友好，生态更好（ELK Stack）