##### docker-composer 搭建 nginx+php

###### 目录

```
├── docker-compose.yml
├── Dockerfile
├── go.mod
├── main
└── main.go
```

###### go Dockerfile

```
FROM golang:alpine

# 以下设置后还是会会出现8小时的时差
# 所以在docker-compose 设置时区
#RUN ln -sf /usr/share/zoneinfo/Asia/Shanghai /etc/localtime
#    && echo "Asia/Shanghai" > /etc/timezone

# 设置环境变量
ENV GO111MODULE on
ENV CGO_ENABLED 0
ENV GOPROXY https://goproxy.io,direct

WORKDIR /go/src/pachong/demo-docker

# 将当前目录挂载到容器中的工作目录
# 在docker-compose 里挂载了目录，方便部署，就不需要在容器里复制了
# COPY . .
```

###### docker-compose.yml

```
version: '3'
services:
  go:
  	# $GOPATH/src/goStudy/pachong/demo-docker/
    build: .
    container_name: shiro-go
    restart: always
    ports:
      - '9099:9099'
    volumes:
      - .:/go/src/pachong/demo-docker #挂载项目目录
      - /etc/localtime:/etc/localtime:ro #本地时间 ro read only
      - /etc/timezone:/etc/timezone:ro #本地时区
    command: /bin/sh -c 'cd /go/src/pachong/demo-docker/ && go mod tidy && go build main.go && ./main' #初始运行命令
```

##### 测试

```
//main.go
package main

import (
	"fmt"
	"github.com/shopspring/decimal"
	"log"
	"net/http"
)

func main() {
	a, _ := decimal.NewFromString("23.23")
	b, _ := decimal.NewFromString("23.2333")
	log.Println(a.Sub(b).Abs().String())
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "hello233")
	})
	http.ListenAndServe(":9099", nil)

}
// 可以在控制台查看输出
// 也可访问127.0.0.1：9099
```

