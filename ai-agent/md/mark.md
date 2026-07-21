### 机器人管控平台
**项目介绍**：基于goframe开发，录入机器人基础信息，地图模块，任务与模型；我这边负责处理mqtt话题与模型管理（包括订阅话题和控制话题，控制话题可以发送机器人行走、表情、动作的话题；订阅话题则是订阅控制话题发送后的结果，以及机器人上报的电量、状态、心跳、动作、表情列表等存入到redis同时开放查询接口）；模型管理是可自定义参数配置的，根据不同的模型适配调用方式；
**技术栈**：
1. mqtt
2. 并发处理
3. redis

### 调度系统
**项目介绍**：基于gin开发，录入算力中心，配置算力集群，机器与算力卡通过prometheus采集，模型与数据集通过minio存储，任务通过算法策略进行编排，通过调度器调用k8s让任务在目标算力卡上训练，可支持手动训练，策略编排（按电价、空闲算力）训练，训练时容器上报任务进度，训练成功后工具目标算力卡以及训练时间段查询prometheus指标，生成工况excel表存入minio；我负责的是算力资源管理、异步任务编排生成策略，以及prometheus这部分
**技术栈**：
1. 异步编程
2. k8s


### 边端智能助手
**项目介绍**：基于langchain==0.2.16+faiss开发，支持多知识库切换，会议纪要、语音转写等；我负责多知识库管理，文档通过markitdown转成md格式进行切片再转成向量嵌入入库，然后通过语音，文字进行问答。
**问题**：
1.rag
* 召回流程:
    - 1. 用户输入
    - 2. 获取会话id（没有则创建）
    - 3. 语音转文字
    - 4. 文字通过cn2an 统一标准化输入文本（五百万=500万）
    - 5. 设置相似度阈值(可配置)
        - 1. faiss构建时开启了 L2 归一化,向量长度为1，检索返回的 distance = 欧式距离 d，所以相似度阈值需要转距离
        - 2. 转换公式：score_threshold = math.sqrt(2 * (1 - base_score_threshold))
    - 6. 设置TopK（可配置）
        - 1. 阈值越严格（比如 0.8 以上）→ 调高 TopK，多捞一点候选，防止筛完为空；
        - 2. 阈值宽松（0.6 左右）→ 调低 TopK，减少无关数据。
        - 3. 阈值越高、文档越碎片化、有 Rerank → TopK 调大；线上高并发、知识库干净 → TopK 调小
    - 7. 标准化输入文本转向量
    - 8. 向量检索
    - 9. 使用 jieba 对标准化输入文本进行分词，rank-bm25 计算相关性分数
    - 10. 重排（使用 RRF（Reciprocal Rank Fusion）融合两种检索结果）
    - 11. 根据会话id查询上下文与重排结果组合prompt
    - 12. 模型调用工具生成回答
    - 13. 保存上下文
    - 14. 流式输出
* **切片策略（文档上万份如何入库）**
  * 格式转换：markitdown 把 pdf/doc/pptx/xlsx 统一转 md，保留标题层级与表格结构
  * 结构优先切片：先按 markdown 标题（#、##、###）分块，再在长块内用 RecursiveCharacterTextSplitter 二次切，避免切断语义
  * 切分参数：chunk_size 一般取 embedding 模型 context 的 1/4（bge-m3 取 512-1024），chunk_overlap 取 10%-20%（约 50-100）
  * 元数据：每片携带 {kb_id, doc_id, source, page, chunk_index, title_path}，便于检索后过滤与引用展示
  * 批量 embedding：langchain `embeddings.embed_documents(list)`，batch_size 控制在 32-64，配合 `asyncio.Semaphore` 防限流
  * 异步流水线：`markitdown → 切片 → embedding → 写 faiss` 全异步，CPU 密集（切片）用 `run_in_executor` 丢线程池
  * 索引选择：<10w 用 `IndexFlatIP`（精度最高）；>10w 改 `IndexIVFFlat`（nlist≈√N，nprobe 10-20）；再大上 HNSW 或 milvus
  * 索引隔离：每个知识库独立 faiss 文件 + 独立 id 映射 json，目录 `indexes/{kb_id}/`
  * 持久化：`faiss.write_index` + 切片原文存 sqlite/ES，支持按 chunk_id 反查原文
  * 增量更新：文档版本号对比，先 `remove_ids` 再 add；删除用软删标记+定期重建
  * 断点续传：每个 doc_id 写入前在 redis 打标 `kb:doc:processed`，崩溃后可重跑未完成部分
  * 上万份入库耗时：embedding 是瓶颈，bge-m3 单卡 A10 约 2000 docs/min，加 vllm/embedding 服务可达 1w+/min
