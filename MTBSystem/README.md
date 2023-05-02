#### env
1. go-micro v2
3. docker
2. mysql(in docker)
4. supervisor(management services in docker)
5. go-vesion 1.14.1 (in docker)
6. go-env mod (in docker)

##### micro install
```
$ go get github.com/micro/micro/v2
$ go get -u github.com/golang/protobuf/proto
$ go get -u github.com/golang/protobuf/protoc-gen-go
$ go get github.com/micro/micro/v2/cmd/protoc-gen-micro
```
##### proto
```
# cd proto
$ protoc --proto_path=. --go_out=:. --micro_out=. srv.proto
```
##### install project
```
# cd project
# build docker
$ sudo bash ./ctrl.sh build
# run docker
$ sudo bash ./ctrl.sh run
# init sql, supervisor
$ sudo bash ./ctrl.sh init conf
# chmod conf
$ sudo bash ./ctrl.sh init chmod
# start docker
$ sudo bash ./ctrl.sh start
# login docker
$ sudo bash ./ctrl.sh login
# cd project(docker)
# install and run srv
$ sudo bash ./build_local.sh all
```
##### create srv
1. make proto code
2. set handler
3. build srv successfully
4. set supervisor.conf
5. ./ctrl.sh init conf
6. ./ctrl.sh init chmod
7. ./ctrl.sh start
8. ./ctrl.sh login
9. cd project(docker)
10. ./build.sh srv

#### swagger
| http://127.0.0.1:18082/swagger/index.html
