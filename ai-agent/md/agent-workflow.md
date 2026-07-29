# Agent 编排引擎

> 一句话：编排引擎是复杂任务的"交通指挥中心"，把多个 Agent/Skill/Plugin 按流程调度起来。

---

## 一、什么是编排引擎？

**一句话**：编排引擎是 Agent 的工作流调度底座，负责把复杂任务拆解成多个节点，按规则串行/并行/分支/循环执行。

### 为什么需要编排引擎？

| 问题 | 说明 |
|---|---|
| 单 Agent 不可控 | 自由迭代容易跑偏、循环 |
| 能力单一 | 一个 Agent 无法同时擅长检索、分析、创作 |
| 成本高 | 简单步骤也走 LLM 推理 |
| 难维护 | 流程无版本、无配置、难治理 |

---

## 二、编排引擎五层架构

```
┌─────────────────────────────────────┐
│  流程配置层  │  JSON/YAML/可视化定义   │
├─────────────────────────────────────┤
│  核心调度层  │  规划器、调度器、状态机  │
├─────────────────────────────────────┤
│  能力适配层  │  适配 Agent/Skill/Plugin │
├─────────────────────────────────────┤
│  执行计算层  │  各专业执行节点          │
├─────────────────────────────────────┤
│  治理观测层  │  日志、监控、重试、降级  │
└─────────────────────────────────────┘
```

---

## 三、四大编排模式

| 模式 | 说明 | 例子 |
|---|---|---|
| **顺序编排** | 节点依次执行 | 解析 → 分片 → 入库 → 索引 |
| **并行编排** | 无依赖节点同时执行 | 同时查知识库、联网、调统计 |
| **条件分支** | 根据结果走不同分支 | 简单问题直接答，复杂问题走分析 |
| **循环迭代** | 满足条件才退出 | 多轮检索补全、反复校验 |

```
顺序：A → B → C
并行：A ─┬─ B
       └─ C
分支：A → 条件判断 → B 或 C
循环：A → B → 满足条件？→ 是结束 / 否继续
```

---

## 四、中心化 vs 去中心化

| 特性 | 中心化编排 | 去中心化协作 |
|---|---|---|
| 控制权 | 统一调度 | Agent 自主移交 |
| 稳定性 | 高 | 低 |
| 可追溯 | 强 | 弱 |
| 灵活性 | 中 | 高 |
| 适用场景 | 企业级生产 | 实验性探索 |

**面试建议**：企业级生产优先中心化编排。

---

## 五、执行链路

```
1. 流程解析：加载配置，生成执行蓝图
2. 任务规划：拆解子任务，分配执行单元
3. 前置校验：参数、权限、风控检查
4. 智能调度：按串/并/分支/循环执行
5. 节点执行：Agent/Skill/Plugin 落地
6. 结果聚合：汇总、去重、结构化
7. 异常容错：重试、熔断、降级
8. 收尾输出：格式化、脱敏、归档
```

---

## 六、数据流转机制

**一句话**：用全局上下文数据池统一存取节点输入输出。

```
节点 A 输出 ──► 全局数据池
                    │
                    ▼
            节点 B 按需取参
```

**好处**：节点解耦、流程可自由组合、新增节点不改旧代码。

---

## 七、高可用保障

| 维度 | 方案 |
|---|---|
| 故障隔离 | 单节点失败不影响整体 |
| 超时管控 | 每节点独立超时 |
| 熔断降级 | 高频失败节点自动熔断 |
| 状态持久化 | 支持中断续跑、重试恢复 |
| 限流 | 全局 QPS 控制 |

---

## 八、State 状态管理

### 1. 为什么需要 State？

编排引擎执行长流程时，需要记录每个节点的输入、输出、状态、异常信息，否则：
- 流程中断后无法续跑
- 多轮迭代无法拿到上一步结果
- 失败时无法定位问题

### 2. State 常见结构

```json
{
  "workflow_id": "wf_123",
  "status": "running",
  "current_node": "node_b",
  "context": {
    "user_query": "查询北京天气",
    "node_a_output": "北京"
  },
  "nodes": {
    "node_a": {
      "status": "success",
      "input": {"city": "北京"},
      "output": "北京",
      "start_time": "2025-01-01T10:00:00",
      "end_time": "2025-01-01T10:00:01"
    },
    "node_b": {
      "status": "running",
      "input": {"city": "北京"},
      "output": null,
      "retry_count": 0
    }
  }
}
```

