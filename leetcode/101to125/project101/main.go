package project100

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 给你一个二叉树的根节点 root ， 检查它是否轴对称
// 输入：root = [1,2,2,3,4,4,3]
// 输出：true
// 输入：root = [1,2,2,null,3,null,3]
// 输出：false
func isSymmetric(root *TreeNode) bool {
	return check(root, root)
}

func check(l, r *TreeNode) bool {
	if l == nil {
		return r == nil
	}
	if r == nil {
		return l == nil
	}
	return l.Val == r.Val && check(l.Left, r.Right) && check(l.Right, r.Left)
}
