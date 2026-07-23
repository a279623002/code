# 树结构与算法面试笔记

> 树是计算机科学中最基础、最重要的数据结构之一。面试中常考的包括二叉树、BST、AVL、红黑树、B/B+树、Trie、堆等，以及与树相关的遍历、递归、动态规划问题。

---

## 一、二叉树（Binary Tree）

### 1. 基本概念

**一句话**：每个节点最多有两个子节点的树。

```
        1          ← 根节点
       / \
      2   3        ← 内部节点
     / \   \
    4   5   6      ← 叶子节点
```

| 概念 | 说明 |
|---|---|
| **根节点** | 最顶层的节点 |
| **叶子节点** | 没有子节点的节点 |
| **高度** | 从该节点到最远叶子的边数 |
| **深度** | 从根到该节点的边数 |
| **满二叉树** | 每个节点要么 0 个子节点，要么 2 个 |
| **完全二叉树** | 除了最后一层，其他层都满，最后一层从左到右填充 |

### 2. 二叉树遍历

```go
type TreeNode struct {
    Val   int
    Left  *TreeNode
    Right *TreeNode
}
```

#### 前序遍历（根左右）

```go
func preorder(root *TreeNode) {
    if root == nil {
        return
    }
    fmt.Println(root.Val)
    preorder(root.Left)
    preorder(root.Right)
}
```

**结果**：1 2 4 5 3 6

#### 中序遍历（左根右）

```go
func inorder(root *TreeNode) {
    if root == nil {
        return
    }
    inorder(root.Left)
    fmt.Println(root.Val)
    inorder(root.Right)
}
```

**结果**：4 2 5 1 3 6

#### 后序遍历（左右根）

```go
func postorder(root *TreeNode) {
    if root == nil {
        return
    }
    postorder(root.Left)
    postorder(root.Right)
    fmt.Println(root.Val)
}
```

**结果**：4 5 2 6 3 1

#### 层序遍历（BFS）

```go
func levelOrder(root *TreeNode) {
    if root == nil {
        return
    }
    queue := []*TreeNode{root}
    for len(queue) > 0 {
        node := queue[0]
        queue = queue[1:]
        fmt.Println(node.Val)
        if node.Left != nil {
            queue = append(queue, node.Left)
        }
        if node.Right != nil {
            queue = append(queue, node.Right)
        }
    }
}
```

**结果**：1 2 3 4 5 6

---

## 二、二叉搜索树（BST）

### 1. 定义

**一句话**：左子树所有节点 < 根节点 < 右子树所有节点。

```
        5
       / \
      3   7
     / \   \
    2   4   8
```

### 2. 性质

- 中序遍历结果是有序数组
- 查找、插入、删除平均 O(log n)，最坏 O(n)（退化成链表）

### 3. 查找代码

```go
func searchBST(root *TreeNode, val int) *TreeNode {
    if root == nil || root.Val == val {
        return root
    }
    if val < root.Val {
        return searchBST(root.Left, val)
    }
    return searchBST(root.Right, val)
}
```

---

## 三、平衡二叉树（AVL）

### 1. 定义

**一句话**：任意节点的左右子树高度差不超过 1 的二叉搜索树。

```
        5              ← 平衡
       / \
      3   7
     /
    2

        5              ← 不平衡
       /
      3
     /
    2
   /
  1
```

### 2. 平衡因子

```
平衡因子 = 左子树高度 - 右子树高度
取值范围：-1, 0, 1
```

### 3. 旋转操作

| 失衡类型 | 旋转方式 |
|---|---|
| LL（左左） | 右旋 |
| RR（右右） | 左旋 |
| LR（左右） | 先左旋，再右旋 |
| RL（右左） | 先右旋，再左旋 |

### 4. AVL vs BST

| 特性 | BST | AVL |
|---|---|---|
| 查找 | 平均 O(log n)，最坏 O(n) | 稳定 O(log n) |
| 插入/删除 | 简单 | 需要旋转维护平衡 |
| 适用 | 查询多、修改少 | 查询和修改都频繁 |

---

## 四、红黑树（Red-Black Tree）

### 1. 定义

**一句话**：一种自平衡二叉搜索树，通过颜色规则和旋转保持大致平衡。

### 2. 五条性质

1. 节点是红色或黑色
2. 根节点是黑色
3. 所有叶子（NIL）是黑色
4. 红色节点的子节点必须是黑色（不能连续红）
5. 从任一节点到其每个叶子的所有路径都包含相同数目的黑色节点

### 3. 红黑树 vs AVL

| 特性 | AVL | 红黑树 |
|---|---|---|
| 平衡度 | 更严格 | 较宽松 |
| 查找 | 更快 | 稍慢 |
| 插入/删除 | 旋转更多 | 旋转更少 |
| 应用 | 数据库索引 | Java HashMap、C++ map |

### 4. 应用

- Java `TreeMap`、`TreeSet`
- Linux 内核 CFS 调度器
- C++ `std::map`、`std::set`

---

## 五、B 树与 B+树

### 1. B 树

**一句话**：多路平衡搜索树，每个节点可以有多个子节点，适合磁盘存储。

```
        [20 | 50 | 80]
       /    |    |    \
  [5,10] [30,40] [60,70] [90,100]
```

**特点**：
- 每个节点存多个 key
- 所有叶子节点在同一层
- 常用于文件系统和数据库索引

### 2. B+树

```
        [20 | 50 | 80]
       /    |    |    \
  [5,10] [30,40] [60,70] [90,100]
      \      |       |        /
       [5,10,20,30,40,50,60,70,80,90,100]  ← 叶子节点链表
```

**特点**：
- 非叶子节点只存索引 key，不存数据
- 所有数据都在叶子节点
- 叶子节点之间用指针连接，方便范围查询

