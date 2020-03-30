package main

import "fmt"

func searchInsert(nums []int, target int) int {
	if nums == nil {
		return 0
	}

	for idx, v := range nums {
		if target <= v {
			return idx
		}
	}

	return len(nums)
}

func main() {
	fmt.Println(searchInsert([]int{1, 3, 5, 6}, 0))
}
