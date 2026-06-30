# go-zero 用户管理微服务

基于 [go-zero](https://github.com/zeromicro/go-zero) 框架构建的用户管理微服务，与 [go-zero-order](../go-zero-order) 订单服务通过 etcd + gRPC 解耦集成。

## 服务关系

```
user-api (端口 8889)
    │
    ├─ 用户表 (go_zero_user.users)  ── 用户自身数据
    │
    └─ order-rpc (通过 etcd 发现)   ── 查询/创建订单
```

**解耦点**：
- user 服务不直接访问 order 数据库。
- user 服务通过 `go-zero-order/order/rpc/orderclient` 调用 order-rpc，仅依赖订单服务的 RPC 契约。
- 订单服务可以独立部署、升级，不影响用户服务。

## 项目结构

```
go-zero-user/
├── common/gorm/            # gorm 通用封装
├── deploy/
│   ├── docker-compose.yaml # MySQL + etcd（同时初始化 order/user 库）
│   └── mysql/init.sql      # 用户表初始化 SQL
├── user/api/
│   ├── etc/user.yaml
│   ├── user.go
│   └── internal/
│       ├── config/
│       ├── handler/        # HTTP handler + Swagger
│       ├── logic/          # 用户逻辑 + 调用 order-rpc
│       ├── model/          # gorm 用户模型
│       ├── svc/
│       └── types/
├── go.mod
└── README.md
```

## 快速开始

### 1. 启动基础设施

```bash
cd deploy
docker-compose up -d
```

### 2. 编译

```bash
# 先编译 order 服务（被 user 依赖）
cd ../go-zero-order
go build -o order-rpc.exe ./order/rpc
go build -o order-api.exe ./order/api

# 再编译 user 服务
cd ../go-zero-user
go build -o user-api.exe ./user/api
```

### 3. 启动服务

```bash
# 终端 1：启动 order-rpc
cd go-zero-order
./order-rpc.exe -f order/rpc/etc/order.yaml

# 终端 2：启动 user-api
cd go-zero-user
./user-api.exe -f user/api/etc/user.yaml
```

## API 接口

| 方法 | 路径 | 说明 |
|------|------|------|
| POST | `/api/user` | 创建用户 |
| GET | `/api/user/:id` | 查询用户 |
| GET | `/api/user/:id/orders` | 查询用户订单（调用 order-rpc） |
| POST | `/api/user/:id/orders` | 为用户下单（调用 order-rpc） |

## Swagger 文档

服务启动后访问：

```
http://127.0.0.1:8889/swagger
```

## 调用示例

### 创建用户

```bash
curl -X POST http://127.0.0.1:8889/api/user \
  -H "Content-Type: application/json" \
  -d '{"username":"张三","phone":"13800138000","email":"zhangsan@example.com"}'
```

### 查询用户

```bash
curl http://127.0.0.1:8889/api/user/1
```

### 为用户下单

```bash
curl -X POST http://127.0.0.1:8889/api/user/1/orders \
  -H "Content-Type: application/json" \
  -d '{"total_amount":299.99,"remark":"用户下单"}'
```

### 查询用户订单

```bash
curl "http://127.0.0.1:8889/api/user/1/orders?page=1&page_size=10"
```

## 配置说明

[user/api/etc/user.yaml](user/api/etc/user.yaml) 关键配置：

```yaml
DB:
  DataSource: root:123456@tcp(127.0.0.1:3306)/go_zero_user?...

OrderRpc:
  Etcd:
    Hosts:
      - 127.0.0.1:2379
    Key: order-rpc
```

## 解耦说明

- `go.mod` 中通过 `replace go-zero-order => ../go-zero-order` 引入订单服务模块。
- user 服务仅使用 `go-zero-order/order/rpc/orderclient` 与 `go-zero-order/order/rpc/types/order`。
- 运行时发现 order-rpc 地址，调用其 gRPC 接口，避免共享数据库或直接耦合实现。
