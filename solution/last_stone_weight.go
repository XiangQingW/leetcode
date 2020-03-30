package main

import (
	"container/heap"
	"fmt"
)

type IntHeap []int

func (h IntHeap) Len() int { return len(h) }
func (h IntHeap) Less(i, j int) bool {
	return h[i] > h[j]
}
func (h IntHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func lastStoneWeight(stones []int) int {
	h := IntHeap(stones)
	heap.Init(&h)

	for 1 < len(h) {
		max1 := heap.Pop(&h)
		max2 := heap.Pop(&h)

		if max1 == max2 {
			continue
		}

		remain := max1.(int) - max2.(int)
		heap.Push(&h, remain)
	}

	if len(h) == 0 {
		return 0
	}

	return h[0]
}

func main() {

}
