package main

import (
	"container/list"
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sortedArrayToBST(nums []int) *TreeNode {
	l := len(nums)
	if nums == nil || l == 0 {
		return nil
	}

	var root, left, right *TreeNode

	mid := l / 2

	left = sortedArrayToBST(nums[:mid])

	if mid+1 < l {
		right = sortedArrayToBST(nums[mid+1:])
	} else {
		right = nil
	}

	root = new(TreeNode)
	root.Val = nums[mid]
	root.Left = left
	root.Right = right
	return root
}

func main() {
}