* **模型选型、集成与性能优化**
  * 选型矩阵（边端智能助手场景）：
    * embedding：BGE-M3（中文/多语言/8192 上下文）、bge-large-zh-v1.5（768 维，精度高）、m3e-large（轻量）
    * chat：Qwen2.5-7B/14B、DeepSeek-V3、ChatGLM4（中文会议场景好）
    * asr：Whisper-large-v3、Paraformer（阿里，中文准）、SenseVoice（多语种+情感）
    * rerank：BGE-Rerank-v2-M3、Cohere Rerank
    * 工具/函数：支持 function call 的模型（Qwen2.5、GPT-4o、DeepSeek）
  * 集成方式：统一 OpenAI 兼容协议，私有模型用 vllm/xinference 起服务，langchain 用 `ChatOpenAI(base_url=...)` 接入
  * 性能优化：
    * embedding 批量化 + GPU 推理（text2vec、FlagEmbedding）
    * faiss 走 C++ 底层，比 numpy 快 50-100 倍
    * prompt 缓存：相同 system + tools 走 prefix cache（vllm 支持）
    * 多知识库并行检索：`asyncio.gather` 各库独立查，再 RRF 融合
    * 流式输出：SSE 降低首字延迟到 200ms 内
    * 连接池：httpx.AsyncClient 复用，避免每次握手
    * 模型分级：简单 query 走 7B，复杂/长文本走 72B/API，节省成本
* **上下文应用**
  * 多轮上下文：以 session_id 为 key 从 redis 拉历史 messages，限制最大 N 轮（如 10）
  * 系统上下文：人设 prompt + 知识库说明（"你只能基于以下参考内容回答"）+ 工具描述
  * 召回上下文：RRF 重排后的 topK 文档以 `[1] xxx\n[2] xxx` 格式注入，附 source/title
  * 上下文拼接模板：`system_prompt + history(messages) + retrieved_context + user_input`
  * 引用约束：提示词明确"回答末尾标注引用编号 [1][2]"，便于溯源
  * 会议纪要场景：额外注入"纪要模板"作为 few-shot，控制输出结构
* **token控制**
  * 三段预算：prompt（系统+历史+召回） + 输出（max_tokens） ≤ 模型 context_window，预留 5%-10% buffer
  * 召回侧：topK 控制在 5-10，单 doc 截断到 500 token，召回拼接后再总截
  * 历史侧：滑动窗口保留最近 6-10 轮，更早的调用 LLM 摘要成 1-2 段历史摘要
  * tiktoken 精确计数：`encoding.encode(text)` → len，OpenAI 模型用 `cl100k_base`，Qwen 用模型自带 tokenizer
  * 超限策略：保留 system + 摘要 + 最近 3 轮 + user，截中间历史（"Lost in the middle" 现象缓解）
  * 多知识库路由：每个库只取 top3 再合并 rerank，避免 token 爆炸
