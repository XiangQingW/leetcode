package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func splitListToParts(root *ListNode, k int) []*ListNode {
	r := []*ListNode{}
	if root == nil || k == 0 {
		return r
	}

	if k == 1 {
		return append(r, root)
	}

	node := root
	l := 0
	for node != nil {
		l += 1
		node = node.Next
	}

	part_len := l / k
	left_nums := l - part_len*k

	for i := 0; i < k; i++ {
		cur_len := part_len
		if 0 < left_nums {
			cur_len += 1
			left_nums -= 1
		}

		head := root
		cur_len -= 1
		for cur_len >= 0 && root != nil {
			if cur_len == 0 {
				tmp := root.Next
				root.Next = nil
				root = tmp
				break
			}
			root = root.Next
			cur_len -= 1
		}

		r = append(r, head)
	}
	return r

}

func main() {

}
