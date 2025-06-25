
#### new
##### 安装goctl
> go-zero 的内置脚手架，是提升开发效率的一大利器，可以一键生成代码、文档、部署 k8s yaml、dockerfile 等
```
$ go install github.com/zeromicro/go-zero/tools/goctl@latest
$ goctl --version
```
##### 完整工程目录
```
mall // 工程名称
├── common // 通用库
│   ├── randx
│   └── stringx
├── go.mod
├── go.sum
└── service // 服务存放目录
    ├── afterSale
    │   ├── api
    │   └── model
    │   └── rpc
    ├── cart
    │   ├── api
    │   └── model
    │   └── rpc
    ├── order
    │   ├── api
    │   └── model
    │   └── rpc
    ├── pay
    │   ├── api
    │   └── model
    │   └── rpc
    ├── product
    │   ├── api
    │   └── model
    │   └── rpc
    └── user
        ├── api
        ├── cronjob
        ├── model
        ├── rmq
        ├── rpc
        └── script
```
##### 生成model
* 方式一
```
$ cd service/user/model
$ goctl model mysql ddl -src user.sql -dir . -c
```
* 方式二
```
$ cd service/user/model
goctl model mysql datasource --url="root:root@tcp(127.0.0.1:3306)/go-zero-shiro" --table="gzs_user" -c --dir . --custType=goZero,gorm --style=goZero
goctl model mysql datasource --url="root:root@tcp(127.0.0.1:3306)/go-zero-shiro" --table="gzs_user" --dir="./model" --custType=goZero,cache,gorm --style=goZero
```
##### 编写api文件
1. 生成api服务
```
$ cd service/user/api
$ goctl api go -api user.api -dir .
```
2. 添加Mysql配置
```
vim service/user/api/internal/config/config.go
```
3. 完善yaml配置
```
vim service/user/api/etc/user-api.yaml
```
4. 完善服务依赖
```
$ vim service/user/api/internal/svc/servicecontext.go
```
5. 填充登录逻辑
```
$ vim service/user/api/internal/logic/loginlogic.go
```
##### rpc服务编写
1. 生成rpc服务
```
$ cd service/user/rpc
$ goctl rpc protoc user.proto --go_out=./types --go-grpc_out=./types --zrpc_out=.
```
2. 添加配置及完善yaml配置项
```
$ vim service/user/rpc/internal/config/config.go
$ vim /service/user/rpc/etc/user.yaml
```
3. 添加资源依赖
```
$ vim service/user/rpc/internal/svc/servicecontext.go
```
4. 添加rpc逻辑
```
$ service/user/rpc/internal/logic/getuserlogic.go
```
##### 使用rpc
1. 添加UserRpc配置及yaml配置项
```
$ vim service/user/api/internal/config/config.go
$ vim service/user/api/etc/user-api.yaml
```
2. 添加依赖
```
$ vim service/user/api/internal/svc/servicecontext.go
```
3. 补充逻辑
```
$ vim /service/user/api/internal/logic/loginlogic.go
```