# Redis 五大基础数据类型 + 三种高级类型 详解（含结构、命令、真实业务场景示例）

## 一、基础5大核心类型（日常开发90%场景使用）

### 1. String 字符串（最基础）
#### 底层结构
二进制安全字符串，可存字符串、数字、二进制图片/序列化对象，单值单key。

#### 核心命令
```redis
set k1 v1 ex 3600    # 设置+过期1小时
get k1
incr num             # 自增（原子）
mset k2 v2 k3 v3     # 批量设置
mget k2 k3
append k1 "_suffix"
```

#### 业务场景示例
- **接口限流计数器**
```redis
# 用户1001每分钟最多访问100次
set limit:user:1001 0 ex 60
incr limit:user:1001
get limit:user:1001
```

- **分布式全局 ID**
`incr order_id` 每次自增生成唯一订单号

- **验证码缓存**
`set sms:13800138000 123456 ex 300` 5 分钟过期

- **简单热点缓存**
商品基础信息、用户基础信息序列化 JSON 存入 String

- **分布式锁（简易版）**
`set lock:goods:100 1 nx ex 10` nx 不存在才设置，防死锁


### 2. List 列表（双向链表）
#### 底层结构
双向链表，左右两端操作 O(1)，中间遍历 O(n)；允许重复元素、有序。

#### 核心命令
```redis
lpush list1 a b c   # 头插
rpush list1 d e     # 尾插
lpop list1          # 左弹出
rpop list1          # 右弹出
lrange list1 0 -1   # 全量查询
ltrim list1 0 99    # 保留前100条，截断
brpop list1 3       # 阻塞弹出3秒，无数据等待
```

#### 业务场景示例
- **消息队列（简易队列）**
生产者：`rpush mq:order "订单JSON"`
消费者：`brpop mq:order 0` 阻塞消费

- **最新消息 / 动态列表**
用户首页 10 条最新浏览记录：`lpush view:user:1001 商品ID` + `ltrim` 限制长度

- **栈、队列结构模拟**
栈：`lpush` + `lpop`；队列：`rpush` + `lpop`

- **日志临时存储**
系统实时日志写入 list，定时脚本批量读取落库


### 3. Set 集合（无序、去重）
#### 底层结构
哈希表实现，无序、自动去重，支持交 / 并 / 差集运算。

#### 核心命令
```redis
sadd set1 a b c a
smembers set1        # 查看所有元素
sismember set1 a     # 判断是否存在
srem set1 a
scard set1           # 集合元素总数
sinter set1 set2     # 交集（共同好友）
sunion set1 set2     # 并集
sdiff set1 set2      # 差集
spop set1 1          # 随机取出元素
```

#### 业务场景示例
- **用户点赞 / 收藏去重**
`sadd like:article:100 1001` 用户 1001 点赞文章 100，自动防重复点赞

- **社交共同好友**
用户 A 好友集合 `friend:1001`，用户 B `friend:1002`
`sinter friend:1001 friend:1002` 算出共同好友

- **抽奖随机抽取**
所有参与用户存入 set，`spop` 随机弹出中奖用户，不重复

- **IP 黑名单**
`sadd black_ip 192.168.1.1`，访问时`sismember`校验


### 4. Hash 哈希（对象存储，类 Map）
#### 底层结构
key → field-value 映射，适合存储结构化对象，无需序列化拆分字段。

#### 核心命令
```redis
hset user:1001 name "张三" age 20 phone "138xxxx"
hget user:1001 name
hgetall user:1001
hmset user:1002 name 李四 age 22
hincrby user:1001 age 1  # 字段自增
hdel user:1001 phone
hexists user:1001 name
```

#### 业务场景示例
- **用户信息缓存（最优方案）**
不用把整个对象序列化 String，只更新修改字段：修改年龄仅`hincrby`

- **商品库存 + 价格复合缓存**
`hset goods:200 stock 99 price 99.9 sales 1000`

