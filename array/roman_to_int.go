package main

import "fmt"

func romanToInt(s string) int {
	if len(s) <= 0 {
		return 0
	}

	char2num := make(map[byte]int)
	char2num['I'] = 1
	char2num['V'] = 5
	char2num['X'] = 10
	char2num['L'] = 50
	char2num['C'] = 100
	char2num['D'] = 500
	char2num['M'] = 1000

	sum := 0
	pre := 0
	for i := len(s) - 1; 0 <= i; i-- {
		c := s[i]
		if num, ok := char2num[c]; ok {
			if num < pre {
				sum -= num
			} else {
				sum += num
			}
			pre = num
		}
	}

	return sum

}

func main() {
	fmt.Println(romanToInt("III"))
	fmt.Println(romanToInt("IV"))
	fmt.Println(romanToInt("IX"))
	fmt.Println(romanToInt("LVIII"))
	fmt.Println(romanToInt("MCMXCIV"))
}
