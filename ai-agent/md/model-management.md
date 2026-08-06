# 模型管理说明

## 1. 整体架构

项目大模型统一入口为 [`app/models/internal/llm_model.py`](file:///d:/Work/nw/langchain-chat/app/models/internal/llm_model.py)。

核心函数 `get_llm_model()` 通过读取配置 `settings.model.run_llm_model`，按名称加载对应模型，并使用全局缓存 `_LLM_MODEL_CACHE` 保证同一进程内只加载一次。

```python
def get_llm_model():
    global _LLM_MODEL_CACHE
    if _LLM_MODEL_CACHE is not None:
        return _LLM_MODEL_CACHE

    run_llm_model = settings.model.run_llm_model
    if run_llm_model == "llm_glm":
        _LLM_MODEL_CACHE = load_glm_model()
    elif run_llm_model == "llm_hm_qwen25_vl":
        _LLM_MODEL_CACHE = load_hm_qwen25_vl_model()
    elif run_llm_model == "llm_qwen25_vl":
        _LLM_MODEL_CACHE = load_qwen25_vl_model()
    else:
        _LLM_MODEL_CACHE = load_glm_model()  # 兜底
    return _LLM_MODEL_CACHE
```

---

## 2. 当前管理的模型

| 配置键 | 模型 | 能力 | 运行环境 | 主要文件 |
|---|---|---|---|---|
| `llm_glm` | 智谱 GLM（chatglm-flash 等） | 纯文本 | 远程 API | [`llm_model.py`](file:///d:/Work/nw/langchain-chat/app/models/internal/llm_model.py#L45) |
| `llm_hm_qwen25_vl` | 后摩 Qwen2.5-VL | 文本 + 图像 | 后摩 NPU（M50，xh1/xh2） | [`hm_qwenvl_model.py`](file:///d:/Work/nw/langchain-chat/app/models/internal/hm_qwenvl_model.py) |
| `llm_qwen25_vl` | 本地 Qwen2.5-VL | 文本 + 图像 | GPU / CUDA | [`gpu_qwenvl_model.py`](file:///d:/Work/nw/langchain-chat/app/models/internal/gpu_qwenvl_model.py) |

### 2.1 智谱 GLM（`llm_glm`）

直接使用 `langchain_openai.ChatOpenAI` 封装，把 `openai_api_base` 指向智谱 OpenAI 兼容接口：

```python
llm_model = ChatOpenAI(
    model_name=settings.model.llm_model_info.llm_glm.llm_model,
    openai_api_key=settings.model.llm_model_info.llm_glm.llm_api_key,
    openai_api_base=settings.model.llm_model_info.llm_glm.llm_base_url,
    temperature=...,
    max_tokens=...,
)
```

### 2.2 后摩 Qwen2.5-VL（`llm_hm_qwen25_vl`）

基于 `tcim_lite` 加载后摩编译后的 `.hmm` 模型：

- `vit_path`：视觉编码器
- `prefill_path`：预填充模型
- `decode_path`：自回归解码模型
- `embedding_path`：量化后的 embedding 权重
- `tokenizer_path`：Qwen2.5-VL tokenizer

加载后返回自定义 `Qwen2VLModel` 对象，不依赖 `transformers`，运行在 NPU 上。

### 2.3 GPU 本地 Qwen2.5-VL（`llm_qwen25_vl`）

使用 `transformers` 原生加载：

- 模型类：`Qwen2_5_VLForConditionalGeneration.from_pretrained`
- 处理器：`AutoProcessor.from_pretrained`
- Tokenizer：`AutoTokenizer.from_pretrained`
- 可选 `flash_attention_2`，支持 AWQ 量化模型

加载后同样返回自定义 `Qwen2VLModel` 对象。

---

## 3. 三种模型的差异

| 维度 | `llm_glm` | `llm_hm_qwen25_vl` | `llm_qwen25_vl` |
|---|---|---|---|
| **依赖** | `langchain_openai` | `tcim_lite` + 后摩运行时 | `transformers`、`qwen_vl_utils` |
| **部署位置** | 云端 API | 后摩 NPU 本地 | GPU 本地 |
| **视觉能力** | 否 | 是 | 是 |
| **推理方式** | HTTP 请求 | NPU 运行 prefill/decode | PyTorch `model.generate` |
| **温度等采样参数** | 支持 | 仅接口占位，实际不支持 | 支持 |
| **流式输出** | LangChain 原生 | 自定义生成器 + 异步队列 | `TextIteratorStreamer` + 线程 |
| **图像输入** | 不支持 | 本地/远程路径 | 本地/远程路径 |

---

## 4. 如何适配 / 兼容 OpenAI

### 4.1 配置层面：GLM 走 OpenAI 兼容协议

智谱 AI 本身提供 OpenAI 兼容的 `/v1/chat/completions` 接口，因此通过 `ChatOpenAI` 即可直接使用，无需额外封装。

### 4.2 输入格式统一

两个 Qwen2.5-VL 本地模型都实现了 `_convert_to_dict_messages()`，支持三类输入：

1. **LangChain Message**：`SystemMessage`、`HumanMessage`、`AIMessage`
2. **字典**：`{"role": "user", "content": "..."}`
3. **元组**：`("user", "...")`

并把 OpenAI 多模态消息中的 `image_url` 自动转换为 Qwen VL 所需的 `image` 类型：

```python
if item.get("type") == "image_url":
    img_url = item.get("image_url", {})
    actual_img_path = img_url.get("url", "") if isinstance(img_url, dict) else str(img_url)
    new_content.append({"type": "image", "image": actual_img_path})
```

这样上层服务只需要构造 OpenAI 格式的消息列表，即可同时用于 GLM API 和本地 VL 模型。

### 4.3 输出格式统一

两个本地 VL 模型都实现了相同的异步接口：

```python
async def astream(messages, max_new_tokens=..., temperature=..., ...) -> AsyncGenerator[GenerationChunk, None]:
    yield GenerationChunk(content="...")

async def generate(messages, ...) -> str:
    full = ""
    async for chunk in self.astream(...):
        full += chunk.content
    return full
```

`GenerationChunk` 是一个简单的 dataclass，与 LangChain 输出一样使用 `.content` 字段。服务层统一按 `chunk.content` 读取，无需关心底层是 API、GPU 还是 NPU。

### 4.4 服务层统一调用

[`app/services/model_manager/qa_service.py`](file:///d:/Work/nw/langchain-chat/app/services/model_manager/qa_service.py) 中统一调用：

```python
self.llm_model = get_llm_model()
async for chunk in self.llm_model.astream(prompt_messages):
    chunk_content = chunk.content.strip() if hasattr(chunk, 'content') and chunk.content else ""
    ...
```

无论是 `ChatOpenAI` 还是自定义 VL 模型，都满足 `astream(messages)` 返回 `chunk.content` 的约定。

### 4.5 响应包装为 OpenAI 格式

`OpenAIChatService` 将模型输出再包装为 OpenAI ChatCompletion / ChatCompletionChunk JSON：

- 非流式：`chat.completion`
- 流式：`chat.completion.chunk`
- 统一包含 `id`、`object`、`created`、`model`、`choices`、`usage` 等字段
- 扩展字段 `extensions` 中返回 `trace_id`、`sources`、`image_count` 等

---

## 5. 配置示例

配置文件参考 [`app/core/config.py`](file:///d:/Work/nw/langchain-chat/app/core/config.py#L60)：

```python
class ModelConfig(BaseModel):
    device: str = "cpu"                    # cpu / hm
    run_llm_model: str = "llm_glm"         # llm_glm / llm_hm_qwen25_vl / llm_qwen25_vl
    supported_models: list = []            # 服务层允许调用的模型名列表
    multimodal_models: list = []           # 哪些模型支持图片输入
    llm_model_info: LLMModelInfoConfig     # 各模型详细配置
    embedding_model: EmbeddingModelConfig  # Embedding 模型配置
    asr_model: ASRModelConfig              # Whisper 语音模型配置
```

`LLMModelInfoConfig` 包含：

- `llm_glm`：`llm_model`、`llm_api_key`、`llm_base_url`、`temperature`、`max_tokens`
- `llm_hm_qwen25_vl`：`vit_path`、`prefill_path`、`decode_path`、`tokenizer_path`、`embedding_path`
- `llm_qwen25_vl`：`model_path`、`device`、`max_memory`、`min_pixels`、`max_pixels`、`default_max_new_tokens`、`default_temperature`

---

## 6. 扩展新模型的建议

如需再接入一种新模型（如 DeepSeek、通义千问 API、本地 Llama 等）：

1. 在 [`app/core/config.py`](file:///d:/Work/nw/langchain-chat/app/core/config.py) 增加对应配置类；
2. 在 `LLMModelInfoConfig` 中新增字段；
3. 在 [`llm_model.py`](file:///d:/Work/nw/langchain-chat/app/models/internal/llm_model.py) 新增 `load_xxx_model()`；
4. 在 `get_llm_model()` 分支中增加判断；
5. 若模型不是 LangChain 原生对象，需要封装 `astream(messages)` / `generate(messages)`，并返回带 `.content` 的 chunk；
6. 在 `model.yaml` 中配置 `run_llm_model` 和对应参数。

核心原则：**上层服务只依赖 `llm_model.astream(messages)` 与 `chunk.content`，底层实现 whatever**。
