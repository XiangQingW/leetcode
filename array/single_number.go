package main

import (
	"fmt"
)

func singleNumber(nums []int) int {
	if nums == nil {
		return 0
	}

	if len(nums) == 0 {
		return 0
	}

	s := make(map[int]bool)

	for _, v := range nums {
		if _, ok := s[v]; ok {
			delete(s, v)
		} else {
			s[v] = true
		}
	}

	for k, _ := range s {
		return k
	}
	return 0

}

func main() {
	nums := []int{4, 1, 2, 1, 2}
	r := singleNumber(nums)
	fmt.Println(r)
}
