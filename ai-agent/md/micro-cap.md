# 一、CAP回顾 + 微服务选型结论
> CAP：Consistency一致性 / Availability可用性 / Partition‑tolerance分区容错
- **P(分区容错)**：分布式网络一定会出现网络抖动、丢包、断连；**微服务必须保证P**，所以只能二选一：`AP` or `CP`
1. **AP（绝大多数业务选择）：可用性优先，最终一致性**
    电商、订单、用户中心、后台管理、网关。
    实现方案：注册中心(Nacos/Eureka) + 本地缓存 + MQ异步补偿 + 最终一致
2. **CP（强一致场景）：一致性优先，牺牲部分可用性**
    分布式锁、资金账户、库存扣减、配置中心。
    实现方案：etcd / zookeeper / redis‑redlock

> 生产最常见选型：**业务微服务选 AP；强一致组件选 CP**

# 二、Go‑Zero + Nacos AP示例（最终一致性，最常用）
## 场景说明
两个微服务：
- user‑service 用户服务
- order‑service 订单服务
下单后异步更新用户积分，**不强实时一致，走MQ最终一致(AP方案)**，网络分区时两个服务都还能对外可用。

## 1、go-zero 依赖 go.mod
```go
module cap-demo

go 1.22

require (
	github.com/zeromicro/go-zero v1.7.0
	github.com/zeromicro/go‑zrpc/nacos v1.7.0
	github.com/streadway/amqp v1.1.0
)
```

## 2、order‑service 订单微服务 (AP，注册到Nacos)
### order.yaml
```yaml
Name: order.rpc
ListenOn: 0.0.0.0:8081
Etcd:
  Hosts:
  - 127.0.0.1:8848
  Key: order.rpc
```

### order.go
```go
package main

import (
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go‑zrpc"
	"github.com/zeromicro/go‑zrpc/resolver/nacos"
	"cap‑demo/rpc/orderpb"
)

type OrderServer struct{}

func (s *OrderServer) CreateOrder(ctx context.Context, req *orderpb.CreateReq) (*orderpb.CreateResp, error) {
	// 1.本地落库订单
	orderId := fmt.Sprintf("ORD‑%d", req.UserId)
	fmt.Printf("创建订单 %s,用户:%d\n",orderId,req.UserId)

	// AP核心:不同步调用用户服务更新积分，发消息异步做最终一致
	err := sendMqUserPointEvent(req.UserId, 10)
	if err != nil{
		// MQ发送失败只打日志，订单仍然成功！可用性优先(AP)
		fmt.Println("mq发送失败，后续补偿任务重试")
	}

	return &orderpb.CreateResp{OrderId:orderId,Success:true},nil
}

func sendMqUserPointEvent(userId int64, addPoint int32) error{
	// rabbitmq发送消息代码省略
	return nil
}

func main() {
	nacos.Register()
	srv := zrpc.MustNewServer(zrpc.RpcServerConf{
		ServiceConf: service.ServiceConf{
			Name: "order.rpc",
		},
		ListenOn: "0.0.0.0:8081",
		Etcd: zrpc.EtcdConf{
			Hosts: []string{"127.0.0.1:8848"},
			Key: "order.rpc",
		},
	})
	orderpb.RegisterOrderServer(srv, &OrderServer{})
	srv.Start()
}
```

## 3、user‑service 用户积分消费服务
```go
package main

// 消费MQ消息，异步更新积分，最终一致性
func ConsumePointMsg(){
	// 监听rabbitmq，收到消息更新用户积分
	// 失败可以重试，保证最终一致
}
```
> AP特性：订单成功后，短时间内查用户积分看不到新增，**短暂不一致，最终一致；网络分区时订单服务不会阻塞等待用户服务**。

# 三、CP示例：Go + Etcd 分布式锁(强一致性)
> CP场景：库存扣减，必须强一致；网络分区时宁可失败不可超卖。
```go
package main

import (
	"context"
	"fmt"
	"go.etcd.io/etcd/client/v3"
	"time"
)

func main(){
	cli,err := clientv3.New(clientv3.Config{
		Endpoints:[]string{"127.0.0.1:2379"},
		DialTimeout:5*time.Second,
	})
	if err != nil{
		panic(err)
	}
	defer cli.Close()

	lease,err := cli.Grant(context.Background(), 10)
	if err != nil{
		panic(err)
	}
	key := "/stock/lock/1001"
	// TryLock
	txn := cli.Txn(context.Background()).
		If(clientv3.Compare(clientv3.CreateRevision(key),"=",0)).
		Then(clientv3.OpPut(key,"locked",clientv3.WithLease(lease.ID))).
		Else(clientv3.OpGet(key))

	resp,err := txn.Commit()
	if err != nil{
		panic(err)
	}
	if !resp.Succeeded{
		// 获取锁失败，CP策略：直接返回失败，不允许继续扣库存(牺牲可用性换一致性)
		fmt.Println("获取锁失败，库存操作拒绝")
		return
	}

	fmt.Println("拿到分布式锁，扣减库存")
	// 库存扣减业务...

	cli.Revoke(context.Background(), lease.ID)
}
```
> CP行为：etcd集群出现分区，锁获取失败，业务直接报错，**不会继续执行业务，保证数据一致性，放弃可用性**

# 四、AP与CP代码选型对比总结
|方案|选型|业务特征|Go实现方式|
|---|---|---|---|
|AP(首选)|可用性优先，最终一致|订单、用户、商品|go‑zero/rpc + Nacos + RabbitMQ/Kafka异步|
|CP|一致性优先|库存、账户资金、分布式锁|etcd / redis redlock|

# 五、微服务开发实操建议（Go栈）
1. 普通业务微服务：**AP + Nacos**，异步消息解耦实现最终一致
2. 强一致资源：单独走CP组件(etcd/redis锁)，**不要强行让整个微服务集群做CP**
3. 补偿机制：MQ死信队列 + 定时任务扫描，修复不一致数据

如果你需要，我可以给你一份**可直接运行的完整Go‑Zero AP Demo，含proto文件、MQ消费者完整代码**。