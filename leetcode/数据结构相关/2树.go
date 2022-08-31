package 数据结构相关

// 树  0/31

//Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//104. 二叉树的最大深度 https://leetcode.cn/problems/maximum-depth-of-binary-tree/description/
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return max(maxDepth(root.Left), maxDepth(root.Right)) + 1
}
func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

//
