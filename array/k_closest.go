package main

import (
	"container/heap"
	"math"
)

type kHeap []node

type node struct {
	distance int
	point    []int
}

func (h kHeap) Len() int { return len(h) }
func (h kHeap) Less(i, j int) bool {
	return h[i].distance < h[j].distance
}
func (h kHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *kHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func (h *kHeap) Push(x interface{}) {
	*h = append(*h, x.(node))
}

func kClosest(points [][]int, K int) [][]int {
	h := &kHeap{}
	heap.Init(h)

	for _, point := range points {
		distance := math.Pow(float64(point[0]), 2.0) + math.Pow(float64(point[1]), 2.0)
		n := node{
			int(distance),
			point,
		}
		h.Push(n)

		if K < len(*h) {
			h.Pop()
		}
	}

	r := [][]int{}
	for _, n := range *h {
		r = append(r, n.point)
	}
	return r
}

func main() {

}
