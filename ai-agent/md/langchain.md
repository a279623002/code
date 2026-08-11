LangChain 1.x（新版）底层架构、整体工作流、工具调用完整原理

##### 1. 底层分层架构

```
业务代码
——————————————————————————————
LangChain Application Layer
  ｜ create_agent / create_openai_tools_agent / RetrievalAgent
  ｜ （上层封装，底层依赖 LangGraph）
——————————————————————————————
LangGraph（状态机运行时，Agent循环引擎）
  ｜ StateGraph、Node、Edge、条件路由、循环、断点持久化
——————————————————————————————
langchain-core【核心底层，LCEL 运行时】
  ｜ Runnable 协议、BaseMessage、BaseTool、回调系统、序列化
——————————————————————————————
Model Integrations（langchain-openai / ollama 等）
  ｜ ChatModel、bind_tools、API 适配层
——————————————————————————————
LLM API（厂商原生 Function Calling / Tool Calling）
```

###### 1. 最底层：Runnable 协议（LCEL 的根基）

所有组件（Prompt、ChatModel、Tool、Parser、自定义逻辑）都实现 Runnable 抽象接口：

```python
class Runnable[Input, Output]:
    def invoke(self, input, config) -> Output      # 同步单次调用
    def ainvoke(self, input, config) -> Awaitable[Output] # 异步
    def stream(self, input, config) -> Iterator   # 流式输出
    def batch(self, inputs, config) -> List[Output]
    def pipe(self, other: Runnable) -> RunnableSequence # | 运算符
```

常用内置 Runnable：
- RunnableSequence：a | b | c 串行管道
- RunnableParallel：并行执行多个分支
- RunnableLambda：普通函数包装成 Runnable
- RunnablePassthrough：透传数据，用于字段重组

###### 2. LCEL:（| 管道语法）、流式、批处理、异步全部建立在这套协议之上。

**本质**：声明式构建一张执行 DAG(有向无环图)，invoke 触发自上而下顺序执行；LCEL 本身没有循环能力，因此多轮 Agent 工具调用必须使用 LangGraph。

###### 3. 消息基础模型（整个工具调用的数据载体）

**langchain_core.messages** 消息体系，所有上下文、工具交互全部依靠消息列表传递：

- HumanMessage：用户提问
- AIMessage：模型输出；携带 tool_calls 代表想要调用工具
```python
tool_calls = [{
    "id": "call-uuid",      # 调用唯一ID，必须配对
    "name": "get_weather",
    "args": {"city": "广州"}
}]
```
- ToolMessage：工具执行结果；必须携带对应 tool_call_id，和 AIMessage 一一配对
**规范**：消息数组 messages 是 Agent 全局上下文，完整历史持续传递给 LLM。

##### 2. 标准 Agent 多轮工具调用

**新版 create_agent**：底层 LangGraph 状态机,不再是 while 循环硬编码，而是自动构建一张二元状态图 StateGraph

**Graph 结构**（两个核心节点）
```
START → model_node
          ↓（条件判断）
    有tool_calls → tools_node → model_node（循环）
    无tool_calls → END
```
###### 1. model_node（LLM 决策节点）
读取 state 中 messages，调用绑定 tools 的 LLM，产出 AIMessage，追加进消息列表。
###### 2. tools_node（工具执行节点）
遍历 ai_msg.tool_calls：
- 支持并行批量执行多个工具调用
- 逐个执行工具，生成多条 ToolMessage
- 全部追加到 messages

