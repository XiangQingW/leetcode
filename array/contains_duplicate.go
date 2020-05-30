package main

func containsDuplicate(nums []int) bool {
	num2exist := make(map[int]bool)

	for _, num := range nums {
		if _, ok := num2exist[num]; ok {
			return true
		}
		num2exist[num] = true
	}

	return false
}
