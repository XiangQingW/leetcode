package main

import "fmt"

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	raw := x
	back := 0

	for x > 0 {
		last_num := x % 10
		x /= 10
		back = back*10 + last_num
	}

	return back == raw
}

func main() {
	fmt.Println(isPalindrome(121))
	fmt.Println(isPalindrome(12))
	fmt.Println(isPalindrome(-121))
	fmt.Println(isPalindrome(10))
	fmt.Println(isPalindrome(1221))
	fmt.Println(isPalindrome(2222))
	fmt.Println(isPalindrome(2224))
	fmt.Println(isPalindrome(0))
	fmt.Println(isPalindrome(5))
}
