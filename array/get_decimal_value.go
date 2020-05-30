package main

import "math"

type ListNode struct {
	Val  int
	Next *ListNode
}

func getDecimalValue(head *ListNode) int {
	l := 0
	node := head
	for node != nil {
		l++
		node = node.Next
	}

	r := 0
	node = head
	for node != nil {
		l--
		if node.Val != 0 {
			r += int(math.Pow(float64(2), float64(l)))
		}
		node = node.Next
	}
	return r
}

func main() {

}