* **function call（参数schema、调用链路、错误处理）**
  * schema 定义：项目里用 Pydantic BaseModel 描述工具入参，langchain `@tool` 装饰器自动生成 OpenAI function schema；示例：
    ```python
    class SwitchKB(BaseModel):
        kb_id: str = Field(description="目标知识库ID")
        reason: str = Field(description="切换原因")
    @tool(args_schema=SwitchKB)
    def switch_knowledge_base(kb_id: str, reason: str) -> str: ...
    ```
  * 调用链路（与本项目召回流程对接）：
    1. user 输入 → 判定是否需要工具（路由 agent / LLM self-decide）
    2. 模型返回 `tool_calls=[{name, args}]`
    3. 客户端解析 args，Pydantic 二次校验（防模型幻觉参数）
    4. 路由到本地函数执行（同步/异步）
    5. 结果封装为 `ToolMessage(tool_call_id=..., content=...)` 追加到 messages
    6. 再次 invoke 模型 → 流式输出最终回答
  * 项目内置工具集：`switch_kb`（切知识库）、`search_kb`（精确检索）、`create_meeting_minutes`（生成纪要）、`web_search`（联网）
  * 错误处理：
    * 参数缺失/类型错：捕获 ValidationError，把错误信息以 tool 形式回传，让模型自纠正（最多 2 轮）
    * 执行超时：`asyncio.wait_for(tool, timeout=10)`，超时返回 `{"error": "timeout"}` 让模型降级回答
    * 工具结果为空：返回 `"未找到相关内容"`，由模型决定兜底（提示用户切换知识库或换 query）
    * 循环调用：限制 max_iter=5，超过强制结束；记录 trace 便于排查
    * 安全沙箱：危险工具（删文档/改配置）走人工确认，返回需要审批的 prompt
* **mcp例子**
  * MCP（Model Context Protocol）= 模型 ↔ 工具的统一协议，类似"工具界的 USB"
  * 项目里用 MCP 把企业微信/钉钉/飞书通知、日历查询、数据库查询封成标准 tool，多模型（Qwen/GPT/Claude）都能直接调用
  * 实际例子：`mcp_server_meeting.py`
    ```python
    from mcp.server import Server, stdio
    app = Server("meeting-tools")
    @app.list_tools()
    async def list_tools(): return [Tool(name="create_meeting",
        description="预定会议室", inputSchema={...})]
    @app.call_tool()
    async def call_tool(name, arguments): ...
    ```
  * 客户端用 langchain-mcp-adapters 接入：`MultiServerMCPClient({...})` → `agent = create_react_agent(llm, tools)`
  * 优势：工具与模型解耦，新增工具不用改 agent 代码；一套 MCP 服务可被多个 agent 复用
* **检索有哪些方式**
  * 稀疏检索（Sparse）：BM25、TF-IDF、Elasticsearch倒排索引；强项是关键词精确匹配
  * 稠密检索（Dense）：向量相似度，faiss/milvus/chroma；强项是语义理解
  * 混合检索（Hybrid）：本项目用的就是 dense + BM25，RRF 融合
  * 全文检索：倒排+分词，Lucene/ES
  * 知识图谱：neo4j 做 entity-relation 检索，适合结构化关系问答
  * 层级检索（Hierarchical）：父块定位→子块精读（适合长文档）
  * 元数据过滤：先按 kb_id/时间/部门过滤候选，再做向量检索（缩小搜索空间，提速+提准）
  * 图增强 RAG：KG + vector 混合，GraphRAG 适合"总结整批文档"
  * HyDE（假设性文档嵌入）：让 LLM 先生成假设答案再 embed，弥补 query-doc 语义 gap
  * 多 Query 召回：让 LLM 生成 3-5 个改写 query 并行检索，扩大召回
* **rerank重排有哪些方式**
  * 统计法：BM25、TF-IDF 单独打分（项目里就是 BM25 这一路）
  * 学习型（Cross-Encoder）：query 和 doc 拼接输入 BERT，输出相关性分
    * BGE-Reranker-v2-M3（中文首选）
    * Cohere Rerank-3（闭源，准但贵）
    * monoT5、T5-Reranker
    * jina-reranker
  * LLM 重排：把 topK 文档让 GPT-4 评 1-5 分，CoT 打分（贵但准，适合小批量）
  * 融合法：
    * RRF（Reciprocal Rank Fusion）：`score = Σ 1/(k+rank_i)`，k 常取 60，无监督、本项目用
    * 加权融合：`α * vector_sim + (1-α) * bm25_score`，α 需 grid search
    * Convex Combination：对分数做 min-max 归一再加权
  * 排序学习：LambdaMART、RankNet（训练阶段用）
  * 项目实践：vector top100 + bm25 top100 → RRF 取 top20，bge-reranker 再精排 top5
