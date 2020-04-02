package main

import (
	"fmt"
)

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	ms := make(map[rune]int)
	mt := make(map[rune]int)

	for _, c := range s {
		ms[c]++
	}

	for _, c := range t {
		mt[c]++
	}

	if len(ms) != len(mt) {
		return false
	}

	for c1, n1 := range ms {
		n2, ok := mt[c1]
		if !ok {
			return false
		}

		if n1 != n2 {
			return false
		}
	}

	return true
}

func main() {
	fmt.Println(isAnagram("", ""))
	fmt.Println(isAnagram("anagram", "nagaram"))
	fmt.Println(isAnagram("rat", "car"))
}
