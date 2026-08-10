# MCP（Model Context Protocol）协议笔记

> 一句话：MCP 是 Anthropic 提出的开放协议，让 LLM 应用通过标准 JSON-RPC 消息调用外部工具、读取资源、复用提示词。

---

## 一、为什么需要 MCP？

### 传统对接方式的问题

每个工具都要单独写适配代码：

```
LLM 应用
   │
   ├── 接 Slack API
   ├── 接 GitHub API
   ├── 接数据库
   └── 每个都要写鉴权、协议转换
```

问题：
- 接入成本高
- 协议不统一
- 工具发现困难
- 安全策略难以统一

### MCP 统一后

```
┌──────────┐     ┌─────────────┐     ┌─────────────┐     ┌─────────────┐
│   Host   │────→│ MCP Client  │────→│ MCP Server  │────→│ 工具/资源/提示 │
│ (LLM应用) │     │ (维护连接)   │     │ (标准协议)   │     │ (实际能力)    │
└──────────┘     └─────────────┘     └─────────────┘     └─────────────┘
```

**核心价值**：
- 一次对接，复用所有 MCP Server
- 统一鉴权、调用、返回格式
- 运行时动态发现能力

---

## 二、MCP 核心角色

| 角色 | 说明 | 例子 |
|---|---|---|
| **Host** | 运行 LLM 的应用 | Claude Desktop、Cursor、自建 Agent |
| **Client** | Host 内的客户端，维护与 Server 的长连接 | 每个 Server 一个 Client |
| **Server** | 对外暴露 Tools/Resources/Prompts 的服务端 | 文件系统 Server、数据库 Server |

**关系**：
- 一个 Host 可以连接多个 MCP Server
- 每个 Server 对应一个 Client，Client 负责维护连接和标准协议通信
- 一个 Server 只负责暴露自己的能力，不内置路由逻辑
- Server 不直接调用其他 Server，由 Host 统筹调度

**三者关系图**：

```
          ┌──────────────────────────────────────────┐
          │                  Host                     │
          │  ┌─────────┐    ┌─────────┐    ┌────────┐ │
          │  │ Client  │    │ Client  │    │ Client │ │
          │  │ 连接 A  │    │ 连接 B  │    │ 连接 C │ │
          │  └────┬────┘    └────┬────┘    └───┬────┘ │
          └───────┼──────────────┼─────────────┼──────┘
                  │              │             │
                  ▼              ▼             ▼
            ┌─────────┐    ┌─────────┐   ┌─────────┐
            │ Server A│    │ Server B│   │ Server C│
            │ 文件系统 │    │ 数据库  │   │  搜索   │
            └─────────┘    └─────────┘   └─────────┘
```

**代码示例：一个简单 MCP Host 连接两个 Server**

```python
# server_a.py：文件系统 Server（stdio 模式）
from mcp.server import Server
from mcp.types import Tool, TextContent

app = Server("file-server")

@app.list_tools()
async def list_tools():
    return [Tool(name="read_file", description="读取文件", inputSchema={"type": "object", "properties": {"path": {"type": "string"}}})]

@app.call_tool()
async def call_tool(name: str, arguments: dict):
    if name == "read_file":
        with open(arguments["path"], "r", encoding="utf-8") as f:
            return [TextContent(type="text", text=f.read())]
    raise ValueError(f"未知工具: {name}")

# 通过 stdio 启动：python server_a.py
```

```python
# server_b.py：数据库 Server（stdio 模式）
from mcp.server import Server
from mcp.types import Tool, TextContent

app = Server("db-server")

@app.list_tools()
async def list_tools():
    return [Tool(name="query_sql", description="执行 SQL", inputSchema={"type": "object", "properties": {"sql": {"type": "string"}}})]

@app.call_tool()
async def call_tool(name: str, arguments: dict):
    if name == "query_sql":
        # 实际调用数据库
        result = db.execute(arguments["sql"])
        return [TextContent(type="text", text=str(result))]
    raise ValueError(f"未知工具: {name}")
```

