FROM ubuntu:14.04

MAINTAINER shiro "970979353@qq.com"

ADD http://mirrors.163.com/.help/sources.list.trusty /etc/apt/sources.list

COPY conf/redis.conf /etc/redis/6379.conf
COPY conf/consul.json /etc/consul/consul.json
COPY supervisor/*.conf /etc/supervisor/conf.d/

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
    wget https://studygolang.com/dl/golang/go1.20.3.linux-amd64.tar.gz && \
    tar zxf go1.20.3.linux-amd64.tar.gz && \
    mkdir -p /data/services/ && \
    mv go/ /data/services/go && \
    rm -rf go1.20.3.linux-amd64.tar.gz consul_0.9.0_linux_amd64.zip && \
    echo "export GO111MODULE=on" >> ~/.bashrc && \
    echo "export GOPROXY=https://goproxy.cn,direct" >> ~/.bashrc && \
    echo "export GOROOT=/data/services/go" >> ~/.bashrc && \
    echo "export PATH=$PATH:/data/services/go/bin" >> ~/.bashrc && \
    sudo ln -s /data/services/go/bin/* /usr/bin/ && \
    go env -w GOPROXY=https://goproxy.cn,direct

EXPOSE 3306 8500 6379 8082 8083 8084 8085 5324 9999

CMD chown -R mysql:mysql /var/lib/mysql && \
    service mysql start && \
    supervisord -c /etc/supervisor/supervisord.conf -n