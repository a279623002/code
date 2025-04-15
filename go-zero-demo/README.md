#### docker
1. 访问AtomHub 可信镜像中心（https://hub.atomgit.com/） 获取镜像拉取
2. docker pull hub.atomgit.com/amd64/consul:1.13.9
3. cd ./go-zero-demo/docker && docker-compose -f docker-compose.yaml up -d
* 启动了4个consul，其中consul1 是主节点，consul2、consul3 是子节点。consul4是提供ui服务的
4. ACL是Consul用来控制访问API与data的。 首先，我们创建一个uuid，也可以使用consul自带的生成key的功能
* docker exec consul1 consul keygen
* 使用这个token，我们创建一个acl.json和acl_client.json文件,分别给与server节点和client节点使用

#### 安装工具
* go install github.com/golang/protobuf/protoc-gen-go@latest
* go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
```
wget https://github.com/protocolbuffers/protobuf/releases/download/v27.3/protoc-27.3-linux-x86_64.zip
unzip protoc-27.3-linux-x86_64.zip
mv bin/protoc /usr/local/bin/
protoc --version
```
* go install github.com/zeromicro/go-zero/tools/goctl@latest
* go/bin 包含protoc-gen-go、protoc-gen-go-grpc、goctl三个工具

##### rpc服务
1. 编写proto文件，定义服务接口
2. 使用goctl工具生成rpc服务代码
```
goctl rpc protoc ./proto/*.proto --go_out=./types --go-grpc_out=./types --zrpc_out=.
```

#### api
1. 编写api.go文件，定义api接口
2. 使用goctl工具生成api代码
```
goctl api go -api ./api/*.api -dir ./  --style=goZero
```
3. demo-api.yaml文件配置rpc服务
```
#注意这个名字和config文件中的名字是对应的
DemoRpc:
  Consul:
  Host: 127.0.0.1:8500
  Key: demo.rpc
```
4. config.go 定义rpc服务
5. internal/svc/serviceContext.go 引入grpc服务
6. logic 调用grpc服务
7. main入口import导入包，完成初始化
```
_ "github.com/zeromicro/zero-contrib/zrpc/registry/consul"
```