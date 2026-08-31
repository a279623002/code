# 微服务 CAP 理论 + 选型策略 + 开发示例

## 一、先回顾 CAP 三要素

> 
> 在**分布式环境（多节点）**下，三者**不可能同时满足**，只能三选二

- **C 一致性 Consistency**：所有节点同一时刻看到一样的数据。更新后任意节点读取都是最新值。
- **A 可用性 Availability**：每个请求都**能收到响应**，不会卡死/超时失败（但返回的数据可能旧）。
- **P 分区容错 Partition tolerance**：网络故障、消息丢包、节点断开，**系统依然可以运行**。

> 
> ⚠️ 重点：**分布式系统 P 是必须保证的，你不能放弃分区容错。**
> 所以真实选型只有两个选择：
> 
> 
> 1. **AP：放弃强一致性，保可用 + 分区容错（绝大多数微服务业务）**
> 2. **CP：放弃可用性，保强一致 + 分区容错（强数据业务）**
> 不存在 CA，放弃 P 就不是分布式。

## 二、业务场景怎么选 AP / CP

### 1）AP（优先选，互联网微服务默认方案）

> 
> 最终一致性，高可用优先，短暂数据不一致可以接受
> 适合绝大多数业务：

- 用户信息、商品列表、订单查询、首页数据、推荐、评论
- 支付状态查询（短暂延迟没关系，最后同步成正确即可）
**最终一致性实现方案**：
- MQ异步通知、本地消息表、Seata AT/TCC、Redis缓存延迟更新、Canal binlog同步> 
> 牺牲强一致，换来极高可用性，网络抖动服务不会挂。

### 2）CP（强一致，可用性降低）

> 
> 网络分区发生时，服务阻塞等待，直到数据达成一致，期间不可用
> 适合强资金、强库存、锁场景：

- 库存扣减、转账交易、分布式锁、订单状态强流转、资金账户
**CP组件例子**：Zookeeper、etcd、Redis Redlock、Seata‑TCC

### 快速选型对照表

| 选型 | 特征 | 典型组件 | 业务案例 |
| --- | --- | --- | --- |
| AP | 高可用，最终一致 | Nacos,Eureka,RocketMQ/Kafka,Seata‑AT | 普通订单、用户、商品 |
| CP | 强一致，网络故障可能不可用 | Zookeeper,Etcd,Seata‑TCC | 库存、资金、分布式锁 |

> 
> ✅ **国内微服务主流选型：注册配置中心 Nacos(AP) + 分布式事务 AT(AP最终一致)**

# 三、开发实战示例：SpringCloud Alibaba AP方案（最常用）

架构：订单服务 `order‑service` 远程调用 库存服务 `stock‑service`，保证**最终一致性（AP）**
技术栈：SpringBoot + SpringCloudAlibaba + Nacos + Seata‑AT + Mysql

## 1. AP思路说明（最终一致性）

1. 订单服务创建订单
2. 远程调用库存扣减
3. 如果库存扣减失败，通过Seata回滚；
4. 允许短暂不一致，最终数据统一，优先保证服务可用。

> 
> Seata‑AT 就是典型 **AP（最终一致性）**，不是强一致。

### 1）父pom依赖

```
<dependencyManagement>
    <dependencies>
        <dependency>
            <groupId>com.alibaba.cloud</groupId>
            <artifactId>spring-cloud-alibaba-dependencies</artifactId>
            <version>2022.0.0.0</version>
            <type>pom</type>
            <scope>import</scope>
        </dependency>
    </dependencies>
</dependencyManagement>
```

### 2）订单服务 application.yml

```
server:
  port: 8081
spring:
  application:
    name: order-service
  datasource:
    driver-class-name: com.mysql.cj.jdbc.Driver
    url: jdbc:mysql://127.0.0.1:3306/db_order?useUnicode=true
    username: root
    password: 123456
  cloud:
    nacos:
      discovery:
        server-addr: 127.0.0.1:8848
    seata:
      tx-service-group: my_tx_group
```

### 3）库存服务 application.yml

