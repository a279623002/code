# MySQL 分库分表完整讲解
## 一、两种拆分类型
### 1、垂直拆分（简单）
- **垂直分库**：按**业务模块**拆成不同数据库
> 例：`db_user`(用户库)、`db_order`(订单库)、`db_goods`(商品库)，库之间表完全不一样，**没有分片路由逻辑**，应用直接连不同库。
- **垂直分表**：一张宽表，字段拆分，放在同一个库不同表
> 例：`user`拆成 `user_base`(基础字段)、`user_ext`(扩展大字段)

> 👉 垂直拆分**不需要特殊CRUD改造**，只是换库/换表名。

### 2、水平拆分（大家常说的「分库分表」，重点）
**同一张逻辑表，数据行分散到多个库、多张物理表**。
> 逻辑表：`t_order`
> 物理实例：4个库 `db_order_0`、`db_order_1`、`db_order_2`、`db_order_3`；每个库里面4张表 `t_order_0 ~ t_order_3`，一共 16 张分片表。

核心概念：
1. **分片键(sharding‑key)**：用来决定这条数据落到哪个库、哪张表的字段，例如 `order_id` / `user_id`
2. **分片算法**：根据分片键计算库下标、表下标
3. **中间件**：负责 SQL 路由，不用自己写大量判断逻辑

## 二、两种实现架构
### 方案A：客户端分片（Sharding‑JDBC，最常用）
Jar包嵌入你的应用程序，**应用内部完成路由计算**，直接访问各个MySQL实例，无额外代理服务。
> Java生态首选，轻量，性能好。

### 方案B：代理层分片（Sharding‑Proxy / MyCat）
独立中间件服务，应用只连接代理；代理接收SQL、解析路由、访问各个MySQL分片，最后合并结果返回。
> 多语言项目适合，独立部署，有网络转发损耗。

> 两者**后续增删改查写法几乎一样**，SQL层面基本无感知。

## 三、常见分片路由算法
假设分片键：`order_id`，总分片数量=16
1. **取模（哈希）分片【最常用】**
```
dbIndex = order_id % 4    -- 算出落到哪个库
tableIndex = order_id / 4 % 4 --算出库内哪张表
```
优点：数据均匀；缺点：扩容麻烦，需要数据迁移。
2. **范围分片（时间分片）**
按月分表：`t_order_202601`、`t_order_202602`
优点：扩容简单；缺点：热点（新数据全部落在最新分片）
3. **一致性哈希**：减少扩容时数据迁移量

> ⚠️ 分片键**一旦选定，绝大多数CRUD最好带上它**。

## 四、分库分表后：增删改查如何操作
> 示例环境：逻辑表`t_order`；分片键`order_id`；分片总数16

### ✅ 1、INSERT 新增
**必须携带分片键！**
1. 应用传入分片键 `order_id`
2. 中间件执行分片算法，算出目标库+目标物理表
3. 路由，数据插入对应分片表

```sql
-- 业务层SQL（你写的SQL不变）
INSERT INTO t_order(order_id,user_id,amount) VALUES(10001, 888, 99.9);
```
> ❌ **禁止不带分片键插入**：中间件不知道插哪里，直接报错。
> ⚠️ 主键不能依赖数据库自增！多个分片自增ID会重复 → 使用雪花算法、UUID生成全局唯一ID。

### ✅ 2、UPDATE 更新
#### 最优方案：带上分片键（定点路由，只访问1个分片）
```sql
UPDATE t_order SET amount=100 WHERE order_id=10001;
```
中间件算出分片位置，只执行单库单表更新，性能高。

#### ❌ 不带分片键更新（广播路由，危险）
```sql
UPDATE t_order SET amount=100 WHERE user_id=888;
```
中间件**向全部16个分片都执行这条SQL**，全分片扫描，性能差、锁冲突风险高。
> 解决方案：如果经常按`user_id`查询更新，就建立**复合分片键**或者搭建**二级索引表/ES搜索引擎**。

### ✅ 3、DELETE 删除
规则和update完全一致
```sql
--推荐（带分片键，定点删除）
DELETE FROM t_order WHERE order_id=10001;
```

### ✅ 4、SELECT 查询，分3种场景
#### 场景1：查询条件带上分片键【最优，强烈推荐】
```sql
SELECT * FROM t_order WHERE order_id=10001;
```
👉 只路由单个分片，查询速度和单库几乎一样。

#### 场景2：不带分片键（无条件 / 条件不含分片键）→广播查询
```sql
SELECT * FROM t_order WHERE user_id=888;
```
👉 中间件发送SQL到**所有分片**，拿到所有分片结果，在内存合并、返回给应用。
> 分片越多性能越差，大系统禁止高频不带分片键查询。

#### 场景3：分页、排序、聚合查询（大坑！）
原生 `limit offset` 在分表直接查询会出错！
> 例：`select * from t_order order by create_time limit 10,20`
每个分片各自返回前30条，再由中间件内存排序，取出最终结果。
- count 总数：需要所有分片count，再累加
- sum、max/min：分片计算，中间件再二次聚合

## 五、跨分片Join问题
分库之后**不同分片库不能做MySQL原生join**。
解决办法：
1. 大表小表：小表做广播表（每个分片库都存一份）
2. 应用层Join：分别查询两张表，代码内存关联
3. 用ES做检索，放弃数据库join

## 六、Sharding‑JDBC最简配置示例（取模分片）
```yaml
spring:
  shardingsphere:
    datasource:
      names: ds0,ds1,ds2,ds3
      #省略四个数据源mysql连接配置
    rules:
      sharding:
        tables:
          t_order:
            actual-data-nodes: ds$->{0..3}.t_order_$->{0..3}
            database-strategy:
              standard:
                sharding-column: order_id
                sharding-algorithm-name: order-mod-db
            table-strategy:
              standard:
                sharding-column: order_id
                sharding-algorithm-name: order-mod-table
```
业务代码完全不需要改路由逻辑，正常写CRUD即可。

## 七、分库分表后常见痛点 & 解决方案
| 问题 | 方案 |
|---|---|
| 分布式事务跨分片更新失败 | Seata AT/TCC事务；尽量业务设计避免跨分片事务 |
| 非分片键查询慢 | 建立二级索引表、同步数据到Elasticsearch |
| ID主键重复 | 雪花算法、Redis生成全局唯一ID |
| 扩容分片需要迁移大量数据 | 提前预留分片数；使用一致性hash；双写迁移方案 |
| join复杂 | 广播表、应用层join、ES检索 |

## 八、什么时候不要上来就分库分表
优化优先级：**索引优化 → SQL优化 → 读写分离 → 单库分表 → 最后才分库分表**。
分库分表带来很高的复杂度，能不分尽量不分。

如果你需要，我可以给你一份可直接运行的 SpringBoot + Sharding‑JDBC 完整Demo工程代码。