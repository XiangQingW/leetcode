package main

import (
	"fmt"
)

func removeElement(nums []int, val int) int {
	if nums == nil {
		return 0
	}

	if len(nums) == 0 {
		return 0
	}

	replaced_pos := 0
	reverse_pos := len(nums) - 1

	for i := 0; i < len(nums); i++ {
		if nums[replaced_pos] != val {
			replaced_pos += 1
			continue
		}

		nums[replaced_pos], nums[reverse_pos] = nums[reverse_pos], nums[replaced_pos]
		reverse_pos -= 1
	}

	l := replaced_pos
	if replaced_pos < len(nums) && nums[replaced_pos] != val {
		l += 1
	}

	return l
}

func main() {
	nums := []int{3, 2, 2, 3}
	r := removeElement(nums, 3)
	fmt.Println(r)
	fmt.Println(nums)

	nums = []int{0, 1, 2, 2, 3, 0, 4, 2}
	r = removeElement(nums, 2)
	fmt.Println(r)
	fmt.Println(nums)
}
