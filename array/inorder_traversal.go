package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func helper(node *TreeNode, res *[]int) {
	if node == nil {
		return
	}
	helper(node.Left, res)
	*res = append(*res, node.Val)
	helper(node.Right, res)
}

func inorderTraversal(root *TreeNode) []int {
	r := []int{}
	helper(root, &r)
	return r
}
