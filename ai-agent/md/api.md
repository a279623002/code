# 国内外主流大模型API调用规范、差异对比手册
## 目录
1. 整体总览对比表（端点、鉴权、协议、SDK、国内可用性）
2. 各模型详细调用方式（curl + Python示例）
3. 核心关键差异：消息格式、System Prompt、Function Calling、多模态、流式、长上下文
4. 工程落地选型区别（兼容迁移、合规、私有化、Agent适配）
5. 面试高频考点总结

## 一、全维度API总览对比表
|对比维度|OpenAI(GPT)|Claude(Anthropic)|Gemini(Google)|Qwen通义千问(阿里云)|DeepSeek深度求索|
| ---- | ---- | ---- | ---- | ---- | ---- |
|**API主端点**|https://api.openai.com/v1/chat/completions|https://api.anthropic.com/v1/messages|https://generativelanguage.googleapis.com/v1beta/models/{model}:generateContent|兼容模式：https://dashscope.aliyuncs.com/compatible-mode/v1/chat/completions<br>原生DashScope：dashscope.aliyuncs.com/api/v1|https://api.deepseek.com/v1/chat/completions<br>兼容Claude：/anthropic|
|**鉴权Header**|Authorization: Bearer sk-xxx|1. x-api-key: sk-xxx<br>2. anthropic-version: 2023-06-01（必填版本头）|x-goog-api-key: xxx|Authorization: Bearer sk-xxx（兼容模式）<br>原生：Authorization: Bearer dashscope-xxx|Authorization: Bearer sk-xxx|
|**原生协议标准**|OpenAI Chat Completions 行业标准|自研Messages API（不兼容OpenAI）|Google自研generateContent 私有协议|双协议：原生DashScope + 完整OpenAI兼容|100%兼容OpenAI Chat Completions，额外兼容Claude格式|
|**官方SDK**|openai(Python/JS)|anthropic(独立包)|google-generativeai|dashscope + openai SDK复用|直接复用openai官方SDK，无需额外包|
|**国内直连可用性**|❌ 无法直连，需代理/中转|❌ 海外接口，国内访问受限|❌ Google服务国内屏蔽|✅ 阿里云国内合规直连|✅ 国产服务商，国内稳定直连|
|**私有化部署支持**|无开源权重，仅Azure云托管|无开源权重|开源权重(Gemini开源小模型)+Google云API|全系列开源权重本地私有化|全系列开源权重可离线私有化|
|**System Prompt存放**|messages数组中role="system"|顶层独立字段`system`，不进messages数组|generationConfig无system，靠首条content携带|兼容模式同OpenAI(messages.system)；原生支持system字段|完全同OpenAI，messages.role="system"|
|**最大上下文窗口**|GPT-4o 128K|Claude Opus 1M / Sonnet 200K|Gemini 3.5 Flash 1M|Qwen3 Max 128K|DeepSeek V4 1M|
|**Function Calling/工具调用**|原生标准化tools，并行调用稳定|原生tool_use，长上下文不易丢失工具定义|Function Declarations数组，多模态联动弱|原生兼容OpenAI tools，配套Qwen-Agent框架|完整复刻OpenAI tools格式，支持链式工具重试|
|**多模态输入格式**|content数组携带image_url|content内type:image/png/base64|parts数组区分text/inline_data图片|兼容OpenAI vision格式，原生支持视频|仅文本，无原生多模态|
|**流式输出(SSE)**|标准SSE data: {...}，结尾data: [DONE]|自研SSE事件结构(content_block_delta)|SSE流式candidates增量|完全对齐OpenAI流式格式|完全对齐OpenAI流式格式|
|**国产特有扩展参数**|无|无|无|enable_thinking（深度思考开关，extra_body透传）|enable_deep_reasoning推理开关|
|**合规数据出境**|数据上传海外服务器，国内政企受限|海外云端存储，不合规|Google海外，国内禁用|阿里云国内机房，数据不出境合规|国内服务器，数据留境内|

## 二、各模型完整调用方式（curl + Python极简示例）
### 1. OpenAI GPT（行业标准OpenAI格式）
#### curl示例
```bash
curl https://api.openai.com/v1/chat/completions \
-H "Authorization: Bearer sk-OPENAI_KEY" \
-H "Content-Type: application/json" \
-d '{
  "model": "gpt-4o",
  "messages": [
    {"role": "system", "content": "你是专业编程助手"},
    {"role": "user", "content": "写一段快速排序代码"}
  ],
  "temperature": 0.7,
  "stream": false
}'