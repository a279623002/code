# Kubernetes（K8s）面试笔记

> 项目背景：调度系统里通过 k8s 让训练任务在目标算力卡上运行。k8s 是容器编排核心工具。

---

## 一、Kubernetes 是什么？

**一句话**：K8s 是一个**容器编排平台**，帮你自动部署、扩缩容、管理容器化应用。

**生活例子**：
- 容器（Docker）= 一个个快递包裹
- K8s = 物流调度中心，决定包裹放在哪辆车、走哪条路、坏了怎么换

---

## 二、架构图与核心组件

```
┌─────────────────────────────────────────────────────────────┐
│                      控制平面（Control Plane）               │
│  ┌──────────┐  ┌──────────┐  ┌──────────┐  ┌────────────┐  │
│  │ API Server│  │ etcd    │  │Scheduler │  │Controller  │  │
│  │           │  │         │  │          │  │Manager     │  │
│  └─────┬─────┘  └─────────┘  └─────┬────┘  └─────┬──────┘  │
└────────┼───────────────────────────┼─────────────┼─────────┘
         │                           │             │
         └───────────────────────────┴─────────────┘
                              │
                    ┌─────────┴─────────┐
                    │     Worker Node    │
                    │  ┌─────────────┐  │
                    │  │   kubelet   │  │
                    │  │ kube-proxy  │  │
                    │  │  Container  │  │
                    │  │  Runtime    │  │
                    │  └──────┬──────┘  │
                    │         │         │
                    │    ┌────┴────┐    │
                    │    │   Pod   │    │
                    │    │ [容器]  │    │
                    │    └─────────┘    │
                    └───────────────────┘
```

### 控制平面组件（Master）

| 组件 | 作用 | 面试一句话 |
|---|---|---|
| **API Server** | 所有请求的入口 | "K8s 的前台接待，所有操作都经过它" |
| **etcd** | 分布式键值存储，所有资源（Pod、Service、Deployment、配置、密钥、权限）,保存集群状态 | "K8s 的数据库",只有apiserver能写入 |
| **Scheduler** | 负责把 Pod 调度到合适的 Node | "HR，决定新员工去哪个部门" |
| **Controller Manager** | 控制器管理器,持续调谐：当前状态 ↔ 期望状态 | "监工，发现实际状态和配置不一样就修复" |

### 工作节点组件（Node）

| 组件 | 作用 |
|---|---|
| **kubelet** | 节点代理，每个 Node 上的"小管家"，对接 apiserver，管理本机所有 Pod，汇报状态 |
| **kube-proxy** | 负责 Service 的网络转发和负载均衡 |
| **Container Runtime** | 真正运行容器，比如 containerd、Docker |

---

## 三、核心概念解析

### 1. Pod

**一句话**：Pod 是 K8s 的**最小调度单位**，一个 Pod 里可以跑一个或多个容器。

**为什么需要 Pod**：
- 同一个 Pod 里的容器**共享网络**（localhost）和**存储卷**
- 适合"主容器 + 辅助容器"场景，比如应用容器 + 日志收集容器

```yaml
# metadata：元数据（名字、命名空间、标签、注解）
# spec：期望规格，你想要资源长成什么样
# status：实际状态，集群真实运行情况（只读，不能手动改）
apiVersion: v1
kind: Pod
metadata:
  name: nginx-pod
spec:
  containers:
    - name: nginx
      image: nginx:1.25
      ports:
        - containerPort: 80
```

**创建并查看**：
```bash
kubectl apply -f nginx-pod.yaml
kubectl get pods
kubectl describe pod nginx-pod
kubectl logs nginx-pod
kubectl delete pod nginx-pod
# 查当前命名空间所有Pod，带IP和节点
kubectl get pods -o wide
# 查kube-system命名空间全部Pod详情
kubectl -n kube-system get pods -owide
# 查看所有节点详细硬件与运行时
kubectl get nodes -o wide
# 查看deployment使用的镜像版本
kubectl get deploy -o wide
# 全命名空间pod宽表
kubectl get pods -A -owide
# 按标签查pod并展示标签
kubectl get pods -l app=order --show-labels
# 实时监控+宽输出
kubectl get deploy -owide -w
# 导出yaml不创建
kubectl create job test --image=busybox --dry-run=client -o yaml
```

### 2. Deployment

**一句话**：管理 Pod 的"控制器"，负责滚动更新、副本数维持、故障自愈。

```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: nginx-deploy
spec:
  replicas: 3
  selector:
    matchLabels:
      app: nginx
  template:
    metadata:
      labels:
        app: nginx
    spec:
      containers:
        - name: nginx
          image: nginx:1.25
          ports:
            - containerPort: 80
```