* **embedding model怎么选**
  * 中文主流选择（按场景）：
    * BGE 系列（BAAI）：bge-large-zh-v1.5（768 维，C-MTEB 强）、bge-m3（1024 维，多语言、长文本 8192）、bge-small-zh（512 维，轻量）
    * M3E：moka-ai/m3e-large，句向量友好
    * 智源 Aquilabed、达摩院达摩、GTE-Qwen2
  * 选型维度：
    * 领域匹配：通用 vs 垂直（医疗/BioBERT、法律）
    * 维度：精度 vs 存储/速度
    * 上下文：长文档必须支持（bge-m3 8192）
    * 速度：QPS、batch 吞吐
    * license：商用是否可（bge-m3 MIT）
    * MTEB/C-MTEB 榜单分数
  * 私有化部署：FlagEmbedding 库，`from FlagEmbedding import BGEM3FlagModel`，支持微调
  * 评估方法：项目里自建 200 条 query-doc 对，Recall@10 > 0.85 才上线
* **embedding维度怎么选**
  * 常见维度：384（mini）、512（small）、768（base）、1024（large）、1536（OpenAI ada）、3072（OpenAI large）
  * 选型 trade-off：
    * 维度↑ → 表达力↑、存储↑、检索速度↓、faiss 内存↑
    * 经验公式：百万级向量 × 1024 维 ≈ 4GB 内存（float32）
  * 业务参考：
    * 知识库 < 10w：768-1024
    * > 100w：512 + 量化（int8）省 4 倍
    * 会议纪要、FAQ：768 足够
  * 一致性约束：训练/入库用 A 模型，查询必须用同一模型；切模型需重新入库
  * 降维：UMAP/PCA 可视化用，生产不建议（丢精度），量化（PQ/SQ）更划算
  * 项目实践：bge-m3 1024 维，faiss 选 `IndexFlatIP`（内积=余弦，归一化后）
* **top20 怎么来的，如何处理噪声**
  * 20 的确定依据：
    * 经验值：业内 RAG 召回→rerank 两阶段常用 top100→top10/20
    * token 预算：单 doc 500 token × 20 = 10k token，留足历史+输出空间
    * 业务效果：A/B 测试，top5/10/20/50 对照人工评分，20 是召回率和 prompt 长度的甜点
    * 配合 rerank：粗召 100，bge-reranker 精排 20 进 prompt
  * 噪声来源：query-doc 语义 gap、切片过粗、embedding 维度不够、文档质量差
  * 噪声处理：
    * 提高相似度阈值：cosine > 0.7（转 distance 阈值用 L2 归一公式）
    * rerank 重排：bge-reranker 把不相关的踢出去
    * MMR（最大边际相关性）：λ 控制相关性与多样性，0.7 偏相关，0.3 偏多样
    * 元数据过滤：按时间/部门/来源缩小候选
    * LLM 上下文压缩：让模型从 topK 抽取关键句再做二次回答
    * prompt 约束："仅基于以下高度相关内容回答，未提到则说不知道"
    * 多路召回去重：vector + bm25 各自取，RRF 融合天然去重
  * 项目实践：vector top100 + bm25 top100 → RRF → top20 → bge-reranker → top5 进 prompt