- **购物车**
key=`cart:user1001`，field = 商品 ID，value = 购买数量
`hincrby cart:1001 10001 1` 加购一件商品


### 5. Sorted Set (ZSet) 有序集合（带权重排序）
#### 底层结构
跳表 + 哈希，元素唯一，每个元素绑定 score 分数，按 score 自动排序，支持范围分页。

#### 核心命令
```redis
zadd rank 90 user1 85 user2 95 user3
zrange rank 0 -1 withscores  # 从小到大
zrevrange rank 0 9 withscores # 从大到小（TOP10）
zincrby rank 5 user2         # 增加分数
zcard rank
zrangebyscore rank 80 100    # 分数区间查询
zrem rank user1
```

#### 业务场景示例
- **排行榜（核心场景）**
直播间送礼榜、商品销量榜、积分排行榜
`zrevrange sales_rank 0 9` 查出销量前 10 商品

- **延迟队列（简易版）**
score 存时间戳，轮询 `zrangebyscore queue 0 当前时间` 取出到期任务

- **带权重的消息推送**
高优先级消息 score 更大，优先读取

- **分页有序数据**
文章按阅读量排序分页展示


## 二、3 种高级特殊类型（拓展场景）

### 1. BitMap 位图（基于 String 实现）
本质是 String，按 bit 位存储，极致节省内存。

#### 核心命令
```redis
setbit sign:20260721 1001 1 # 用户1001今日签到
getbit sign:20260721 1001
bitcount sign:20260721      # 今日签到总人数
bitop and res sign1 sign2   # 连续两天都签到用户
```

#### 场景示例
- 每日用户签到统计（百万用户仅几 KB 内存）
- 在线用户状态标记、布隆过滤器底层辅助


### 2. HyperLogLog 基数统计（去重总数，模糊估算）
海量数据下低成本统计独立元素数量，误差 0.81%，不存储原始数据。

#### 核心命令
```redis
pfadd uv:page1 user1 user2 user1
pfcount uv:page1  # 独立访客数
pfmerge uv:total uv:page1 uv:page2
```

#### 场景示例
- 页面 UV 统计（不去重，只算多少人访问）
- 直播间独立访客、活动访客总量统计


### 3. Geo 地理位置（基于 ZSet 封装）
存储经纬度，计算两点距离、范围内周边点位。

#### 核心命令
```redis
geoadd city 116.40 39.90 beijing
geodist city beijing shanghai km  # 两地距离km
georadius city 116.40 39.90 100 km # 方圆100km城市
```

#### 场景示例
- 附近门店、附近人、网约车司机位置检索
- 同城商品、周边服务筛选


## 三、各类型选型速查表（快速判断用什么）

| 需求场景 | 推荐类型 | 理由 |
| --- | --- | --- |
| 简单 KV、计数器、验证码、分布式锁 | String | 最轻量，原子自增 |
| 结构化对象（用户 / 商品多字段） | Hash | 局部更新，无需序列化 |
| 消息队列、浏览历史、日志列表 | List | 两端快速操作，阻塞消费 |
| 去重、交集差集（点赞、好友、黑名单） | Set | 自动去重，集合运算 |
| 排行榜、延时任务、有序分页 | ZSet | 按分数自动排序，范围查询 |
| 签到、状态标记、海量布尔值 | BitMap | 极致节省内存 |
| 海量独立访客 UV 统计 | HyperLogLog | 极低内存，只算基数 |
| 经纬度、附近门店 / 人 | Geo | 内置距离、范围查询 |


## 四、补充区分易混类型
- List：有序可重复；Set：无序不可重复；ZSet：有序不可重复（靠分数）
- String 适合整体更新；Hash 适合局部字段更新
- HyperLogLog 只能统计数量，拿不到具体用户；Set 可以拿到完整元素
- BitMap 只适合 0/1 二值标记，大量 ID 存储比 Set 省千倍内存

---

## 五、缓存更新策略

