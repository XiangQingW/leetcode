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

func averageOfLevels(root *TreeNode) []float64 {
	if root == nil {
		return []float64{}
	}

	r := []float64{}
	cur_level_len := 1

	cur_level_node := list.New()
	cur_level_node.PushBack(root)

	for 0 < cur_level_len {
		subed_len := cur_level_len
		sum := 0
		next_level_len := 0

		for 0 < subed_len {
			head := cur_level_node.Front()
			node := head.Value.(*TreeNode)
			sum += node.Val

			if node.Left != nil {
				cur_level_node.PushBack(node.Left)
				next_level_len += 1
			}

			if node.Right != nil {
				cur_level_node.PushBack(node.Right)
				next_level_len += 1
			}

			cur_level_node.Remove(head)
			subed_len -= 1
		}

		r = append(r, float64(sum)/float64(cur_level_len))
		cur_level_len = next_level_len

		fmt.Println(r)
	}

	return r
}

func main() {
	root := new(TreeNode)
	root.Val = 3

	node1 := new(TreeNode)
	node1.Val = 9

	node2 := new(TreeNode)
	node2.Val = 20

	node3 := new(TreeNode)
	node3.Val = 15

	node4 := new(TreeNode)
	node4.Val = 15

	root.Left = node1
	root.Right = node2

	node2.Left = node3
	node2.Right = node4

	averageOfLevels(root)

}
