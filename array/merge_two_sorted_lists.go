package main

import "fmt"

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	var head, cur *ListNode

	for l1 != nil || l2 != nil {
		var min *ListNode
		if l1 == nil {
			min = l2
			l2 = l2.Next
		} else if l2 == nil {
			min = l1
			l1 = l1.Next
		} else {
			if l1.Val <= l2.Val {
				min = l1
				l1 = l1.Next
			} else {
				min = l2
				l2 = l2.Next
			}
		}

		if head == nil {
			head = min
			cur = head
		} else {
			cur.Next = min
			cur = cur.Next
		}
	}

	return head
}

func main() {
	fmt.Println("hello")
}
