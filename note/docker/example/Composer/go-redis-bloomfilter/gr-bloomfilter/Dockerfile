FROM golang:alpine

# 以下设置后还是会会出现8小时的时差
# 所以在docker-compose 设置时区
#RUN ln -sf /usr/share/zoneinfo/Asia/Shanghai /etc/localtime
#    && echo "Asia/Shanghai" > /etc/timezone

# 设置环境变量
ENV GO111MODULE on
ENV CGO_ENABLED 0
ENV GOPROXY https://goproxy.io,direct

WORKDIR /files/go/gr-bloomfilter

# 将当前目录挂载到容器中的工作目录
# 在docker-compose 里挂载了目录，方便部署，就不需要在容器里复制了
# COPY . .