* **rag结果怎么评分、优化**
  * 评分指标三层：
    * 检索侧：Recall@K、MRR、NDCG@K、Hit Rate
    * 生成侧：faithfulness（忠实度）、answer_relevancy、context_relevancy
    * 端到端：用户采纳率、点赞率、Badcase 率
  * 工具：ragas 框架（自动跑上述指标）、LangSmith（链路追踪+标注）
  * LLM-as-Judge：GPT-4 按 1-5 分打分，打分 prompt 要给 rubric
  * 评估集：200-500 条人工标注 QA，覆盖多知识库、长尾 query、badcase
  * 优化方向（按 ROI 排序）：
    1. 切片：调 chunk_size/overlap，结构化优先
    2. embedding：换 bge-m3 或领域微调
    3. 加 rerank：bge-reranker 提升 10-20%
    4. query 改写：HyDE / multi-query
    5. 加引用：减少幻觉
    6. prompt 工程：few-shot、结构化输出
    7. 微调：在垂直领域数据上 SFT/LoRA
* **出现过什么错误，怎么排查什么原因，怎么处理**
    * **幻觉问题**
      * 现象：模型回答里出现知识库里没有的内容，编造 API/参数/人名
      * 原因：模型先验太强、检索为空时仍强行回答、上下文窗口被无关文档占满
      * 排查：对比 answer 与 retrieved docs 的 ragas faithfulness 分；看 recall 是否为空
      * 处理：
        * 提示词加约束："若参考内容不足，回答'我不知道'，不要编造"
        * 强制引用：要求回答末尾标 `[1][2]`，可点击溯源
        * 检索为空时走兜底分支（不调用 LLM 生成）
        * 提高阈值、增加 query 改写
        * 输出后用 LLM-as-Judge 检测 hallucination，命中则重生成
    * **上下文干扰**
      * 现象：topK 中混进无关文档，模型被带偏答非所问
      * 原因：切片粒度太粗、embedding 区分度差、阈值过低
      * 排查：人工看 topK 命中率，算 context_precision
      * 处理：
        * 细化切片，chunk_size 调小
        * 引入 rerank
        * MMR 多样性去冗余
        * 元数据预过滤（按 kb_id/时间）
        * 提示词："忽略与问题无关的段落"
    * **长上下文退化**
      * 现象：召回 20 条后效果反而比 5 条差（Lost in the middle）
      * 原因：attention 分散在大量低相关 token，关键信息被淹没
      * 排查：对比不同 K 值的 ragas 分；用 attention 可视化看权重分布
      * 处理：
        * 召回数控制在 5-10，不要贪多
        * rerank 把最相关的放首尾（位置 bias）
        * LLM 上下文压缩：先抽关键句再回答
        * 多路召回后分而治之：每路单独生成再汇总（Map-Reduce）
    * **工具调用失效**
      * 现象：参数错（字符串当 int）、死循环调用、忽略工具结果
      * 原因：schema 描述不清、模型理解偏差、错误未回传
      * 排查：langsmith trace 看每次 tool_call 的 args/response；统计参数错误率
      * 处理：
        * tool description 写详细，给参数示例
        * few-shot：在 system 放"用户问 X → 调用 tool Y → 拿到 Z → 回答 W"
        * Pydantic 校验失败时把 error 信息塞回 messages，让模型重试
        * 限制 max_iter=5 防死循环
        * 工具执行加 timeout（10s）和幂等保证
2. **agent**
* **transformer**
  * 架构：Encoder-Decoder（原始）/ Decoder-only（GPT/Llama/Qwen）/ Encoder-only（BERT）
  * 核心模块：Multi-Head Self-Attention + FFN + Residual + LayerNorm
  * 项目落地：所有 LLM（Qwen2.5、ChatGLM4、DeepSeek）都是 Decoder-only Transformer；BERT 用于 cross-encoder rerank
* **attension机制**
  * Self-Attention：`Q·K^T / √d → softmax → ·V`，捕捉序列内依赖
  * Multi-Head：多组 QKV 并行，捕捉不同子空间关系
  * Cross-Attention：Q 来自一个序列、K/V 来自另一序列（rerank 里 query-doc 拼接就是 cross）
  * Causal Mask：自回归生成时屏蔽未来 token
  * 优化：Flash Attention（显存省 5x、速度 2-4x）、Paged Attention（vllm 推理）
  * GQA / MQA：分组查询注意力，省 KV cache
  * 项目里用：rerank 模型本质就是 cross-encoder attention