| 策略 | 流程 | 优点 | 缺点 |
|---|---|---|---|
| **Cache Aside（旁路缓存）** | 读：先读缓存，没有则读 DB 并回写；写：先更新 DB，再删缓存 | 最常用，逻辑简单 | 短暂不一致 |
| **Read/Write Through** | 读写都走缓存层，缓存负责和 DB 同步 | 对业务透明 | 缓存组件复杂 |
| **Write Behind（异步回写）** | 先写缓存，异步批量写 DB | 写性能极高 | 丢数据风险大 |

**推荐**：日常业务用 **Cache Aside + 设置过期时间 + 异步消息补偿**。

```
Cache Aside 读流程：
    用户 ──→ 缓存
              ↓ 命中
            返回
              ↓ 未命中
            读 DB
            写缓存
            返回

Cache Aside 写流程：
    用户 ──→ 更新 DB
              ↓
            删除缓存（不是更新缓存）
```

> 为什么写操作是**删缓存**而不是更新缓存？
> 因为并发写时"更新缓存"容易把旧值覆盖成新值；删除缓存让下次读时从 DB 重新加载最新值。

---

## 六、缓存淘汰（回收）策略

Redis 6 种淘汰策略：

| 策略 | 含义 |
|---|---|
| `noeviction` | 默认，内存满了直接拒绝写入 |
| `allkeys-lru` | 所有 key 中，淘汰最近最少使用 |
| `allkeys-lfu` | 所有 key 中，淘汰使用频率最少 |
| `allkeys-random` | 所有 key 中随机淘汰 |
| `volatile-lru` | 带过期时间的 key 中，淘汰最近最少使用 |
| `volatile-lfu` | 带过期时间的 key 中，淘汰使用频率最少 |
| `volatile-ttl` | 带过期时间的 key 中，淘汰即将过期的 |
| `volatile-random` | 带过期时间的 key 中随机淘汰 |

**生产推荐**：
- 缓存场景：`allkeys-lru`
- 需要区分冷热：`allkeys-lfu`

配置：
```redis
maxmemory 2gb
maxmemory-policy allkeys-lru
```

---

## 七、Pipeline

### 1. 概念图

```
普通模式（往返 3 次）：
    客户端 ──→ set a 1 ──→ 服务端
    服务端 ──→ OK ──→ 客户端
    客户端 ──→ set b 2 ──→ 服务端
    服务端 ──→ OK ──→ 客户端
    客户端 ──→ set c 3 ──→ 服务端
    服务端 ──→ OK ──→ 客户端

Pipeline 模式（打包 1 次发送）：
    客户端 ──→ [set a 1, set b 2, set c 3] ──→ 服务端
    服务端 ──→ [OK, OK, OK] ──→ 客户端
```

### 2. 原理

Pipeline 不是原子操作，只是把多个命令打包成一次网络请求，减少 RTT（往返时间）。

### 3. 应用

```go
// go-redis Pipeline 示例
pipe := rdb.Pipeline()
incr := pipe.Incr(ctx, "counter")
pipe.Expire(ctx, "counter", time.Hour)
_, err := pipe.Exec(ctx)
fmt.Println(incr.Val())
```

**适用**：批量写、批量读，能显著提升吞吐量。
**不适用**：需要事务原子性的场景（用 `MULTI/EXEC` 或 Lua）。

---

## 八、缓存三大问题

### 1. 缓存穿透

**现象**：查询一个**一定不存在**的数据，缓存没有，DB 也没有，每次请求都打到 DB。

**解决**：
- **布隆过滤器**：在缓存前加一层，快速判断 key 是否可能存在
- **缓存空值**：把不存在的 key 也缓存起来（设置较短过期时间）
- **参数校验**：非法 id 直接拦截

```
请求 ──→ 布隆过滤器？
              ↓ 不存在
            直接返回空
              ↓ 可能存在
            查缓存 ──→ 查 DB
```

### 2. 缓存击穿

**现象**：某个**热点 key 突然过期**，大量请求同时打到 DB。

