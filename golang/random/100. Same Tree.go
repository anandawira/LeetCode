package random

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSameTree(p *TreeNode, q *TreeNode) bool {
	// If both is nil
	if p == nil && q == nil {
		return true
	} else if p == nil || q == nil {
		return false
	}

	if p.Val != q.Val {
		return false
	}

	isLeftTreeSame := isSameTree(p.Left, q.Left)
	isRightTreeSame := isSameTree(p.Right, q.Right)

	return isLeftTreeSame && isRightTreeSame
}
