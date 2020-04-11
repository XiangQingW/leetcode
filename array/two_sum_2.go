package main

import (
	"fmt"
)

func twoSum(numbers []int, target int) []int {
	if len(numbers) < 2 {
		return nil
	}

	start := 0
	end := len(numbers) - 1

	for start < end {
		sum := numbers[start] + numbers[end]
		if sum == target {
			return []int{start + 1, end + 1}
		}

		if sum < target {
			start += 1
		} else {
			end -= 1
		}
	}

	return nil

}

func main() {
	r := twoSum([]int{2, 7, 11, 15}, 9)
	fmt.Println(r)
}
