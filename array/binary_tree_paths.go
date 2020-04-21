package main

import (
	"fmt"
	"strconv"
	"strings"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func deepSearch(node *TreeNode, visited_node *TreeNode, st *[]*TreeNode) {
	if node == nil {
		return
	}

	if node.Right == visited_node {
		return
	}

	for node != nil {
		if node.Left != visited_node {
			node = node.Left
		} else {
			node = node.Right
		}
		if node != nil {
			*st = append(*st, node)
		}
	}
}

func binaryTreePaths(root *TreeNode) []string {
	if root == nil {
		return []string{}
	}

	r := []string{}
	st := []*TreeNode{root}

	deepSearch(root, nil, &st)
	for len(st) != 0 {
		last_node := st[len(st)-1]
		if last_node.Left == nil && last_node.Right == nil {
			path := []string{}
			for _, node := range st {
				path = append(path, strconv.Itoa(node.Val))
			}
			r = append(r, strings.Join(path, "->"))
		}
		st = st[:len(st)-1]
		node := st[len(st)-1]
		deepSearch(node, last_node, &st)
	}

	return r
}

func main() {
	fmt.Println("hello")
}