**解决**：
- **互斥锁**：只允许一个线程去加载，其他等待
- **逻辑过期**：不真正设置 TTL，业务代码判断是否需要重建
- **热点 key 永不过期**：后台异步更新

```go
// 互斥锁示例（简化版）
func getHotKey(key string) string {
    val, ok := rdb.Get(ctx, key).Result()
    if ok == nil {
        return val
    }

    // 只有一个人去拿锁
    locked, _ := rdb.SetNX(ctx, "lock:"+key, "1", 10*time.Second).Result()
    if locked {
        defer rdb.Del(ctx, "lock:"+key)
        val = queryDB(key)
        rdb.Set(ctx, key, val, time.Minute)
    } else {
        time.Sleep(50 * time.Millisecond)
        return getHotKey(key) // 重试
    }
    return val
}
```

### 3. 缓存雪崩

**现象**：大量 key **同时过期**，或 Redis 宕机，导致 DB 压力骤增。

**解决**：
- **过期时间加随机值**：避免同时失效
- **多级缓存**：本地缓存（Caffeine）+ Redis + DB
- **限流降级**：熔断、兜底数据
- **Redis 高可用**：主从 + 哨兵 / Cluster

```go
// 过期时间随机
expire := time.Duration(300+rand.Intn(60)) * time.Second
rdb.Set(ctx, key, val, expire)
```

| 问题 | 原因 | 解决 |
|---|---|---|
| 缓存穿透 | 查询不存在数据 | 布隆过滤器、缓存空值 |
| 缓存击穿 | 热点 key 过期 | 互斥锁、逻辑过期 |
| 缓存雪崩 | 大量 key 同时过期 / Redis 宕机 | 随机 TTL、多级缓存、高可用 |

---

## 九、主从复制与主节点选举

### 1. 主从复制

```
        ┌──────────┐
        │  Master  │ ← 写请求
        │  主节点   │
        └────┬─────┘
             │ 复制数据
    ┌────────┼────────┐
    ↓        ↓        ↓
┌───────┐ ┌───────┐ ┌───────┐
│Slave 1│ │Slave 2│ │Slave 3│ ← 读请求
└───────┘ └───────┘ └───────┘
```

**复制方式**：
- 全量复制：初次同步，RDB 文件
- 增量复制：后续通过复制积压缓冲区同步写命令

### 2. 哨兵模式（Sentinel）

```
┌─────────┐     ┌─────────┐     ┌─────────┐
│Sentinel1│────→│Sentinel2│────→│Sentinel3│
└────┬────┘     └────┬────┘     └────┬────┘
     │               │               │
     └───────────────┼───────────────┘
                     ↓
              ┌──────────┐
              │  Master  │
              └────┬─────┘
                   │
        ┌──────────┼──────────┐
        ↓          ↓          ↓
    ┌───────┐  ┌───────┐  ┌───────┐
    │Slave 1│  │Slave 2│  │Slave 3│
    └───────┘  └───────┘  └───────┘
```

**主节点选举流程**：
1. Sentinel 发现 Master 主观下线（SDOWN）
2. 多个 Sentinel 互相确认，达到客观下线（ODOWN）
3. 选举一个 Sentinel 作为 Leader
4. Leader 从 Slave 中选一个作为新 Master（优先级、复制偏移量、runid）
5. 其他 Slave 重新指向新 Master
6. 通知客户端新的主节点地址

### 3. Redis Cluster

```
┌──────────┐     ┌──────────┐     ┌──────────┐
│ Master A │────→│ Master B │────→│ Master C │
│ Slot 0-5460 │  │ Slot 5461-10922│ │ Slot 10923-16383│
└────┬─────┘     └────┬─────┘     └────┬─────┘
     │                │                │
┌────┴─────┐     ┌────┴─────┐     ┌────┴─────┐
│ Slave A1 │     │ Slave B1 │     │ Slave C1 │
└──────────┘     └──────────┘     └──────────┘
```

- 数据分片：16384 个 slot，每个 key 通过 CRC16(key) % 16384 决定 slot
- 无中心架构，节点之间 gossip 协议通信
- 主从自动故障转移

