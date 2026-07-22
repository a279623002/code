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
│                      控制平面（Control Plane）                 │
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
| **etcd** | 分布式键值存储，保存集群状态 | "K8s 的数据库" |
| **Scheduler** | 负责把 Pod 调度到合适的 Node | "HR，决定新员工去哪个部门" |
| **Controller Manager** | 维持集群期望状态 | "监工，发现实际状态和配置不一样就修复" |

### 工作节点组件（Node）

| 组件 | 作用 |
|---|---|
| **kubelet** | 每个 Node 上的"小管家"，负责创建/销毁 Pod，汇报状态 |
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
- Service 通过 **selector** 找到后端 Pod
- kube-proxy 维护 iptables / IPVS 规则，把请求转发到 Pod
- 每个 Service 分配一个 ClusterIP，集群内稳定访问

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