```python
# host.py：Host 统筹多个 Client
import asyncio
from mcp import ClientSession, StdioServerParameters
from mcp.client.stdio import stdio_client

async def run():
    # 配置两个本地 Server
    servers = [
        StdioServerParameters(command="python", args=["server_a.py"]),
        StdioServerParameters(command="python", args=["server_b.py"]),
    ]

    all_tools = {}
    for params in servers:
        async with stdio_client(params) as (read, write):
            async with ClientSession(read, write) as session:
                await session.initialize()
                tools = await session.list_tools()
                all_tools[params.args[0]] = tools

    # LLM 根据用户问题决定调用哪个 Server 的哪个 Tool
    # 例如：用户说"查一下订单表"，LLM 选择 db-server.query_sql
    print("可用工具:", all_tools)

asyncio.run(run())
```

**代码说明**：
- Host 启动时分别为每个 Server 创建一个 Client
- Client 通过 `session.initialize()` 完成握手
- `list_tools()` 动态发现 Server 提供的工具
- LLM 根据用户意图选择 Tool，Client 通过 `call_tool()` 发起调用
- Server 只负责执行自己的工具，不感知其他 Server

---

## 三、MCP 三种能力

| 能力 | 说明 | 用途 |
|---|---|---|
| **Tools** | 可调用的工具/函数 | 查天气、发邮件、执行 SQL |
| **Resources** | 只读数据，供上下文使用 | 文件内容、文档、配置信息 |
| **Prompts** | 预定义提示词模板 | 代码审查、报告生成、翻译模板 |

**原则**：
- Tools 是"动作"，会改变外部状态
- Resources 是"数据"，只读
- Prompts 是"模板"，帮助 LLM 生成更好的请求

---

## 四、核心协议方法

MCP 基于 **JSON-RPC 2.0**，Server 必须响应以下标准方法：

### 1. 生命周期方法

| 方法 | 作用 |
|---|---|
| `initialize` | 握手，协商协议版本、能力 |
| `initialized/notification` | 客户端通知初始化完成 |
| `notifications/cancelled` | 取消正在进行的请求 |

### 2. 能力发现方法

| 方法 | 作用 |
|---|---|
| `tools/list` | 返回 Server 提供的工具列表 |
| `resources/list` | 返回资源列表 |
| `prompts/list` | 返回提示词模板列表 |

### 3. 能力调用方法

| 方法 | 作用 |
|---|---|
| `tools/call` | 调用指定工具 |
| `resources/read` | 读取指定资源 |
| `prompts/get` | 获取提示词模板内容 |

### 4. 通知方法

| 方法 | 作用 |
|---|---|
| `notifications/tools/list_changed` | 工具列表变化时通知 Client |
| `notifications/resources/list_changed` | 资源列表变化时通知 Client |
| `notifications/prompts/list_changed` | 提示词列表变化时通知 Client |

---

## 五、典型交互流程

```
Host                    MCP Client              MCP Server
 │                          │                       │
 │  启动                    │                       │
 │─────────────────────────→│                       │
 │                          │  1. initialize        │
 │                          │──────────────────────→│
 │                          │  返回协议版本、能力     │
 │                          │←──────────────────────│
 │                          │  2. tools/list        │
 │                          │──────────────────────→│
 │                          │  返回可用工具列表       │
 │                          │←──────────────────────│
 │                          │                       │
 │  用户提问：北京天气如何    │                       │
 │─────────────────────────→│                       │
 │                          │  3. tools/call        │
 │                          │   (weather, city=北京) │
 │                          │──────────────────────→│
 │                          │  返回天气结果           │
 │                          │←──────────────────────│
 │                          │                       │
 │  生成最终回答             │                       │
 │←─────────────────────────│                       │
```

---

## 六、通信方式

