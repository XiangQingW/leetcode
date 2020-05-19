package main

import "fmt"

func letterCasePermutation(S string) []string {
	r := []string{}

	if len(S) == 0 {
		return r
	}

	cur_c := S[0]
	var other_c byte
	if 'a' <= cur_c && cur_c <= 'z' {
		other_c = cur_c - 'a' + 'A'
	} else if 'A' <= cur_c && cur_c <= 'Z' {
		other_c = cur_c + 'a' - 'A'
	}

	remains := letterCasePermutation(S[1:])

	if len(remains) == 0 {
		if other_c != 0 {

			return []string{string(cur_c), string(other_c)}
		} else {

			return []string{string(cur_c)}
		}
	}

	for _, remain := range remains {
		strs := string(cur_c) + remain
		r = append(r, strs)

		if other_c != 0 {
			strs = string(other_c) + remain
			r = append(r, strs)
		}
	}
	return r
}

func main() {
	r := letterCasePermutation("ab")
	fmt.Println(r)
}