* **token核心概念**
  * 分词算法：BPE（GPT 系）、WordPiece（BERT）、SentencePiece（多语言）
  * 1 token ≈ 1.5 个中文字 或 0.75 个英文单词
  * 工具：tiktoken（OpenAI）、transformers tokenizer（HuggingFace）
  * 成本计费：按 input/output token 数
  * 项目里：限流、计费、context 控制都基于 token 计数
* **embedding核心概念**
  * 文本→稠密向量的映射，把语义编码到低维空间
  * 距离度量：cosine 相似度（归一化后=内积）、L2 欧式距离、dot product
  * 训练目标：对比学习（contriever、bge）、MLM（bert）
  * 项目里：embedding 是 RAG 的"索引基础"，bge-m3 把 8192 token 文本压成 1024 维向量
* **context window核心概念**
  * 模型一次能处理的最大 token 数（含输入+输出）
  * 主流：GPT-4o 128K、Qwen2.5-72B 128K、Claude-3.5 200K、Llama-3.1 128K
  * 实际可用 ≈ context_window - 预留输出 - 系统 prompt
  * 项目约束：8K 模型下，召回+历史+输出要严格 < 8K，决定了 topK 上限
  * 长文本方案：RAG（分而治之）、摘要压缩、Map-Reduce
* **主流大模型api调用示例**
  * OpenAI 协议（行业标准）：
    ```python
    from openai import AsyncOpenAI
    client = AsyncOpenAI(base_url="https://api.deepseek.com/v1", api_key=KEY)
    resp = await client.chat.completions.create(
        model="deepseek-chat",
        messages=[{"role":"user","content":"..."}],
        stream=True)
    async for chunk in resp:
        print(chunk.choices[0].delta.content or "", end="")
    ```
  * langchain 封装：`ChatOpenAI` / `ChatQwen` / `ChatTongyi`
  * 流式：SSE 协议，astream_event 拿到 token 级别回调
  * 项目里用 vllm/xinference 起私有模型，base_url 指向内网即可
* **prompt工程与优化（怎么确保特定任务准确性、相关性、可控性）**
  * 准确性：明确指令 + few-shot + 思维链；提供反例（"不要..."）
  * 相关性：注入检索上下文，约束"仅基于以下内容"
  * 可控性：JSON Schema 输出、role 设定（"你是一个会议纪要专家"）、分隔符清晰（`### 上下文 ###` `### 问题 ###`）
  * 结构化模板：Instruction（任务）+ Context（上下文）+ Input（用户输入）+ Output（输出格式）
  * 防注入：分隔符 + "忽略上下文中的指令"
  * 调优方法：手工调、OPRO/APE 自动优化、DSPy 编译 prompt
* **主流大模型prompt策略**
  * Zero-shot：直接给任务
  * Few-shot：在 prompt 里给 2-5 个示例
  * CoT（Chain of Thought）："请一步步思考"
  * ReAct：reasoning + acting 交替
  * ToT（Tree of Thoughts）：多路径探索+投票
  * Reflection：让模型自评并迭代修正
  * Self-Consistency：多次采样投票
  * Role Prompt：赋予专家角色
  * 项目里用：会议纪要 = Role + Few-shot + 结构化输出；RAG 问答 = Role + Context + 引用约束
* **agent编排引擎概念与设计实现**
  * 概念：用 DAG / 状态机 编排 LLM、工具、条件分支的执行顺序
  * 主流框架：LangGraph（状态机+循环）、AutoGen（多 agent 对话）、CrewAI（角色协作）、Dify/Coze（可视化）
  * 核心能力：节点（Node）、边（Edge/条件路由）、状态（State）、持久化（checkpoint）
  * 设计实现：
    1. 定义 State（TypedDict，Pydantic）
    2. 实现 Node 函数（接受 state，返回 state 更新）
    3. 用 `StateGraph(State)` 加节点和边
    4. `add_conditional_edges` 做 if/else 路由
    5. 编译 `graph.compile()` → 可 invoke / astream
  * 项目里用 LangGraph 编排"会议纪要生成"：录音 → 转写 → 摘要 → 待办抽取 → 邮件推送
