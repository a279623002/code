# AI应用开发面试八股文

## 大模型基础

### 1. 大语言模型(LLM)的基本原理
- 基于Transformer架构
- 自回归预测下一个token
- 预训练 + 微调 + RLHF
- 涌现能力：Scaling Law

### 2. Transformer架构的核心组件
- Encoder-Decoder结构
- Self-Attention自注意力机制
- Multi-Head Attention多头注意力
- Feed-Forward Network前馈网络
- Layer Normalization和残差连接
- Positional Encoding位置编码

### 3. Attention机制的计算
- Q(Query)、K(Key)、V(Value)
- Attention(Q,K,V) = softmax(QK^T/√dk)V
- Masked Attention：Decoder中遮挡未来token
- Cross Attention：Encoder-Decoder之间

### 4. 什么是Token？
- 模型处理的基本单位
- 可以是字、词、子词
- BPE（Byte-Pair Encoding）算法
- 不同模型有不同的Tokenizer

### 5. 大模型的局限性
- 知识截止日期
- 幻觉问题
- 上下文长度限制
- 推理成本高
- 多模态能力有限

## RAG技术

### 6. RAG（检索增强生成）的概念
- Retrieve：从知识库检索相关文档
- Augment：将检索结果作为上下文
- Generate：基于上下文生成回答
- 解决知识更新和幻觉问题

### 7. RAG的完整流程
1. 文档预处理
2. 向量化（Embedding）
3. 存储到向量数据库
4. 用户查询向量化
5. 相似度检索Top-K
6. 构建Prompt
7. LLM生成回答

### 8. 为什么选择FAISS作为向量数据库？
- Facebook开源，性能优秀
- 本地部署，无需额外服务
- 延迟低，成本低
- 适合中小型知识库场景
- 支持多种索引类型（FlatL2、IVFFlat等）

### 9. Embedding模型的选择
- Sentence Transformers
- bge系列（bge-large-zh-v1.5）
- OpenAI text-embedding-ada-002
- 中文任务优先选择中文Embedding模型

### 10. 向量相似度计算方法
- 余弦相似度：最常用，-1到1
- L2距离（欧氏距离）
- 内积
- FAISS默认L2，可转换为余弦相似度

### 11. 如何优化RAG的检索效果？
- 调整chunk_size和chunk_overlap
- 在chunk前添加文件名、章节等元信息增强语义
- 调整match_threshold过滤低质量结果
- 选择合适的Embedding模型
- 重排序（Rerank）
- 混合检索（BM25 + 向量）

### 12. 文档分块策略
- 固定大小分块：简单但可能破坏语义
- 递归字符分块：按标点符号智能分块
- 语义分块：按段落、章节、逻辑单元划分
- 重叠分块：保留上下文连贯性

### 13. RAG的挑战和解决方案
- 检索不相关：优化检索策略、混合检索
- 上下文窗口限制：更好的分块、摘要
- 时效性：定期更新向量库
- 多跳问题：迭代检索
- 多语言：使用多语言Embedding

## LangChain

### 14. LangChain的核心概念
- Chain：多个组件串联执行
- LLM：大语言模型抽象
- PromptTemplate：提示词模板
- Memory：对话记忆
- Agent：智能体
- Tools：工具调用
- VectorDB：向量数据库集成

### 15. Chain的类型
- LLMChain：最基本的Chain
- SequentialChain：顺序执行多个Chain
- RouterChain：根据输入路由到不同Chain
- TransformChain：数据转换
- RetrievalQAChain：检索问答

### 16. Memory的类型
- ConversationBufferMemory：保存所有对话历史
- ConversationSummaryMemory：总结对话历史
- ConversationBufferWindowMemory：保存最近K条
- ConversationSummaryBufferMemory：混合方式
- ConversationKGMemory：知识图谱记忆

### 17. Agent的工作原理
- ReAct模式：Reasoning + Acting
- 规划：理解任务，规划步骤
- 工具调用：使用各种工具获取信息
- 反思：评估结果，修正策略
- 执行：完成任务

### 18. 如何创建自定义Agent？
- 定义Tools
- 创建PromptTemplate
- 初始化Agent
- 执行AgentExecutor

### 19. LangChain的优势和劣势
- 优势：组件丰富、生态完善、快速上手
- 劣势：抽象层多、调试困难、性能开销

## 提示词工程(Prompt Engineering)