**常用命令**：
```bash
kubectl apply -f nginx-deploy.yaml
kubectl get deployments
kubectl get pods -l app=nginx
kubectl scale deployment nginx-deploy --replicas=5
kubectl rollout status deployment/nginx-deploy
kubectl rollout history deployment/nginx-deploy
kubectl rollout undo deployment/nginx-deploy   # 回滚

```

### 3. Service

**一句话**：给一组 Pod 提供一个**稳定访问入口**，自动做负载均衡。

**三种类型**：

| 类型 | 作用 |
|---|---|
| **ClusterIP** | 集群内部访问（默认） |
| **NodePort** | 通过每个 Node 的端口暴露服务 |
| **LoadBalancer** | 云厂商负载均衡器暴露服务 |

```yaml
apiVersion: v1
kind: Service
metadata:
  name: nginx-svc
spec:
  selector:
    app: nginx
  ports:
    - port: 80
      targetPort: 80
  type: ClusterIP
```

### 4. ConfigMap / Secret

| 类型 | 用途 |
|---|---|
| **ConfigMap** | 存非敏感配置，如配置文件、环境变量 |
| **Secret** | 存敏感信息，如密码、Token，Base64 编码 |

```bash
# 创建 ConfigMap
kubectl create configmap app-config \
  --from-literal=ENV=prod \
  --from-literal=LOG_LEVEL=info

# 创建 Secret
kubectl create secret generic db-secret \
  --from-literal=password=123456
```

### 5. Volume / PVC

**一句话**：解决容器重启数据丢失问题。

| 类型 | 特点 |
|---|---|
| **emptyDir** | 临时目录，Pod 删除就没了 |
| **hostPath** | 挂载宿主机目录 |
| **PVC** | 向 StorageClass 申请持久化存储 |

```yaml
apiVersion: v1
kind: PersistentVolumeClaim
metadata:
  name: data-pvc
spec:
  accessModes:
    - ReadWriteOnce
  resources:
    requests:
      storage: 10Gi
```

### 6. Namespace

**一句话**：资源隔离的"虚拟集群"。

```bash
kubectl create ns dev
kubectl get pods -n dev
kubectl get all --all-namespaces
```

---

## 四、常用命令速查表

### 增删改查

| 命令 | 作用 |
|---|---|
| `kubectl apply -f xxx.yaml` | 创建/更新资源 |
| `kubectl delete -f xxx.yaml` | 删除资源 |
| `kubectl get pods` | 查看 Pod |
| `kubectl get pods -o wide` | 查看 Pod 详情（节点/IP） |
| `kubectl describe pod xxx` | 查看事件和状态 |
| `kubectl logs xxx` | 查看日志 |
| `kubectl logs -f xxx` | 实时查看日志 |
| `kubectl exec -it xxx -- /bin/sh` | 进入容器 |

### 调度和扩缩容

| 命令 | 作用 |
|---|---|
| `kubectl scale deploy xxx --replicas=5` | 扩缩容 |
| `kubectl rollout status deploy xxx` | 查看滚动更新状态 |
| `kubectl rollout undo deploy xxx` | 回滚 |
| `kubectl cordon node1` | 标记节点不可调度 |
| `kubectl drain node1` | 驱逐节点上 Pod，用于维护 |

### 调试

| 命令 | 作用 |
|---|---|
| `kubectl get events` | 查看集群事件 |
| `kubectl top pod` | 查看资源使用 |
| `kubectl port-forward svc/xxx 8080:80` | 本地转发访问服务 |
| `kubectl cp xxx:/tmp/a.log ./a.log` | 拷贝容器文件 |

---

## 五、Pod 生命周期与状态

```
Pending → ContainerCreating → Running → Succeeded / Failed
```

| 状态 | 含义 |
|---|---|
| **Pending** | 已提交，还没调度或镜像在拉取 |
| **ContainerCreating** | 正在创建容器 |
| **Running** | 运行中 |
| **CrashLoopBackOff** | 启动失败反复重启，常见是程序退出或健康检查不过 |
| **ImagePullBackOff** | 镜像拉取失败 |
| **OOMKilled** | 内存超过限制被 kill |
| **Evicted** | 节点资源不足被驱逐 |

**排查命令**：
```bash
kubectl describe pod <pod-name>     # 看 Events
kubectl logs <pod-name> --previous  # 看上一次崩溃日志
```

---

## 六、面试高频问题与答案

### Q1：K8s 和 Docker 的区别？

