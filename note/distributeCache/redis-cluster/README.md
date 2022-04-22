1. 下载镜像
   ```
   docker pull redis:5.0.7
   ```
   
2. 创建 6 个 redis 节点, 3 主 3 从方式，从为了同步备份，主进行 slot 数据分片
   ```
   bash ./run.sh
   ```
   
3. 使用 redis 建立集群
   ```
   #进入到一个redis容器
   docker exec -it redis-7001 bash
   
   redis-cli --cluster create 192.168.1.89:7001 192.168.1.89:7002 192.168.1.89:7003 192.168.1.89:7004 192.168.1.89:7005 192.168.1.89:7006 --cluster-replicas 1
   #其中--cluster-replicas 1代表 一个master后有几个slave，1代表为1个slave节点
   #过程中会提示，输入 yes 继续    集群自动分配
   ```
   
4. 测试redis集群
   ```
   #随便进入某个 redis 容器
   docker exec -it redis-7002 /bin/bash
   #获取集群信息
   redis-cli -c -h 192.168.56.10 -p 7006 cluster info
   #获取集群节点
   redis-cli -c -h 192.168.56.10 -p 7006  cluster nodes
   #使用 redis-cli 的 cluster 方式进行连接
   redis-cli -c -h 192.168.56.10 -p 7006
   ```
   
5. 相关命令

   1. 集群创建

      ```
      redis-cli --cluster create 192.168.1.89:7001 192.168.1.89:7002 192.168.1.89:7003 192.168.1.89:7004 192.168.1.89:7005 192.168.1.89:7006 --cluster-replicas 1
      ```

      * 这样就创建了一个具有3个主节点和3个从节点的集群
      * --cluster-replicas 主节点有一个从节点
      * 虽然指定了每个主节点都有一个从节点，但哪个是7001的从节点，却是随机分配的，直到集群创建完毕，才能确定是7004、7005还是7006

   2. 增加主节点

      ```
      redis-cli --cluster add-node 192.168.1.89:7007 192.168.1.89:7001
      ```

      * 192.168.1.89:7007 要向集群添加新的节点

      * 192.168.1.89:7001 原集群中任意节点

      * 由于它还没有分配到 hash slots 哈希槽，所以它还没有数据，不会参与到从节点升级到主节点的选举中

      * 执行 resharding 指令来为它分配 hash slots，这会进入交互式命令行，由用户输入相关信息

        ```
        redis-cli --cluster reshard 127.0.0.1:7000
        
        只需要指定一个节点，redis会自动发现其他节点。
        
        How many slots do you want to move (from 1 to 16384)?
        target node id？
        from what nodes you want to take those keys？
        第一个问题需要填写，如1000. 
        第二个问题可以通过命令查看：redis-cli -p 7001 cluster nodes | grep myself
        第三个问题：all，这样会从每个节点上移动一部分 hash slots到新节点
        
        然后开始迁移，每迁移一个key就会输出一个点。
        ```

      * 执行下面的指令查看集群是否正常

        ```
        redis-cli --cluster check 192.168.1.89:7001
        ```

   3. 增加从节点

      ```
      redis-cli --cluster add-node 192.168.1.89:7008 192.168.1.89:7001 --cluster-slave
      ```

      * 192.168.1.89:7008 添加从节点

      * 192.168.1.89:7001 为集群中任意节点

      * 该指令与增加主节点语法一致，与添加主节点不同的是，显式指定了是从节点

      * 这会为该从节点随机分配一个主节点，优先从那些从节点数目最少的主节点中选取

      * 如果要在添加从节点时就为其指定主节点，需要指定master-id，执行下面的指令

        ```
        redis-cli --cluster add-node 192.168.1.89:7008 192.168.1.89:7001 --cluster-slave --cluster-master-id ***
        ```

   4. 删除节点

      ```
      redis-cli --cluster del-node 192.168.1.89:7001 <node-id>
      ```

      * 只能删除从节点或者空的主节点
      * 192.168.1.89:7001 为集群中任意节点
      * node-id为要删除的节点的id

   5. 获取集群信息

      ```
      redis-cli -c -h 192.168.1.89 -p 7006 cluster info
      ```

   6. 获取集群节点

      ```
      redis-cli -c -h 192.168.1.89 -p 7006  cluster nodes;
      ```

   7. 使用 redis-cli 的 cluster 方式进行连接

      ```
      redis-cli -c -h 192.168.1.89 -p 7006
      ```

      