FROM ubuntu:14.04

MAINTAINER shiro "970979353@qq.com"

ADD http://mirrors.163.com/.help/sources.list.trusty /etc/apt/sources.list

COPY conf/redis.conf /etc/redis/6379.conf
COPY conf/consul.json /etc/consul/consul.json
COPY supervisor/*.conf /etc/supervisor/conf.d/
# COPY 复制指令，从上下文目录中复制文件或者目录到容器里指定路径
# ADD 指令和 COPY 的使用格类似（同样需求下，官方推荐使用 COPY）。功能也类似，不同之处 ADD 会自动复制并解压到 <目标路径>

RUN apt-get update && \
    apt-get -y install build-essential && \
    apt-get -y install openssh-server && \
    apt-get -y install libssl-dev && \
    apt-get -y install git && \
    apt-get -y install vim && \
    apt-get -y install wget && \
    apt-get -y install curl && \
    apt-get -y install unzip

RUN apt-get -y install supervisor && \
    apt-get -y install redis-server && \
    apt-get -y install mysql-server && \
    apt-get -y install mysql-client && \
    mkdir -p /data/services/consul-0.9/bin/ && \
    wget https://releases.hashicorp.com/consul/0.9.0/consul_0.9.0_linux_amd64.zip && \
    unzip consul_0.9.0_linux_amd64.zip && \
    mv ./consul /usr/local/bin/ && \
    mkdir /data/consul/ && \
    mkdir -p /data/logs/gologs/ && \
    mysql_install_db && \
    update-rc.d -f mysql defaults && \
    wget https://studygolang.com/dl/golang/go1.14.1.linux-amd64.tar.gz && \
    tar zxf go1.14.1.linux-amd64.tar.gz && \
    mkdir -p /data/services/ && \
    mv go/ /data/services/go && \
    rm -rf go1.14.1.linux-amd64.tar.gz consul_0.9.0_linux_amd64.zip && \
    echo "export GO111MODULE=on" >> ~/.bashrc && \
    echo "export GOPROXY=https://goproxy.cn,direct" >> ~/.bashrc && \
    echo "export GOROOT=/data/services/go" >> ~/.bashrc && \
    echo "export PATH=$PATH:/data/services/go/bin" >> ~/.bashrc && \
    sudo ln -s /data/services/go/bin/* /usr/bin/

# 仅仅只是声明端口
EXPOSE 3306 8500 6379 8082 8083 8084 8085 5324 9999

CMD chown -R mysql:mysql /var/lib/mysql && \
    service mysql start && \
    supervisord -c /etc/supervisor/supervisord.conf -n
# CMD 类似于 RUN 指令，用于运行程序，但二者运行的时间点不同
# CMD 在docker run 时运行
# RUN 是在 docker build

# build build -t shiro/go-env .
# run   docker run --name=$Container -p 13306:3306 -p 16379:6379 -p 18500:8500 -p 19999:9999 -p 15324:5324 -p 18082:8082 -p 18083:8083 -p 18084:8084 -p 18085:8085 -p 18090:8090 -d -v `pwd`:/data/deploy/$ProjectName shiro/go-env