**答**：
- Docker 是**容器技术**，负责打包和运行单个容器
- K8s 是**容器编排工具**，负责管理很多容器（调度、扩缩容、自愈、服务发现）
- 关系：K8s 用 Docker/containerd 作为底层运行时

### Q2：Pod 和容器是什么关系？

**答**：
- 容器是 Docker 层面的运行单元
- Pod 是 K8s 的最小调度单元
- 一个 Pod 里可以跑多个容器，它们共享网络和存储

### Q3：Deployment 和 Pod 的区别？

**答**：
- Pod 是一次性运行的单元
- Deployment 管理一组 Pod，负责维持副本数、滚动更新、自动恢复
- 实际生产都用 Deployment，很少直接创建 Pod

### Q4：Service 是怎么实现负载均衡的？

**答**：
- Service 通过 spec.selector 根据标签匹配 Pod，生成后端端点列表 Endpoints；
- apiserver 维护 Endpoints 资源，实时同步所有正常运行 Pod 的 IP + 端口；
- 全节点 kube-proxy 持续监听 apiserver 的 Service、Endpoints 变更；
- kube-proxy 在本机写入转发规则，所有访问 Service ClusterIP 的流量被转发到真实 Pod。

### Q5：K8s 是怎么做滚动更新的？

**答**：
1. 创建新版本的 ReplicaSet
2. 先启动新 Pod
3. 等新 Pod Ready 后，逐步减少旧 Pod
4. 旧 ReplicaSet 保留，方便回滚

```bash
kubectl set image deployment/nginx-deploy nginx=nginx:1.26
kubectl rollout status deployment/nginx-deploy
kubectl rollout undo deployment/nginx-deploy
```

### Q6：Pod 一直 CrashLoopBackOff 怎么排查？

**答**：
1. `kubectl describe pod` 看 Events（如 OOM、镜像拉取失败）
2. `kubectl logs` 看应用错误
3. `kubectl logs --previous` 看崩溃前日志
4. 检查资源限制、健康检查配置、启动命令

### Q7：什么是 Helm？

**答**：
- Helm 是 K8s 的"包管理工具"
- 把一组 K8s 资源模板化打包成 Chart
- 类似 apt/yum，用于简化复杂应用部署

### Q8：K8s 如何实现服务发现？

**答**：
- 集群内通过 **Service 的 DNS 名**访问，如 `nginx-svc.default.svc.cluster.local`
- CoreDNS 负责把 Service 名解析为 ClusterIP
- Pod 间通过 Service + selector 自动发现并负载均衡

### Q9：Resource 的 requests 和 limits 有什么区别？

| 类型 | 作用 |
|---|---|
| **requests** | 调度时保证能给 Pod 的最小资源 |
| **limits** | Pod 实际能使用的最大资源，超过会被限流或 kill |

```yaml
resources:
  requests:
    cpu: "100m"
    memory: "128Mi"
  limits:
    cpu: "500m"
    memory: "512Mi"
```

### Q10：你们项目里怎么用的 K8s？

**答**（贴合调度系统项目）：
- 训练任务通过调度器生成 k8s Job/Deployment YAML
- 指定 GPU 资源（`nvidia.com/gpu: 1`），让任务调度到目标算力卡
- 训练容器挂载 PVC 读取数据集、写入模型
- 训练完成后通过 prometheus 采集该时间段指标，生成工况报告

### Q11：资源下发-集群部署-用户访问全流程
**答**：
- 阶段1： 用户提交资源（Deployment+Service），经由apiserver完成全集群配置同步
- 阶段2： 客户端发起请求，经由DNS、kube-proxy、CNI，最终到达业务Pod

**详细流程**
**步骤 1：客户端 kubectl 请求接入 apiserver**

运维执行 `kubectl apply -f nginx-all.yaml`：
1. kubectl 携带 token/证书访问 kube-apiserver
2. apiserver 完成认证、鉴权、准入控制校验
3. 校验通过，把 Deployment、Service 写入 etcd

> 重点：只有 apiserver 能读写 etcd，其他组件不直连。

**步骤 2：controller-manager 感知 etcd 变动，创建 Pod**

kube-controller-manager 持续监听 apiserver 资源变化：
1. Deployment 控制器监听到新增 Deployment
2. 控制器发现当前没有对应 Pod，向 apiserver 发起请求创建Pod对象并写入etcd
3. Pod 此时只存在etcd，但容器还没启动

**步骤 3：scheduler 调度 Pod**

调度器持续监听未绑定 node 的 Pod：
1. 预选：过滤资源不足、污点、亲和性不满足的节点
2. 优选：打分选出最优节点（如 node2）
3. scheduler 调用 apiserver，更新 Pod 的spec.nodeName绑定节点，数据落 etcd

