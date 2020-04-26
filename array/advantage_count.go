package main

import (
	"fmt"
	"sort"
)

func getIdxByValue(v2idx *map[int][]int, v int) int {
	if entry, ok := (*v2idx)[v]; ok {
		fmt.Println(entry, v)
		idx := entry[len(entry)-1]
		entry = entry[:len(entry)-1]
		(*v2idx)[v] = entry
		return idx
	}
	return 0
}

func advantageCount(A []int, B []int) []int {
	b_v2idx := make(map[int][]int)

	for i := 0; i < len(A); i++ {
		if entry, ok := b_v2idx[B[i]]; ok {
			entry = append(entry, i)
			b_v2idx[B[i]] = entry
		} else {
			b_v2idx[B[i]] = []int{i}
		}
	}
	sort.Ints(A)
	sort.Ints(B)

	r := make([]int, len(A))
	last_pos := len(A) - 1
	b_idx := 0

	for _, a := range A {
		b := B[b_idx]
		if a <= b {
			last_b := B[last_pos]

			idx := getIdxByValue(&b_v2idx, last_b)
			fmt.Println(b_v2idx)
			r[idx] = a

			last_pos--
			continue
		}

		idx := getIdxByValue(&b_v2idx, b)
		fmt.Println(b_v2idx)
		r[idx] = a

		b_idx++
	}

	return r
}

func main() {
	r := advantageCount([]int{2, 0, 4, 1, 2}, []int{1, 3, 0, 0, 2})
	fmt.Println(r)
}
