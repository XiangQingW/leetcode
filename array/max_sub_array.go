package main

import "math"

func maxSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	ans := nums[0]
	unit_cur_max := nums[0]

	for _, num := range nums {
		if 0 < unit_cur_max {
			unit_cur_max += num
		} else {
			unit_cur_max = num
		}

		ans = int(math.Max(float64(ans), float64(unit_cur_max)))
	}

	return ans
}

func main() {

}
