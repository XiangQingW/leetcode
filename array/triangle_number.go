package main

import (
	"fmt"
	"sort"
)

func binarySearchMatchedLen(nums []int, left_len int) int {
	if len(nums) == 0 {
		return 0
	}

	mid := len(nums) / 2
	for 0 <= mid && mid < len(nums) {
		if mid == 0 {
			if left_len < nums[mid] {
				mid += 1
			} else {
				mid -= 1
			}
			continue
		}

		if left_len < nums[mid-1] && nums[mid] <= left_len {
			break
		}
		fmt.Println(nums)
		fmt.Println(left_len, mid)
		if nums[mid] <= left_len {
			mid = mid / 2
		} else {
			old := mid
			mid = (mid + len(nums) - 1) / 2
			if mid == old {
				mid += 1
			}
		}
	}

	if mid == -1 {
		return 0
	} else if mid == len(nums) {
		return len(nums)
	} else {
		return mid
	}
}

func triangleNumber(nums []int) int {
	if nums == nil {
		return 0
	}

	l := len(nums)
	if l < 3 {
		return 0
	}

	sort.Sort(sort.Reverse(sort.IntSlice(nums)))

	count := 0
	for i := 1; i < l; i++ {
		max_len := nums[i-1]

		for j := i; j < l; j++ {
			second_len := nums[j]

			if j == l-1 {
				continue
			}

			left_len := max_len - second_len
			count += binarySearchMatchedLen(nums[j+1:], left_len)
		}
	}

	return count
}

func main() {
	fmt.Println(triangleNumber([]int{58, 97, 2, 8, 2, 55, 14, 54, 2}))
}