**步骤 4：kubelet 拉起容器**

node2 上的 kubelet 一直监听 apiserver 本机 Pod 列表：
1. kubelet 发现有 Pod 调度到本机，拉取 Pod 完整 spec
2. 调用容器运行时（containerd）拉取镜像，创建并启动容器
3. kubelet 配置容器网络（CNI）、数据卷、就绪 / 存活探针；
4. 容器运行后，kubelet 持续上报 Pod 状态（running/error）给 apiserver，更新 etcd 里 Pod 的 status 字段

**步骤 5：Endpoints 生成，kube-proxy 同步规则**

1. Endpoints 控制器（controller-manager 内置）：根据 Service 的spec.selector匹配所有带app=nginx的就绪 Pod，自动生成 Endpoints 资源，保存所有正常 Pod 的 IP + 端口；
2. 集群所有节点 kube-proxy 监听 Service/Endpoints 变化
3. kube-proxy 在本机 Linux 内核（ipvs/iptables）写入转发规则，绑定 Service 的 ClusterIP，做好四层负载均衡配置

**步骤 6：CoreDNS 同步域名**

CoreDNS 监听 apiserver 内所有 Service，自动生成 DNS 记录：
nginx.default.svc.cluster.local 固定解析为 Service 的 ClusterIP。

至此集群配置全部就绪。

#### 阶段 2：用户访问流量全链路

**场景 A：集群内 Pod 访问 Service 域名**

客户端 `client-pod` 执行：
```bash
curl nginx.default.svc.cluster.local
```

1. **DNS 解析**：CoreDNS 返回 Service ClusterIP
2. **本机 ipvs 负载均衡**：内核根据 ipvs 规则选中一个后端 Pod IP
3. **CNI 路由**：
   - 同节点：Calico 本机网桥直达目标容器
   - 跨节点：Calico IPIP/VXLAN 封装，送到目标节点后解包
4. **业务处理**：nginx 容器处理请求，响应原路返回

> **联动自愈**：Pod 就绪探针失败 → Endpoints 剔除该 Pod → kube-proxy 删除转发规则 → 流量不再发到故障 Pod。

**场景 B：外网通过 NodePort 访问**

用户访问 `节点公网IP:30080`：
1. 流量进入任意宿主机 30080 端口
2. kube-proxy 将 NodePort 流量导向 Service ClusterIP
3. 后续流程同场景 A

**场景 C：外网通过 Ingress 访问（生产主流）**

用户访问 `nginx.test.com`：
1. 域名解析到 Ingress Controller 的公网 IP
2. 流量到达 ingress-nginx Pod
3. Ingress Controller 根据域名匹配后端 Service
4. 在集群内部访问 Service（重复场景 A 链路）
5. 业务响应经 Ingress 返回外网用户

> Ingress 层实现 SSL、路径路由、限流等七层能力。

**极简主干时间线**
kubectl → apiserver → etcd
→ controller-manager 创建 Pod → scheduler 选节点 → kubelet 启动容器
→ Endpoints 更新 → kube-proxy 配置内核负载规则 + CoreDNS 生成域名
→ 用户请求 DNS 解析 → 内核负载 → CNI 路由 → Pod 接收请求

---

## 七、项目实战：一个训练任务的 YAML

```yaml
apiVersion: batch/v1
kind: Job
metadata:
  name: train-job-001
spec:
  template:
    spec:
      restartPolicy: Never
      containers:
        - name: train
          image: train-image:v1
          command: ["python", "train.py"]
          resources:
            limits:
              nvidia.com/gpu: 1
          volumeMounts:
            - name: data
              mountPath: /data
            - name: output
              mountPath: /output
      volumes:
        - name: data
          persistentVolumeClaim:
            claimName: dataset-pvc
        - name: output
          persistentVolumeClaim:
            claimName: model-pvc
```

**提交任务**：
```bash
kubectl apply -f train-job.yaml
kubectl get job train-job-001
kubectl logs -f job/train-job-001
```

---

## 八、一句话总结

- **K8s 是容器编排平台**：自动部署、调度、扩缩、自愈
- **控制平面**：API Server + etcd + Scheduler + Controller Manager
- **工作节点**：kubelet + kube-proxy + Container Runtime
- **核心资源**：Pod（最小单元）、Deployment（管理 Pod）、Service（访问入口）
- **排查思路**：看状态 → `describe` 看事件 → `logs` 看日志 → `exec` 进容器

> **面试口诀：Pod 是员工，Deployment 是主管，Service 是前台，Scheduler 是 HR，etcd 是档案室**
