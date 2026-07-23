# 计算机网络协议面试笔记

> 项目背景：边端智能助手用 HTTP/WebSocket/SSE 与前端交互，调度系统用 gRPC/HTTP 做微服务通信。网络协议是前后端、微服务、AI 集成的共同基础。

---

## 一、OSI 七层模型

```
┌─────────────────────────────────────┐
│  第7层 应用层  │ HTTP、DNS、FTP、SMTP   │
│  第6层 表示层  │ SSL/TLS 加密、编码转换  │
│  第5层 会话层  │ 会话管理、连接维护      │
│  第4层 传输层  │ TCP、UDP              │
│  第3层 网络层  │ IP、ICMP、路由         │
│  第2层 数据链路层│ MAC、交换机           │
│  第1层 物理层  │ 网线、光纤、电信号      │
└─────────────────────────────────────┘
```

**记忆口诀**：
> **应表会传网数物**（从上到下）

**常考四层模型（TCP/IP）**：

| 层 | 协议 |
|---|---|
| 应用层 | HTTP、DNS、FTP |
| 传输层 | TCP、UDP |
| 网络层 | IP、ICMP |
| 网络接口层 | MAC、以太网 |

---

## 二、TCP vs UDP

### 1. 对比

| 特性 | TCP | UDP |
|---|---|---|
| 连接 | 面向连接（三次握手） | 无连接 |
| 可靠性 | 可靠传输，保证顺序 | 不保证可靠，不保证顺序 |
| 流量控制 | 有滑动窗口 | 无 |
| 拥塞控制 | 有 | 无 |
| 头部开销 | 大（20字节） | 小（8字节） |
| 速度 | 慢 | 快 |
| 场景 | HTTP、文件传输、数据库 | DNS、视频直播、游戏、IoT |

### 2. TCP 三次握手

```
客户端 ────── SYN=1, seq=x ─────→ 服务端
客户端 ←── SYN=1, ACK=1, seq=y, ack=x+1 ─── 服务端
客户端 ───── ACK=1, seq=x+1, ack=y+1 ───→ 服务端

状态变化：
CLOSED → SYN_SENT → ESTABLISHED
CLOSED → LISTEN → SYN_RCVD → ESTABLISHED
```

**为什么是三次？**
- 第一次：客户端证明能发
- 第二次：服务端证明能收、能发
- 第三次：客户端证明能收

### 3. TCP 四次挥手

```
客户端 ───── FIN=1, seq=u ─────→ 服务端
客户端 ←──── ACK=1, ack=u+1 ───── 服务端
客户端 ←──── FIN=1, seq=v ───── 服务端
客户端 ───── ACK=1, ack=v+1 ───→ 服务端

状态变化：
ESTABLISHED → FIN_WAIT_1 → FIN_WAIT_2 → TIME_WAIT → CLOSED
ESTABLISHED → CLOSE_WAIT → LAST_ACK → CLOSED
```

**为什么 TIME_WAIT 等 2MSL？**
- 确保最后一个 ACK 对方收到
- 等网络中残留报文消失，防止新连接收到旧数据

### 4. TCP 如何保证可靠传输？

1. **序列号/确认应答**：每个包有编号，收到后回 ACK
2. **超时重传**：没收到 ACK 就重发
3. **滑动窗口**：批量发送，提高利用率
4. **流量控制**：接收方告诉发送方自己还能收多少
5. **拥塞控制**：慢启动、拥塞避免、快重传、快恢复

---

## 三、HTTP

### 1. HTTP 报文结构

```
请求：
GET /index.html HTTP/1.1
Host: www.example.com
User-Agent: Mozilla/5.0
Accept: text/html

响应：
HTTP/1.1 200 OK
Content-Type: text/html
Content-Length: 1234

<html>...</html>
```

### 2. 常见状态码

| 状态码 | 含义 |
|---|---|
| 200 | 成功 |
| 301/302 | 重定向 |
| 400 | 请求参数错误 |
| 401 | 未认证 |
| 403 | 无权限 |
| 404 | 资源不存在 |
| 500 | 服务器内部错误 |
| 502 | 网关错误 |
| 503 | 服务不可用 |

