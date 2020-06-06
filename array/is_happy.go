package main

import "math"

func helper(n int) int {
	r := 0
	for n != 0 {
		left := n % 10
		n /= 10

		r += int(math.Pow(float64(left), 2.0))
	}

	return r
}

func isHappy(n int) bool {
	num2exist := make(map[int]bool)

	for true {
		if n == 1 {
			return true
		}
		n = helper(n)

		if _, ok := num2exist[n]; ok {
			return false
		}
		num2exist[n] = true
	}

	return false
}

func main() {

}