---

## 十、分布式缓存

### 1. 特点
- 多台 Redis 节点共同承担数据
- 支持水平扩展
- 需要解决数据分片、复制、故障转移

### 2. 集群模式对比

| 模式 | 架构 | 特点 |
|---|---|---|
| **主从复制** | 一主多从 | 读写分离，手动故障转移 |
| **哨兵模式** | 主从 + Sentinel | 自动故障转移，可读可写 |
| **Cluster 模式** | 多主多从 + slot 分片 | 自动分片、自动 failover、推荐 |
| **Codis/Twemproxy** | 代理中间件 | 早期方案，现在用得少 |

---

## 十一、Lua 脚本

### 1. “Redis 单线程，命令本来串行，为什么还要 Lua？”

串行 ≠ 业务逻辑原子
- 串行：所有命令排队依次执行；
- 原子逻辑：一组关联操作必须连续执行，中间不能插入别人的命令。
多条独立命令只是排队，中间可以被插队；Lua 脚本是一个整体任务，执行期间队列里其他命令全部等待。

### 2. 示例：原子扣库存

```lua
-- stock.lua
local key = KEYS[1]
local num = tonumber(ARGV[1])
local stock = tonumber(redis.call('get', key) or '0')

if stock >= num then
    redis.call('decrby', key, num)
    return 1  -- 扣减成功
else
    return 0  -- 库存不足
end
```

```go
// go-redis 执行 Lua
script := redis.NewScript(luaScript)
result, err := script.Run(ctx, rdb, []string{"stock:1001"}, 1).Result()
```

### 3. 应用场景
- 分布式锁释放（判断是不是自己加的锁再删）
- 原子扣库存
- 限流滑动窗口

---

## 十二、面试高频问题

### Q1：Redis 为什么快？

**答**：
1. 纯内存操作
2. 单线程避免上下文切换和锁竞争（Redis 6 网络 IO 多线程，命令执行仍单线程）
3. 高效数据结构（SDS、跳表、压缩列表、整数集合）
4. IO 多路复用（epoll）

### Q2：Redis 是单线程还是多线程？

**答**：
- Redis 6 之前：完全单线程
- Redis 6/7：网络 IO 多线程，**命令执行仍是单线程**
- 所以 Redis 命令天然原子，不需要锁

### Q3：持久化 RDB 和 AOF 的区别？

| 特性 | RDB | AOF |
|---|---|---|
| 文件 | 二进制快照 | 命令日志 |
| 恢复速度 | 快 | 慢 |
| 数据安全 | 可能丢最近一次快照 | 丢得少（appendfsync always/everysec） |
| 文件大小 | 小 | 大 |

**生产**：通常 RDB + AOF 混合持久化。

### Q4：分布式锁怎么实现？

```redis
set lock:order:1001 request_id nx ex 30
```

要求：
- `nx`：不存在才设置，保证互斥
- `ex`：设置过期时间，防止死锁
- 释放时用 Lua 判断 request_id 是否一致再删除

### Q5：Redis 和 Memcached 区别？

| 特性 | Redis | Memcached |
|---|---|---|
| 数据结构 | 丰富（5+3 种） | 只有 String |
| 持久化 | 支持 RDB/AOF | 不支持 |
| 集群 | 原生 Cluster | 需客户端实现 |
| 多线程 | 6 之后网络 IO 多线程 | 多线程 |

---

## 十三、一句话总结

- **更新策略**：Cache Aside 最常用，写 DB 删缓存
- **淘汰策略**：`allkeys-lru` 是生产标配
- **Pipeline**：打包命令减少网络 RTT，非原子
- **三大问题**：穿透用布隆过滤器、击穿用互斥锁、雪崩用随机 TTL + 多级缓存
- **高可用**：主从复制 → 哨兵自动选举 → Cluster 自动分片
- **Lua**：在 Redis 里原子执行多命令，适合扣库存、释放锁
