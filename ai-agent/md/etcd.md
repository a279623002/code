# etcd 面试笔记

> etcd 是分布式键值存储系统，基于 Raft 协议实现强一致性。常用于服务发现、配置中心、分布式锁、Leader 选举。调度系统和 go-zero 微服务都用 etcd 做服务注册发现。

---

## 一、etcd 是什么？

**一句话**：etcd 是一个高可用、强一致的分布式 **key-value 存储系统**，使用 Go 语言编写，基于 Raft 共识算法。

**核心特点**：
- 强一致性：基于 Raft，读写都能保证线性一致性
- 高可用：多节点集群，自动 Leader 选举
- Watch 机制：监听 key 变化，实时通知
- TTL：支持 key 过期
- 事务：支持多 key 原子操作
- 安全：支持 TLS 认证

---

## 二、架构图

```
┌─────────────────────────────────────────────┐
│              etcd Cluster                    │
│                                              │
│   ┌─────────┐     ┌─────────┐     ┌─────────┐ │
│   │  Node 1 │◄───►│  Node 2 │◄───►│  Node 3 │ │
│   │ Leader  │     │Follower │     │Follower │ │
│   │         │     │         │     │         │ │
│   │ 处理写请求 │     │复制日志   │     │复制日志   │ │
│   └─────────┘     └─────────┘     └─────────┘ │
│        ▲                                     │
│        │         心跳 / 日志复制              │
└────────┼─────────────────────────────────────┘
         │
    ┌────┴────┐
    │  Client │
    └─────────┘
```

---

## 三、核心概念

| 概念 | 说明 |
|---|---|
| **Raft** | 分布式一致性算法，负责 Leader 选举和日志复制 |
| **Leader** | 处理所有写请求，协调日志复制 |
| **Follower** | 接收并复制 Leader 日志 |
| **Candidate** | Leader 失效时，Follower 转为 Candidate 参与选举 |
| **Term** | 任期号，单调递增，用于区分不同 Leader |
| **Log Entry** | 日志条目，包含操作指令和任期号 |
| **WAL** | 预写日志，保证数据持久化 |
| **MVCC** | 多版本并发控制，每个 key 有多个历史版本 |
| **Revision** | 全局版本号，每次写操作递增 |

---

## 四、Raft 协议原理

### 1. Leader 选举

```
1. 所有节点初始为 Follower
2. 选举超时（election timeout）后没收到 Leader 心跳，变成 Candidate
3. Candidate 增加 Term，给自己投票，向其他节点请求投票
4. 获得半数以上选票，成为 Leader
5. Leader 周期性发送心跳，维持权威
```

**选举规则**：
- 一个 Term 内，一个节点只能投一票
- Candidate 的日志必须至少和自己一样新，才能获得选票
- 如果同时多个 Candidate 竞争，可能 split vote，重新选举

### 2. 日志复制

```
Client ──写请求──► Leader
Leader ──AppendEntries RPC──► Followers
Followers ──确认──► Leader
Leader ──提交──► 应用到状态机
Leader ──响应 Client
```

**提交条件**：
- 日志被半数以上节点复制成功
- 才能提交并返回给客户端

### 3. 安全性

- **Election Restriction**：只有日志足够新的节点才能当选 Leader
- **Commit Restriction**：Leader 只能提交自己任期内的日志

---

## 五、etcd 使用方式

### 1. 命令行操作

```bash
# 写入
etcdctl put /config/db_host "127.0.0.1"

# 读取
etcdctl get /config/db_host

# 删除
etcdctl del /config/db_host

# 监听变化
etcdctl watch /config/db_host

# 租约（TTL）
etcdctl lease grant 60
# lease 32695410dcc0ca06
etcdctl put /lock/serviceA "locked" --lease=32695410dcc0ca06

# 事务
etcdctl txn -i
```

### 2. Go 客户端示例

```go
package main

import (
    "context"
    "fmt"
    "time"

    clientv3 "go.etcd.io/etcd/client/v3"
)

func main() {
    cli, err := clientv3.New(clientv3.Config{
        Endpoints:   []string{"localhost:2379"},
        DialTimeout: 5 * time.Second,
    })
    if err != nil {
        panic(err)
    }
    defer cli.Close()

    ctx := context.Background()

    // 写入
    _, err = cli.Put(ctx, "foo", "bar")
    if err != nil {
        panic(err)
    }

    // 读取
    resp, err := cli.Get(ctx, "foo")
    if err != nil {
        panic(err)
    }
    for _, ev := range resp.Kvs {
        fmt.Printf("%s : %s\n", ev.Key, ev.Value)
    }

    // 监听
    watchCh := cli.Watch(ctx, "foo")
    for wresp := range watchCh {
        for _, ev := range wresp.Events {
            fmt.Printf("Type: %s Key:%s Value:%s\n", ev.Type, ev.Kv.Key, ev.Kv.Value)
        }
    }
}
```