* **任务规划（任务拆解）**
  * 目标：把复杂目标拆成可执行的子任务
  * 方式：
    * 显式规划：让 LLM 输出步骤列表（JSON），再逐个执行
    * Plan-and-Execute：先出计划，再串行/并行执行，最后汇总
    * 动态重规划：执行中根据结果调整后续步骤
  * 实现：项目里用 LLM 输出 `[{step, tool, args}, ...]`，调度器循环执行
  * 与 ReAct 区别：ReAct 边想边做，Plan-Execute 先想后做，适合长链路
* **state状态管理**
  * State 是 agent 的"工作内存"，在节点间流转
  * 结构：`{messages: [...], current_step, kb_id, retrieved_docs, ...}`
  * 实现：LangGraph 用 `StateGraph` + reducer（默认覆盖，可自定义 add_message 追加）
  * 持久化：Checkpoint（sqlite/postgres），断点恢复
  * 项目里：会议纪要 state 包含 `transcript / summary / todos / recipients`
* **记忆机制（短期记忆、长期记忆设计以及应用场景）**
  * 短期：当前 session 的 messages（redis 存，TTL 1 天）
  * 长期：用户偏好、历史总结、行为画像（向量库 + DB）
  * 场景：
    * 短期用于多轮对话连贯性
    * 长期用于个性化（如"用户上次偏好 markdown 格式"）
  * 实现：项目里 redis 存短期（key=`session:{id}`），milvus 存长期 user_profile
  * 记忆压缩：每 10 轮调 LLM 总结成 1 段历史摘要
  * 召回：从长期记忆里 RAG 取出相关片段注入 prompt
* **工作流编排概念与设计实现**
  * 概念：用可视化/代码方式编排多步骤流程
  * 元素：节点（LLM/工具/条件）、边（流转）、触发器、错误处理
  * 框架：Dify/Coze（低代码）、LangGraph/Temporal（代码）
  * 项目里：会议纪要工作流 = [上传音频] → [ASR 转写] → [LLM 摘要] → [LLM 抽待办] → [人审] → [邮件/IM 推送]
  * 实现：用 LangGraph 编排，每个节点是 async 函数
* **多轮对话设计与实现**
  * session 管理：session_id 关联 redis 里的 messages list
  * 指代消解：把"它/那个"改写成上文实体（LLM 改写或 NER）
  * 上下文压缩：超 token 时 LLM 摘要
  * 主题切换检测：embedding 相似度 < 阈值时开新 session
  * 项目里：用户问"那它的参数呢？" → 自动补全成"XXX 模型的参数是？"
* **错误恢复机制设计与实现**
  * 错误类型：模型超时、工具失败、参数错、token 超限、限流
  * 恢复策略：
    * 重试：tenacity 库，指数退避（1s, 2s, 4s），最多 3 次
    * 降级：主模型挂了切备用模型（Qwen → GPT）
    * 回滚：事务性步骤回退已做的修改
    * 兜底：工具失败时让 LLM 用通用知识回答
  * 实现：项目里每个 agent step 包 try/except，失败时记录 trace + 走兜底
  * 熔断：连续失败 N 次熔断 5 分钟，保护下游
* **回复判断机制与实现**
  * 判定类型：是否需要检索、是否需要工具、是否需要切知识库、是否要追问
  * 实现：
    * 路由 LLM：分类器（zero-shot 或小模型 BGE-M3 + classifier）
    * 规则兜底：包含"查/搜"→ 检索，包含"打开/关闭"→ 工具
  * 项目里用小 LLM（Qwen2.5-3B）做意图路由，节省 token
* **意向识别机制与实现**
  * 识别：会议纪要 / 知识问答 / 闲聊 / 工具指令 / 投诉
  * 方案：
    * 分类 prompt + few-shot（简单）
    * 小模型微调（BERT + 分类头，准）
    * LLM 输出 enum（灵活）
  * 项目里：用户首条消息过分类器，路由到对应 agent（MeetingAgent/QAAgent/ToolAgent）
  * 多轮修正：用户说"不对，我要查订单"→ 切到 QAAgent
