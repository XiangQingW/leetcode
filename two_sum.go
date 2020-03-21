package main

import (
	"fmt"
)

func twoSum(nums []int, target int) []int {
	var2index := make(map[int]int)

	for index, value := range nums {
		var2index[value] = index
	}

	for index, value := range nums {
		left := target - value
		left_idx, ok := var2index[left]

		if left_idx == index {
			continue
		}

		if ok {
			fmt.Println("hello")
			return []int{index, left_idx}
		}
	}
	return []int{0, 0}
}

func main() {
	r := twoSum([]int{3, 2, 4}, 6)
	fmt.Println(r)
}
