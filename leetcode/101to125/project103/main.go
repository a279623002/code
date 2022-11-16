package project100

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 给你二叉树的根节点 root ，返回其节点值的 锯齿形层序遍历 。（即先从左往右，再从右往左进行下一层遍历，以此类推，层与层之间交替进行）。
// 输入：root = [3,9,20,null,null,15,7]
// 输出：[[3],[20,9],[15,7]]
func zigzagLevelOrder(root *TreeNode) [][]int {
	res := [][]int{}
	if root == nil {
		return res
	}
	queue := []*TreeNode{root}
	for i := 0; len(queue) > 0; i++ {
		curr := []int{}
		q := queue
		queue = nil
		for _, v := range q {
			curr = append(curr, v.Val)
			if v.Left != nil {
				queue = append(queue, v.Left)
			}
			if v.Right != nil {
				queue = append(queue, v.Right)
			}
		}

		if i%2 == 1 {
			for n, m := 0, len(curr); n < m/2; n++ {
				curr[n], curr[m-1-n] = curr[m-1-n], curr[n]
			}
		}
		res = append(res, curr)
	}
	return res
}
