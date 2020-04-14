package main

import (
	"container/heap"
	"fmt"
)

type Heap []int

func (this Heap) Len() int {
	return len(this)
}

func (this Heap) Less(i, j int) bool {
	return this[i] < this[j]
}

func (this Heap) Swap(i, j int) {
	this[i], this[j] = this[j], this[i]
}

func (this *Heap) Push(x interface{}) {
	*this = append(*this, x.(int))
}

func (this *Heap) Pop() interface{} {
	l := len(*this)
	last_v := (*this)[l-1]

	*this = (*this)[0 : l-1]
	return last_v
}

func findKthLargest(nums []int, k int) int {
	h := new(Heap)
	heap.Init(h)

	for _, num := range nums {
		heap.Push(h, num)
		if k < len(*h) {
			heap.Pop(h)
		}
	}

	return (*h)[0]
}

func main() {
	r := findKthLargest([]int{3, 2, 1, 5, 6, 4}, 2)
	fmt.Println(r)

	r2 := findKthLargest([]int{3, 2, 3, 1, 2, 4, 5, 5, 6}, 4)
	fmt.Println(r2)
}
