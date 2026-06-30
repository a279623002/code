# go-zero 订单管理微服务

基于 [go-zero](https://github.com/zeromicro/go-zero) 框架构建的订单管理微服务示例，集成 gorm、etcd 服务注册发现、Swagger 文档。

## 项目结构

```
go-zero-order/
├── common/
│   ├── gorm/               # gorm 通用封装（日志适配 go-zero logx）
│   └── etcd/               # etcd 服务注册/发现示例
├── deploy/
│   ├── docker-compose.yaml # MySQL + etcd 一键启动
│   └── mysql/init.sql      # 订单表初始化 SQL
├── order/
│   ├── api/                # HTTP API 网关（端口 8888）
│   │   ├── etc/order.yaml
│   │   ├── internal/
│   │   │   ├── config/
│   │   │   ├── handler/    # HTTP handler + Swagger 文档
│   │   │   ├── logic/
│   │   │   ├── svc/
│   │   │   └── types/
│   │   └── order.go
│   └── rpc/                # gRPC 订单服务（端口 8080）
│       ├── etc/order.yaml
│       ├── internal/
│       │   ├── config/
│       │   ├── logic/
│       │   ├── model/      # gorm 订单模型
│       │   ├── server/
│       │   └── svc/
│       ├── order.proto
│       ├── types/order/    # protobuf 生成代码
│       └── order.go
├── go.mod
└── README.md
```

## 技术栈

- **Web 框架**：go-zero
- **ORM**：gorm + mysql driver
- **RPC**：gRPC + protobuf
- **服务注册/发现**：etcd
- **文档**：Swagger UI 3.0

## 前置依赖

- Go 1.22+
- Docker & Docker Compose（推荐） 或 本地 MySQL + etcd

## 快速开始

### 1. 启动基础设施

使用 Docker Compose 启动 MySQL 与 etcd：

```bash
cd deploy
docker-compose up -d
```

验证：

```bash
docker ps
# 应看到 go-zero-mysql 与 go-zero-etcd
```

### 2. 初始化数据库

Docker Compose 已自动挂载 `deploy/mysql/init.sql` 创建数据库与表。若使用本地 MySQL，请手动执行该 SQL。

### 3. 编译服务

```bash
# 编译 RPC 服务
go build -o order-rpc.exe ./order/rpc

# 编译 API 网关
go build -o order-api.exe ./order/api
```

### 4. 启动 RPC 服务

```bash
./order-rpc.exe -f order/rpc/etc/order.yaml
```

启动成功后，RPC 服务会注册到 etcd（Key: `order-rpc`）。

### 5. 启动 API 网关

另开一个终端：

```bash
./order-api.exe -f order/api/etc/order.yaml
```

API 网关会通过 etcd 发现 `order-rpc` 服务。

## API 接口

| 方法 | 路径 | 说明 |
|------|------|------|
| POST | `/api/order` | 创建订单 |
| GET | `/api/order/:id` | 查询订单详情 |
| GET | `/api/order/list` | 订单列表（支持 user_id/page/page_size） |
| PUT | `/api/order/:id/status` | 更新订单状态 |

## Swagger 文档

服务启动后访问：

```
http://127.0.0.1:8888/swagger
```

Swagger JSON 地址：

```
http://127.0.0.1:8888/swagger/swagger.json
```

## 配置说明

### order/rpc/etc/order.yaml

```yaml
Name: order-rpc
ListenOn: 0.0.0.0:8080

DB:
  DataSource: root:123456@tcp(127.0.0.1:3306)/go_zero_order?charset=utf8mb4&parseTime=True&loc=Local
  MaxIdleConns: 10
  MaxOpenConns: 100
  ConnMaxLifetime: 3600

Etcd:
  Hosts:
    - 127.0.0.1:2379
  Key: order-rpc
```

### order/api/etc/order.yaml

```yaml
Name: order-api
Host: 0.0.0.0
Port: 8888

OrderRpc:
  Etcd:
    Hosts:
      - 127.0.0.1:2379
    Key: order-rpc
```

## 订单状态

| 值 | 状态 |
|----|------|
| 1 | 待支付 |
| 2 | 已支付 |
| 3 | 已发货 |
| 4 | 已完成 |
| 5 | 已取消 |

## 调用示例

### 创建订单

```bash
curl -X POST http://127.0.0.1:8888/api/order \
  -H "Content-Type: application/json" \
  -d '{"user_id":10001,"total_amount":199.99,"remark":"测试订单"}'
```

### 查询订单

```bash
curl http://127.0.0.1:8888/api/order/1
```

### 订单列表

```bash
curl "http://127.0.0.1:8888/api/order/list?user_id=10001&page=1&page_size=10"
```

### 更新订单状态

```bash
curl -X PUT http://127.0.0.1:8888/api/order/1/status \
  -H "Content-Type: application/json" \
  -d '{"status":2}'
```

## 重新生成 protobuf 代码

修改 `order/rpc/order.proto` 后，执行：

```bash
cd order/rpc
protoc --go_out=. --go_opt=paths=source_relative \
  --go-grpc_out=. --go-grpc_opt=paths=source_relative \
  order.proto
```

生成后请将 `order.pb.go` 与 `order_grpc.pb.go` 移动到 `types/order/` 目录。

## 注意事项

- 默认数据库账号为 `root` / `123456`，生产环境请修改。
- etcd 默认无认证，生产环境请启用 TLS 与认证。
- go-zero 的 etcd 服务发现由 `zrpc` 自动完成，无需手动调用 `common/etcd` 中的代码；该目录仅作为显式示例。