### 3. HTTP/1.1 vs HTTP/2 vs HTTP/3

| 特性 | HTTP/1.1 | HTTP/2 | HTTP/3 |
|---|---|---|---|
| 连接 | 串行 / 管线化 | 多路复用 | 基于 QUIC |
| 头部 | 重复传输 | HPACK 压缩 | QPACK 压缩 |
| 服务器推送 | 不支持 | 支持 | 支持 |
| 传输层 | TCP | TCP | UDP（QUIC） |

### 4. HTTP 方法

| 方法 | 作用 | 幂等性 |
|---|---|---|
| GET | 获取资源 | 幂等 |
| POST | 创建资源 | 不幂等 |
| PUT | 更新资源（全量） | 幂等 |
| PATCH | 更新资源（局部） | 不一定幂等 |
| DELETE | 删除资源 | 幂等 |

---

## 四、HTTPS 与 SSL/TLS

### 1. HTTPS 是什么？

HTTP + SSL/TLS = HTTPS，在传输层对数据进行加密。

### 2. TLS 握手过程

```
客户端 ───── 支持的 TLS 版本、加密套件、随机数 ───→ 服务端
客户端 ←──── 证书、选择的加密套件、随机数 ───────── 服务端
客户端 ───── 用公钥加密预主密钥 ─────────────────→ 服务端
双方用随机数 + 预主密钥生成会话密钥，后续用对称加密通信
```

### 3. 对称加密 vs 非对称加密

| 类型 | 特点 | 用途 |
|---|---|---|
| 对称加密 | 快，双方用同一个密钥 | 传输数据 |
| 非对称加密 | 慢，公钥加密私钥解密 | 握手时交换密钥 |

### 4. 证书验证

1. 客户端拿到服务端证书
2. 用 CA 公钥验证证书签名
3. 检查域名、有效期、是否吊销
4. 确认无误才继续握手

---

## 五、WebSocket

### 1. 概念

**一句话**：基于 TCP 的**全双工**通信协议，客户端和服务端可以互相主动发消息。

```
HTTP 升级过程：
客户端 ───── GET /chat HTTP/1.1 ─────→ 服务端
            Upgrade: websocket
            Connection: Upgrade
客户端 ←──── HTTP/1.1 101 Switching Protocols ─── 服务端
            之后变成全双工 TCP 长连接
```

### 2. 与 HTTP 对比

| 特性 | HTTP | WebSocket |
|---|---|---|
| 方向 | 客户端请求，服务端响应 | 双向 |
| 连接 | 短连接 | 长连接 |
| 实时性 | 差 | 好 |
| 头部开销 | 大 | 小 |

### 3. 应用场景

- 在线聊天
- 股票行情实时推送
- 多人协作编辑
- 游戏实时同步

### 4. 项目应用

```python
# FastAPI WebSocket 示例
from fastapi import FastAPI, WebSocket

app = FastAPI()

@app.websocket("/ws")
async def websocket_endpoint(ws: WebSocket):
    await ws.accept()
    while True:
        data = await ws.receive_text()
        await ws.send_text(f"收到: {data}")
```

---

## 六、SSE（Server-Sent Events）

### 1. 概念

**一句话**：服务端单向推送，基于 HTTP，适合流式输出。

### 2. 与 WebSocket 对比

| 特性 | SSE | WebSocket |
|---|---|---|
| 方向 | 服务端 → 客户端单向 | 双向 |
| 协议 | 基于 HTTP | 独立协议 |
| 重连 | 浏览器自动重连 | 需手动实现 |
| 场景 | 日志流、股票推送、AI 流式回答 | 聊天、游戏 |

### 3. 项目应用

```python
# FastAPI SSE 流式输出 LLM
from fastapi import FastAPI
from fastapi.responses import StreamingResponse

app = FastAPI()

async def stream():
    for token in ["你", "好", "，", "世", "界"]:
        yield f"data: {token}\n\n"

@app.get("/sse")
def sse():
    return StreamingResponse(stream(), media_type="text/event-stream")
```

---