* **推理策略（推理连（ReAct、Plan-Execute）和思维链（Cot）等）**
  * CoT：让模型逐步推理，"让我们一步步想"
  * ReAct：Thought → Action → Observation 循环，agent 经典范式
    ```python
    Thought: 需要查知识库
    Action: search_kb(query="XXX")
    Observation: [文档内容]
    Thought: 信息够回答了
    Final Answer: ...
    ```
  * Plan-Execute：先列计划再执行，适合长链
  * Reflexion：执行后自我反思，更正重试
  * Tree of Thoughts：多路径搜索+投票
  * 项目里：简单问答用 CoT，工具调用用 ReAct，长流程用 Plan-Execute
* **架构模式**
  * 核心组件（项目映射）：
    * 规划模块（planning）：LLM 拆解任务 → `planner.py`
    * 记忆模块（memory）：redis 短期 + 向量库长期 → `memory/`
    * 工具调用模块（tool-use）：MCP + langchain tools → `tools/`
    * 行动模块（action）：执行 tool + 输出生成 → `executor.py`
    * 感知模块（perception）：ASR、OCR、用户输入解析 → `perception/`
  * 架构模式：
    1. **单体工作流**
       * ReAct：单 agent 边想边做，适合工具少的场景
       * tool-use：纯 function call，无显式 reasoning
       * planning：Plan-and-Execute，长链路
    2. **多智能体协作**：主 agent 分派子 agent（检索 agent / 写作 agent / 审核 agent），各司其职
    3. **生产级长时运行**：断点续传、checkpoint、人工审批、状态恢复
* **出现过什么问题，怎么定位，线上怎么持续优化**
  * 常见问题：
    * 召回不准：query-doc 语义 gap → 加 query 改写
    * 模型超时：vllm OOM/限流 → 加 retry + 备用模型
    * 成本高：长 context 浪费 → 加压缩+缓存
    * 幻觉：检索为空仍生成 → 加空召回兜底
    * 工具循环：max_iter 失控 → 强制熔断
  * 定位手段：
    * LangSmith / Phoenix trace 全链路
    * 日志：每次请求记 query / retrieved / response / 耗时
    * 监控：成功率、token 用量、P95 延迟、badcase 率
  * 持续优化：
    * 每周跑评估集（ragas + 人工）
    * 收集 badcase 入库反哺切片/检索
    * A/B 测试新 prompt / 新模型
    * 用户反馈闭环（点赞/点踩）
* **怎么云端部署或本地GPU部署**
  * 云端部署：
    * vllm / TGI（HuggingFace）：高吞吐推理，单卡 100+ QPS
    * 阿里云 PAI / 火山方舟 / 腾讯 TI：一键部署 Qwen/DeepSeek
    * AWS SageMaker / Azure ML：海外
    * k8s + Helm chart 编排多模型
  * 本地 GPU 部署：
    * ollama：开箱即用，`ollama run qwen2.5:7b`
    * xinference：兼容 OpenAI 协议，支持多种模型
    * LM Studio：图形化
    * llama.cpp / vllm：高性能，命令行
  * 硬件选型：
    * 7B 模型：16GB 显存（4090/3090）
    * 14B：24GB（4090）
    * 72B：2×A100 80G 或 4×4090（量化后）
  * 项目里：embedding 走 CPU/GPU 轻量服务，rerank 走 GPU 小服务，chat 走 vllm 大服务，三层解耦独立扩缩容
* **流式输出**
  * 用什么方式逐token处理输出：sdk，sse
  * 三层管道：字节流（网络层）-> 事件流（协议层）-> 状态机（语义层）-> 业务逻辑
  * sse：每条消息以data: 开头，后跟 token 内容，最后以\n\n 结束,一次read()可能读取不完整，需要保存下来等下一次再拼接上，最后解析渲染出来