---

## 六、Watch 机制

```
Client ──Watch("/config")──► etcd
                              │
                              │ key 变化
                              ↓
Client ◄──Event(PUT/DELETE)─── etcd
```

**特点**：
- 基于 gRPC 长连接
- 可以监听单个 key 或前缀
- 支持历史版本回放

```go
// 监听前缀
rch := cli.Watch(ctx, "/services/", clientv3.WithPrefix())
for wresp := range rch {
    for _, ev := range wresp.Events {
        fmt.Printf("%s %q : %q\n", ev.Type, ev.Kv.Key, ev.Kv.Value)
    }
}
```

---

## 七、分布式锁

### 基于 etcd 的分布式锁

```go
import (
    "go.etcd.io/etcd/client/v3/concurrency"
)

func lock(cli *clientv3.Client) {
    s, err := concurrency.NewSession(cli)
    if err != nil {
        panic(err)
    }
    defer s.Close()

    mu := concurrency.NewMutex(s, "/lock/my-lock/")

    // 加锁
    if err := mu.Lock(context.Background()); err != nil {
        panic(err)
    }

    // 执行业务
    doSomething()

    // 释放锁
    if err := mu.Unlock(context.Background()); err != nil {
        panic(err)
    }
}
```

**原理**：
- 利用 etcd 创建带 TTL 的 key
- 创建成功表示获取锁
- 释放锁时删除 key
- 通过 Watch 监听锁释放，实现阻塞等待

---

## 八、服务注册与发现

```
服务启动：
    serviceA ──put /services/user-service/192.168.1.10:8080 {"weight":10}──► etcd
    serviceB ──put /services/user-service/192.168.1.11:8080 {"weight":10}──► etcd

客户端发现：
    Client ──get /services/user-service --prefix──► etcd
    etcd 返回所有节点地址

动态感知：
    Client ──watch /services/user-service --prefix──► etcd
    服务上下线实时通知
```

**go-zero 中的应用**：
- rpc 服务启动时向 etcd 注册地址
- api 网关从 etcd 拉取 rpc 地址
- 通过 watch 实现服务实例动态更新

---

## 九、etcd vs ZooKeeper vs Consul

| 特性 | etcd | ZooKeeper | Consul |
|---|---|---|---|
| 协议 | Raft | ZAB | Raft |
| 一致性 | 强一致 | 强一致 | 强一致 |
| 接口 | HTTP/gRPC | 自定义协议 | HTTP/DNS |
| Watch | 基于 gRPC 流 | 一次性触发 | 支持 |
| 多数据中心 | 不支持 | 不支持 | 支持 |
| 健康检查 | 无 | 无 | 有 |
| 适用场景 | K8s、服务发现、配置中心 | 分布式协调 | 服务网格、健康检查 |

---

## 十、面试高频问题

### Q1：etcd 和 Redis 的区别？

**答**：
- etcd 强调一致性和高可用，基于 Raft；Redis 强调性能，单机/主从/Cluster
- etcd 适合配置、服务发现、锁；Redis 适合缓存、计数器、会话
- etcd 写操作需要半数以上确认；Redis 主节点直接响应

### Q2：Raft 是怎么选举 Leader 的？

**答**：Follower 在选举超时后变成 Candidate，增加 Term 并发起投票。获得半数以上选票且日志足够新的节点成为 Leader。Leader 定期发送心跳维持权威。

### Q3：etcd 如何保证一致性？

**答**：
- 所有写请求走 Leader
- Leader 通过 Raft 日志复制到半数以上 Follower 才提交
- 读请求默认走 Leader，保证线性一致性

### Q4：etcd 的 Watch 有什么用？

**答**：监听 key 或前缀的变化，实时推送事件。用于配置中心热更新、服务发现动态感知。

### Q5：etcd 集群为什么建议奇数节点？

**答**：Raft 需要半数以上节点存活。3 节点可容忍 1 个故障，5 节点可容忍 2 个故障。偶数节点并不能提高容错能力，反而增加选举分裂概率。

---

## 十一、一句话总结

- **etcd**：基于 Raft 的强一致分布式 KV 存储
- **Raft**：Leader 选举 + 日志复制 + 半数提交
- **Watch**：实时监听 key 变化，配置中心和服务发现核心
- **分布式锁**：基于 TTL + Watch 实现
- **服务发现**：服务注册地址，客户端 Watch 动态感知
