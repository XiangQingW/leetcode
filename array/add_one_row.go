package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func addOneRow(root *TreeNode, v int, d int) *TreeNode {
	if root == nil || d == 1 {
		return &TreeNode{
			Val:   v,
			Left:  root,
			Right: nil,
		}
	}

	pre_level_node := []*TreeNode{root}
	cur_level_node := []*TreeNode{}

	for 2 < d {
		for _, node := range pre_level_node {
			if node.Left != nil {
				cur_level_node = append(cur_level_node, node)
			}
			if node.Right != nil {
				cur_level_node = append(cur_level_node, node)
			}
		}
		d -= 1
		pre_level_node = cur_level_node
		cur_level_node = cur_level_node[:0]
	}

	for _, node := range pre_level_node {
		old_left := node.Left
		old_right := node.Right

		new_left := TreeNode{
			Val:   v,
			Left:  old_left,
			Right: nil,
		}
		node.Left = &new_left

		new_right := TreeNode{
			Val:   v,
			Left:  nil,
			Right: old_right,
		}
		node.Right = &new_right
	}
	return root
}

func main() {

	leaf1 := TreeNode{
		Val:   3,
		Left:  nil,
		Right: nil,
	}
	leaf2 := TreeNode{
		Val:   1,
		Left:  nil,
		Right: nil,
	}
	node := TreeNode{
		Val:   2,
		Left:  &leaf1,
		Right: &leaf2,
	}
	root := TreeNode{
		Val:   4,
		Left:  &node,
		Right: nil,
	}

	addOneRow(&root, 1, 3)
}
