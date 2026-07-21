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