## 七、DNS

### 1. 作用

把域名解析成 IP 地址。

### 2. 解析过程

```
用户输入 www.example.com
    ↓
浏览器缓存 → 操作系统缓存 → 本地 DNS 服务器
    ↓
根 DNS 服务器 → 顶级域 DNS（.com）→ 权威 DNS（example.com）
    ↓
返回 IP 地址
```

### 3. 记录类型

| 类型 | 作用 |
|---|---|
| A | 域名 → IPv4 |
| AAAA | 域名 → IPv6 |
| CNAME | 域名 → 另一个域名 |
| MX | 邮件服务器 |
| TXT | 文本记录，常用于验证 |

---

## 八、SCP

**一句话**：基于 SSH 的安全文件拷贝命令。

```bash
# 本地 → 远程
scp file.txt user@remote:/path/

# 远程 → 本地
scp user@remote:/path/file.txt ./

# 目录
scp -r dir user@remote:/path/
```

与 FTP 区别：SCP 全程加密，更安全。

---

## 九、A2A（Agent-to-Agent）

### 1. 概念

**一句话**：智能体之间的通信协议，让不同厂商/框架的 AI Agent 能互相发现、协作。

### 2. 为什么需要 A2A？

传统协议（HTTP/gRPC）只能传数据，A2A 还要传递：
- Agent 能力描述（我能做什么）
- 任务上下文（多轮状态）
- 安全凭证
- 工具调用约定

### 3. A2A 核心能力

| 能力 | 说明 |
|---|---|
| Agent Card | 描述 Agent 的能力、输入输出、认证方式 |
| Task | 一个可执行的任务单元 |
| Artifact | 任务产出结果 |
| Message | Agent 之间交换的消息 |

### 4. 应用场景

- 一个 Agent 调用另一个 Agent 完成子任务
- 多 Agent 协作完成复杂工作流
- 跨平台 Agent 生态互通

---

## 十、面试高频问题

### Q1：TCP 和 UDP 的区别？

**答**：
- TCP 面向连接、可靠、有流量和拥塞控制，适合 HTTP、文件传输
- UDP 无连接、快、不保证可靠，适合视频、游戏、DNS

### Q2：三次握手为什么不能两次？

**答**：两次握手只能证明客户端能发、服务端能收能发，但无法证明客户端能收。如果客户端没收到 SYN+ACK，服务端就认为连接已建立，会浪费资源。

### Q3：HTTP 和 HTTPS 的区别？

**答**：
- HTTPS = HTTP + SSL/TLS
- HTTPS 通过证书验证服务端身份
- 传输数据加密，防止中间人攻击
- 默认端口 443（HTTP 是 80）

### Q4：WebSocket 和 SSE 的区别？

**答**：
- WebSocket 双向、基于 TCP 新协议
- SSE 单向（服务端 → 客户端）、基于 HTTP
- SSE 更简单，浏览器自动重连；WebSocket 更灵活，适合双向实时场景

### Q5：输入 URL 到页面显示发生了什么？

**答**：
1. DNS 解析域名得到 IP
2. TCP 三次握手建立连接
3. TLS 握手（HTTPS）
4. 发送 HTTP 请求
5. 服务端处理并返回响应
6. 浏览器解析 HTML、CSS、JS，渲染页面

### Q6：什么是 QUIC？

**答**：QUIC 是基于 UDP 的传输协议，HTTP/3 用它。特点是：
- 握手更快（0-RTT / 1-RTT）
- 连接迁移（IP 变了连接还在）
- 解决 TCP 队头阻塞

---

## 十一、一句话总结

- **OSI**：应表会传网数物，七层模型记清楚
- **TCP**：可靠传输靠三次握手 + 四次挥手 + 确认重传
- **UDP**：快但不保证可靠，适合实时场景
- **HTTP/HTTPS**：无状态请求响应，HTTPS 加了 TLS 加密
- **WebSocket**：全双工长连接，双向实时
- **SSE**：基于 HTTP 的服务端单向推送
- **DNS**：域名 → IP 的翻译官
- **A2A**：Agent 之间的协作协议