### 3. B 树 vs B+树

| 特性 | B 树 | B+树 |
|---|---|---|
| 数据存储 | 所有节点都存 | 只有叶子存 |
| 范围查询 | 需要中序遍历 | 叶子链表直接遍历 |
| IO 次数 | 更多 | 更少 |
| 应用 | 文件系统 | 数据库索引（MySQL InnoDB） |

---

## 六、Trie 树（前缀树）

### 1. 定义

**一句话**：专门用于处理字符串的多叉树，路径上的字符组成字符串。

```
                    root
                   /    \
                  a      b
                 /        \
                p          e
               / \          \
              p   r          e
             /     \          \
            l*      e*         r*
            |
            e*

* 表示单词结尾
存储：apple, app, are, beer
```

### 2. 应用

- 搜索引擎自动补全
- 拼写检查
- IP 路由最长前缀匹配
- 敏感词过滤

### 3. Go 实现

```go
type TrieNode struct {
    children map[rune]*TrieNode
    isEnd    bool
}

type Trie struct {
    root *TrieNode
}

func NewTrie() *Trie {
    return &Trie{root: &TrieNode{children: make(map[rune]*TrieNode)}}
}

func (t *Trie) Insert(word string) {
    node := t.root
    for _, ch := range word {
        if node.children[ch] == nil {
            node.children[ch] = &TrieNode{children: make(map[rune]*TrieNode)}
        }
        node = node.children[ch]
    }
    node.isEnd = true
}

func (t *Trie) Search(word string) bool {
    node := t.root
    for _, ch := range word {
        if node.children[ch] == nil {
            return false
        }
        node = node.children[ch]
    }
    return node.isEnd
}
```

---

## 七、堆（Heap）

### 1. 定义

**一句话**：完全二叉树，父节点值大于（大顶堆）或小于（小顶堆）子节点。

```
大顶堆：              小顶堆：
       9                    1
      / \                  / \
     7   6                2   3
    / \ / \              / \ / \
   5  4 3  2            4  5 6  7
```

### 2. 应用

- 优先队列
- Top K 问题
- 堆排序
- 定时任务调度

### 3. Top K 问题

```go
// 找数组中最大的 K 个数，用小顶堆
func topK(nums []int, k int) []int {
    h := &IntHeap{}
    heap.Init(h)
    for _, n := range nums {
        heap.Push(h, n)
        if h.Len() > k {
            heap.Pop(h)
        }
    }
    return *h
}

type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] } // 小顶堆
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *IntHeap) Push(x interface{}) { *h = append(*h, x.(int)) }
func (h *IntHeap) Pop() interface{} {
    old := *h
    n := len(old)
    x := old[n-1]
    *h = old[:n-1]
    return x
}
```

---

## 八、常见面试题

### Q1：二叉树的前中后序遍历，递归和非递归怎么写？

**答**：
- 递归：按顺序访问根、左、右
- 非递归：用栈模拟递归过程
- 层序：用队列 BFS

### Q2：BST 中序遍历有什么特点？

**答**：BST 中序遍历结果是一个**递增有序数组**。

### Q3：AVL 和红黑树有什么区别？

**答**：
- AVL 更严格平衡，查找更快，但插入删除旋转更多
- 红黑树较宽松，插入删除旋转少，综合性能更优
- 红黑树在实际库中更常见

### Q4：为什么数据库索引用 B+树而不是红黑树？

**答**：
- B+树是多路树，树高更低，磁盘 IO 更少
- B+树叶子节点链表，范围查询极快
- 红黑树是二叉树，节点多、树高高，不适合磁盘场景

### Q5：Trie 树和哈希表的区别？

**答**：
- Trie 适合前缀匹配、自动补全
- 哈希表适合精确查找
- Trie 空间换时间，前缀查询效率高

### Q6：堆和栈的区别？

| 特性 | 堆（数据结构） | 栈 |
|---|---|---|
| 存取 | 按优先级 | 先进后出 |
| 底层 | 完全二叉树 | 数组/链表 |
| 应用 | Top K、优先队列、排序 | 函数调用、表达式求值 |

### Q7：怎么判断一棵二叉树是否平衡？

```go
func isBalanced(root *TreeNode) bool {
    _, ok := height(root)
    return ok
}

func height(root *TreeNode) (int, bool) {
    if root == nil {
        return 0, true
    }
    lh, lok := height(root.Left)
    rh, rok := height(root.Right)
    if !lok || !rok || abs(lh-rh) > 1 {
        return 0, false
    }
    return max(lh, rh) + 1, true
}
```

### Q8：二叉树最大深度和最小深度？

```go
// 最大深度
func maxDepth(root *TreeNode) int {
    if root == nil {
        return 0
    }
    return max(maxDepth(root.Left), maxDepth(root.Right)) + 1
}

// 最小深度（到最近叶子）
func minDepth(root *TreeNode) int {
    if root == nil {
        return 0
    }
    if root.Left == nil {
        return minDepth(root.Right) + 1
    }
    if root.Right == nil {
        return minDepth(root.Left) + 1
    }
    return min(minDepth(root.Left), minDepth(root.Right)) + 1
}
```

---

## 九、一句话总结

- **二叉树**：基础，掌握前中后序和层序遍历
- **BST**：左 < 根 < 右，中序有序
- **AVL**：严格平衡，旋转维护
- **红黑树**：宽松平衡，插入删除效率高
- **B/B+树**：多路树，数据库和文件系统索引
- **Trie**：前缀树，自动补全、敏感词
- **堆**：优先队列、Top K、堆排序

> **面试口诀：BST 中序变有序，AVL 红黑保平衡，B+树做数据库索引，Trie 做前缀匹配，堆做 Top K**
