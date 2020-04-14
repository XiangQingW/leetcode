package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type BSTIterator struct {
	Stack []*TreeNode
}

func Constructor(root *TreeNode) BSTIterator {
	Stack := []*TreeNode{}

	for root != nil {
		Stack = append(Stack, root)
		root = root.Left
	}

	return BSTIterator{
		Stack: Stack,
	}
}

/** @return the next smallest number */
func (this *BSTIterator) Next() int {
	l := len(this.Stack)
	last_node := this.Stack[l-1]

	this.Stack = this.Stack[:l-1]

	next_node := last_node.Right
	for next_node != nil {
		this.Stack = append(this.Stack, next_node)
		next_node = next_node.Left
	}

	return last_node.Val
}

func (this *BSTIterator) HasNext() bool {
	return len(this.Stack) != 0
}

func main() {
	// obj := Constructor(root)
	// param_1 := obj.Next()
	// param_2 := obj.HasNext()

}
