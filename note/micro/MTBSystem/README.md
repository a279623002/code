protoc --proto_path=. --go_out=:. --micro_out=. user-sr
v.proto

github.com/micro/micro/v2/cmd/protoc-gen-micro

build_local.sh 是在docker执行的脚本
ctrl.sh 是构建容器、环境的脚本