| 方式 | 说明 | 场景 |
|---|---|---|
| **stdio** | 标准输入输出，本地进程间通信 | 本地 Server，如文件系统、SQLite |
| **SSE / HTTP** | Server-Sent Events + HTTP POST | 远程 Server，可网络访问 |

**stdio 特点**：
- 启动一个本地子进程
- Host 通过 stdin 发请求，stdout 收响应
- 简单、安全、适合本地工具

**SSE/HTTP 特点**：
- Server 独立部署
- 支持远程访问和多客户端连接
- 需要处理鉴权、网络稳定性

---

## 七、Server 的职责边界

**MCP Server 只负责**：
- 暴露 Tools/Resources/Prompts
- 实现标准 JSON-RPC 方法
- 执行被调用的具体工具

**MCP Server 不负责**：
- 路由到其他 Server
- 编排多个工具调用
- LLM 推理决策
- 全局会话管理

这些由 **Host** 或上层 Agent 框架负责。

---

## 八、与 Function Call 的区别

| 对比 | Function Call | MCP |
|---|---|---|
| 定位 | 模型调用函数的能力 | 模型与外部系统连接的标准协议 |
| 范围 | 单次调用单个函数 | 可暴露多个工具、资源、提示词 |
| 发现 | 调用方预先知道函数定义 | 运行时通过 `tools/list` 动态发现 |
| 生态 | 各模型接口略有差异 | 跨模型、跨平台统一标准 |
| 部署 | 通常内嵌在应用中 | 独立 Server，可本地或远程 |

---

## 九、面试高频问题

### Q1：MCP 是什么？解决了什么问题？

**答**：MCP 是 Anthropic 提出的开放协议，让 LLM 应用通过标准 JSON-RPC 调用外部工具、读取资源、复用提示词。解决了每个工具都要单独对接、协议不统一、能力发现难的问题。

### Q2：MCP 有哪些核心角色？

**答**：Host（运行 LLM 的应用）、Client（维护与 Server 的连接）、Server（暴露 Tools/Resources/Prompts）。Server 只负责暴露能力，不内置路由逻辑。

### Q3：MCP Server 必须实现哪些协议方法？

**答**：生命周期方法如 `initialize`；发现方法如 `tools/list`、`resources/list`、`prompts/list`；调用方法如 `tools/call`、`resources/read`、`prompts/get`。

### Q4：MCP 和 Function Call 有什么区别？

**答**：Function Call 是模型调用函数的能力，MCP 是连接模型与外部系统的标准协议。MCP 支持运行时动态发现多个工具，更适合构建开放生态。

### Q5：MCP Server 内置路由逻辑吗？

**答**：不内置。MCP Server 只暴露自己的能力，由 Host 根据用户请求决定调用哪个 Server 的哪个 Tool。路由和编排属于 Host 或上层 Agent 的职责。

### Q6：通信方式有哪些？

**答**：stdio 适合本地进程间通信，SSE/HTTP 适合远程 Server。

---

## 十、项目口述模板

> 请介绍你项目中的 MCP。

我在项目中按照 MCP 标准协议实现了多个 MCP Server，用于把数据库查询、文件读取、外部 API 调用等能力统一暴露给 LLM 应用。每个 Server 通过 `initialize` 完成握手，Host 通过 `tools/list` 动态发现可调用的工具。当用户提问时，LLM 决定调用哪个 Tool，Client 通过 `tools/call` 发起标准 JSON-RPC 请求，Server 执行具体逻辑并返回结构化结果。MCP Server 只负责暴露和执行自身能力，不涉及路由和编排，整体接入成本低、协议统一、能力可复用。

---

## 十一、一句话总结

- **MCP = LLM 与外部世界的 USB-C 接口**
- **Server 只暴露 Tools/Resources/Prompts**
- **标准方法：initialize、tools/list、tools/call、resources/read、prompts/get**
- **无内置路由，Host 负责调度和编排**