### 20. 提示词设计原则
- 指令清晰明确
- 提供示例（Few-Shot Learning）
- 按角色设定（System Prompt）
- 格式规范
- 引导思维过程（Chain of Thought）

### 21. Few-Shot Learning
- 在提示中提供几个示例
- 帮助模型理解任务格式
- 示例要多样化
- 示例数量3-10个为宜

### 22. Chain of Thought (CoT)
- 让模型逐步思考
- "让我们一步步思考"
- 提升复杂推理任务效果
- Zero-Shot CoT

### 23. 提示词常见技巧
- 结构化输出（JSON、XML）
- 角色设定
- 约束条件
- 风格要求
- 温度参数设置

### 24. 提示词注入攻击
- 用户输入恶意提示词
- 绕过安全限制
- 防护措施：输入过滤、输出验证、Sandbox

## Agent与智能体

### 25. Agent系统的核心模块
- 任务规划：将复杂任务拆解为步骤
- 记忆管理：短期记忆、长期记忆
- 工具调用：搜索、API、数据库等
- 反思与自修正：评估结果，优化策略
- 决策机制：选择下一步行动

### 26. 主流Agent框架
- LangChain Agent
- AutoGPT
- BabyAGI
- Openclaw
- Hermes Agent
- Claude Code

### 27. 多Agent协作
- 多个Agent分工协作
- 每个Agent有专长
- Agent之间通信
- 协调机制

### 28. Agent的观测体系
- 全链路追踪
- 决策日志
- 工具调用记录
- 性能指标
- 效果评估

## 项目经验相关

### 29. 如何处理空向量库？
- 场景：新建知识库，没有文档时FAISS需要至少一个向量
- 解决方案：
  - 实现add_zero_faiss()方法
  - 通过嵌入空字符串获取维度
  - 手动构建IndexFlatL2 + InMemoryDocstore + 空映射字典
  - 保存空库到磁盘，后续可直接加载

### 30. 大文件上传内存溢出问题
- 场景：上传100MB+超大文档
- 解决方案：
  - 分批处理文档，batch_size=50
  - 使用流式写入临时文件
  - 实时记录处理进度

### 31. FAISS相似度分数计算问题
- 场景：FAISS返回L2距离，用户习惯余弦相似度0-1范围
- 解决方案：
  - 公式：score_threshold = sqrt(2 * (1 - base_score_threshold))
  - 支持用户输入0-1阈值，内部自动转换为L2距离

### 32. Word文档解析效果差
- 场景：UnstructuredWordLoader对中文文档解析不理想
- 解决方案：
  - 自定义CustomUnstructuredWordLoader继承UnstructuredLoader
  - 使用Mammoth库进行底层解析
  - 提升中文文档处理质量

### 33. 模型加载速度优化
- 场景：每次请求都重新加载Embedding模型，响应延迟高
- 解决方案：
  - 全局单例模式缓存模型
  - @app.on_event("startup")服务启动时预加载
  - 懒加载可选配置

### 34. 向量库性能优化
- 模型预加载机制：全局缓存，模型只加载一次
- 向量库缓存：避免频繁磁盘IO
- 批处理插入：batch_faiss_insert()批量50条
- 冷启动优化：服务启动时预加载默认知识库，首次响应延迟降低80%+

### 35. 工程化设计经验
- 分层架构：API层→Service层→Model层
- 配置管理：基于Pydantic的配置模型，支持YAML热更新
- 统一响应格式：success_response/error_response/warn_response
- 模块化日志：不同模块独立日志文件
- API密钥认证：Bearer Token鉴权
- CORS跨域支持
- 健康检查接口/health

### 36. Loader工厂模式
- 字典映射文件扩展名到对应Loader
- 易于添加新格式支持
- 配置化参数，无需改代码
- 多设备支持：CPU/GPU两种运行模式

### 37. Whisper会议纪要转写
- 语音转文字
- 长音频分片处理
- 说话人识别（可选）
- 摘要生成

### 38. OpenAI兼容API设计
- 标准Chat Completions接口
- 支持流式/非流式响应
- 可直接对接OpenAI SDK
- 降低迁移成本

### 39. 系统监控和运维
- Loguru日志记录关键操作
- /health健康检查接口
- 配置文件支持动态更新
- 完善的异常处理和日志记录

### 40. 如何处理多轮对话？
- 当前版本专注于知识库检索
- 可通过在Prompt中添加上下文历史实现
- 后续可集成LangChain的ConversationBufferMemory
