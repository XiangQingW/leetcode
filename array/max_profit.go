package main

import "math"

func maxProfit(prices []int) int {
	if len(prices) <= 1 {
		return 0
	}

	res := 0
	cur_min := prices[0]

	for i := 1; i < len(prices); i++ {
		res = int(math.Max(float64(res), float64(prices[i]-cur_min)))
		cur_min = int(math.Min(float64(cur_min), float64(prices[i])))
	}

	return res
}

func main() {

}