### 3. State 常见实现方式

| 方式 | 说明 | 适用场景 |
|---|---|---|
| 内存 State | 用 `map` 或对象存当前流程状态 | 短流程、单实例 |
| Redis | 键值存储，支持过期和分布式 | 多实例、需中断续跑 |
| 数据库 | MySQL/PostgreSQL 持久化 | 长流程、强一致性 |
| 事件日志 | Kafka + 事件溯源 | 高并发、需回放 |

### 4. State 核心操作

```python
class WorkflowState:
    def __init__(self, workflow_id: str, store: dict):
        self.workflow_id = workflow_id
        self.store = store  # 可以是内存 dict 或 Redis

    def get_node_state(self, node_id: str):
        return self.store.get(self.workflow_id, {}).get("nodes", {}).get(node_id)

    def set_node_state(self, node_id: str, state: dict):
        wf = self.store.setdefault(self.workflow_id, {"nodes": {}})
        wf["nodes"][node_id] = state

    def update_context(self, key: str, value):
        wf = self.store.setdefault(self.workflow_id, {"context": {}})
        wf["context"][key] = value

    def get_context(self):
        return self.store.get(self.workflow_id, {}).get("context", {})
```

### 5. 中断续跑示例

```python
def run_workflow(workflow_id, config):
    state = load_state(workflow_id) or create_state(workflow_id)

    for node in config["nodes"]:
        node_id = node["id"]

        # 已经成功的节点跳过
        if state.get_node_state(node_id)?.status == "success":
            continue

        try:
            result = execute_node(node, state.get_context())
            state.set_node_state(node_id, {
                "status": "success",
                "output": result,
            })
            state.update_context(f"{node_id}_output", result)
        except Exception as e:
            state.set_node_state(node_id, {
                "status": "failed",
                "error": str(e),
                "retry_count": state.get_node_state(node_id).get("retry_count", 0) + 1
            })
            if should_retry(node, state):
                run_workflow(workflow_id, config)  # 重试
            else:
                handle_failure(state)
            break
```

### 6. State 持久化流程

```
1. 节点执行前：写入状态为 running
2. 节点执行中：实时保存上下文
3. 节点成功后：更新状态为 success，保存输出
4. 节点失败后：更新状态为 failed，记录异常
5. 流程重启时：读取 State，从失败的节点继续执行
```

---

## 九、面试高频问题

### Q1：什么是 Agent 编排引擎？

**答**：编排引擎是复杂任务的工作流调度底座，把多个 Agent/Skill/Plugin 按规则串联/并联执行，解决单 Agent 不可控、能力单一、成本高等问题。

### Q2：编排引擎有哪些模式？

**答**：顺序、并行、条件分支、循环迭代。

### Q3：为什么企业级优先中心化编排？

**答**：中心化可控、可追溯、可治理、易运维；去中心化灵活但容易混乱，不适合线上业务。

### Q4：节点间数据怎么流转？

**答**：通过全局上下文数据池，每个节点把输出写入数据池，后续节点按需读取。

### Q5：编排引擎怎么做容错？

**答**：节点级故障隔离、超时控制、熔断降级、状态持久化、限流。

### Q6：State 在编排引擎里起什么作用？

**答**：State 记录流程每个节点的输入、输出、状态、异常，支持中断续跑、重试恢复、问题定位。常用实现方式有内存 State、Redis、数据库、事件日志。

### Q7：中断续跑怎么实现？

**答**：节点执行前后持久化 State，流程重启时先读取 State，已经 success 的节点跳过，从失败或 running 的节点继续执行。

---

## 十、项目口述模板

> 请介绍你项目中的编排引擎。

我在项目中落地了中心化 Agent 编排引擎。通过 JSON/YAML 配置定义工作流，支持顺序、并行、条件分支、循环迭代四种编排模式。引擎负责任务拆解、节点调度、数据流转和结果聚合，下层统一适配子 Agent、Skill、Plugin 和 RAG 能力。节点间通过全局上下文数据池解耦，新增流程无需改核心代码。同时用 State 记录每个节点的输入输出和状态，支持中断续跑和失败重试，配套节点级故障隔离、超时重试、熔断降级和状态持久化，保证复杂流程稳定可控。

---

## 十一、一句话总结

- **编排引擎 = 复杂任务的调度指挥中心**
- **四大模式：顺序、并行、分支、循环**
- **企业级选中心化，可控可治理**
- **数据池解耦节点，异常隔离保稳定**