```
server:
  port: 8082
spring:
  application:
    name: stock-service
  datasource:
    url: jdbc:mysql://127.0.0.1:3306/db_stock?useUnicode=true
    username: root
    password: 123456
  cloud:
    nacos:
      discovery:
        server-addr: 127.0.0.1:8848
    seata:
      tx-service-group: my_tx_group
```

> 
> seata 需要每个库创建回滚日志表 `undo_log`

```
CREATE TABLE `undo_log` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `branch_id` bigint(20) NOT NULL,
  `xid` varchar(100) NOT NULL,
  `context` varchar(128) NOT NULL,
  `rollback_info` longblob NOT NULL,
  `log_status` int(11) NOT NULL,
  `log_created` datetime NOT NULL,
  `log_modified` datetime NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `ux_undo_log` (`xid`,`branch_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
```

### 4）订单服务代码，开启分布式事务 @GlobalTransactional（AP最终一致）

```
@RestController
@RequestMapping("/order")
public class OrderController {

    @Autowired
    private OrderMapper orderMapper;

    @Autowired
    private StockFeignClient stockFeignClient;

    //全局事务注解，Seata AT模式 AP最终一致性
    @GlobalTransactional
    @PostMapping("/create")
    public String createOrder(Long productId,Integer count){
        //1.创建订单
        Order order = new Order();
        order.setProductId(productId);
        order.setCount(count);
        orderMapper.insert(order);

        //2.远程扣减库存
        stockFeignClient.deduct(productId,count);

        return "下单成功";
    }
}
```

### 5）Feign远程调用接口

```
@FeignClient("stock-service")
public interface StockFeignClient {
    @PostMapping("/stock/deduct")
    String deduct(@RequestParam Long productId, @RequestParam Integer count);
}
```

### 6）库存服务

```
@RestController
@RequestMapping("/stock")
public class StockController {

    @Autowired
    private StockMapper stockMapper;

    @PostMapping("/deduct")
    public String deduct(Long productId,Integer count){
        stockMapper.deductStock(productId,count);
        //模拟异常，触发分布式回滚
        // int i= 1/0;
        return "库存扣减成功";
    }
}
```

> 
> 👉 AT模式工作原理（AP）：
> 一阶段：执行业务SQL，记录undo_log，**立刻提交本地事务，释放锁**（高可用）
> 二阶段：成功什么都不做；失败异步补偿回滚。
> 执行期间不阻塞，优先可用性，属于**最终一致性 AP**。

# 四、CP 方案示例（强一致，TCC）

如果库存场景你要用CP强一致，选 Seata‑TCC，网络分区时会锁住资源直到事务完成，可用性下降。
TCC三段式：Try - Confirm - Cancel

```
@LocalTCC
public interface StockTccService {
    //Try：预留资源，冻结库存
    @TwoPhaseBusinessAction(name = "stockTcc",commitMethod = "confirm",rollbackMethod = "cancel")
    boolean tryDeduct(@BusinessActionContextParameter(paramName = "pid") Long pid,
                       @BusinessActionContextParameter(paramName = "cnt") Integer cnt);
    //Confirm：正式扣减
    boolean confirm(BusinessActionContext ctx);
    //Cancel：释放冻结库存，回滚
    boolean cancel(BusinessActionContext ctx);
}
```

> 
> TCC属于手动补偿方案，偏向CP，适合资金等高可靠场景，开发成本高。

# 五、注册中心 CAP对比

1. **Nacos 默认 AP**：集群网络分区，依然可以注册发现，保证可用，短时间可能读到旧实例列表。
Nacos 也可以切换成 CP模式（使用Raft）
2. **Eureka AP**：经典AP，网络故障节点不会剔除
3. **Zookeeper CP**：网络分区时，没有过半节点可用，整个集群不可用

# 六、总结开发选型口诀

1. **90%普通微服务业务：选 AP（最终一致性），Nacos+Seata‑AT**
2. **资金、库存、锁等高风险业务：选 CP，TCC / Redis分布式锁 / etcd**
3. **永远不要幻想CAP三选CA，分布式系统P必选**

如果你需要，我可以给你一份完整可运行的 AP微服务项目demo，包含nacos、seata服务端启动脚本。