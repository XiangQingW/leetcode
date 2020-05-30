package main

import "sort"

func arrayPairSum(nums []int) int {
	sort.Ints(nums)
	sum := 0

	for idx, num := range nums {
		if idx%2 == 0 {
			sum += num
		}
	}
	return sum
}
