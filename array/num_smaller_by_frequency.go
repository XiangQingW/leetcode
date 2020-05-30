package main

import "sort"

func getF(s string) int {
	min := byte('z')
	for i := 0; i < len(s); i++ {
		if s[i] < min {
			min = s[i]
		}
	}
	r := 0
	for i := 0; i < len(s); i++ {
		if s[i] == min {
			r++
		}
	}
	return r
}

func numSmallerByFrequency(queries []string, words []string) []int {
	count := []int{}
	for _, word := range words {
		f := getF(word)
		count = append(count, f)
	}
	sort.Ints(count)

	r := []int{}
	for _, q := range queries {
		f := getF(q)

		l := sort.Search(len(count), func(i int) bool { return count[i] > f })
		r = append(r, len(count)-l)
	}

	return r
}

func main() {

}
