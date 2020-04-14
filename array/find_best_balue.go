package main

import (
	"fmt"
	"math"
	"sort"
)

func findBestValue(arr []int, target int) int {
	sort.Ints(arr)
	l := len(arr)

	sum := 0
	for index, num := range arr {
		left_nums := l - index
		fmt.Println(float64(target-sum) / float64(left_nums))
		avg := int(math.Floor(float64(target-sum)/float64(left_nums) + 0.4))
		if avg <= num {
			return avg
		}
		sum += num
	}

	return arr[l-1]
}

func main() {
	r := findBestValue([]int{4, 3, 9}, 10)
	fmt.Println(r)

	r = findBestValue([]int{2, 3, 5}, 10)
	fmt.Println(r)

	r = findBestValue([]int{60864, 25176, 27249, 21296, 20204}, 56803)
	fmt.Println(r)

	r = findBestValue([]int{1547, 83230, 57084, 93444, 70879}, 71237)
	fmt.Println(r)
}
