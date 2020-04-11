package main

import (
	"container/list"
	"fmt"
)

type Node struct {
	Val      int
	Children []*Node
}

func maxDepth(root *Node) int {
	if root == nil {
		return 0
	}

	cur_level := []*Node{root}
	next_level := []*Node{}

	level := 0
	for 0 < len(cur_level) {
		level += 1
		for _, node := range cur_level {
			if node == nil {
				continue
			}
			next_level = append(next_level, node.Children...)
		}
		cur_level = next_level
		next_level = nil
	}

	return level
}

func main() {
}
