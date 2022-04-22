#创建6个实例【端口为：7001----7006】
for port in $(seq 7001 7006); \
do \
mkdir -p ./data/node-${port}/conf
touch ./data/node-${port}/conf/redis.conf
cat << EOF >./data/node-${port}/conf/redis.conf
port  ${port}
cluster-enabled yes
cluster-config-file nodes.conf
cluster-node-timeout 5000
cluster-announce-ip 192.168.1.89
cluster-announce-port ${port}
cluster-announce-bus-port 1${port}
appendonly yes
EOF
docker run -p ${port}:${port} -p 1${port}:1${port} --name redis-${port} \
-v /home/siro/文档/code/note/dynamicProgramming/redis-cluster/data/node-${port}/data:/data \
-v /home/siro/文档/code/note/dynamicProgramming/redis-cluster/data/node-${port}/conf/redis.conf:/etc/redis/redis.conf \
-d redis:5.0.7 redis-server /etc/redis/redis.conf; \
done

#停止【可选】
#docker stop $(docker ps -a |grep redis-700 | awk '{ print $1}')
#删除【可选】
#docker rm $(docker ps -a |grep redis-700 | awk '{ print $1}')
