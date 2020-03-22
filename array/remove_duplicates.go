package main

import "fmt"

func removeDuplicates(nums []int) int {
	if nums == nil {
		return 0
	}

	if len(nums) == 0 {
		return 0
	}

	r := 1
	pre_v := nums[0]

	for idx, v := range nums {
		if idx == 0 {
			continue
		}

		if v == pre_v {
			continue
		}

		pre_v = v
		nums[r] = v
		r += 1
	}

	return r
}

func main() {
	a := []int{1, 1, 2}
	r := removeDuplicates(a)
	fmt.Println(r)
	fmt.Println(a)
}
