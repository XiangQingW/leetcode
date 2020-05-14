package main

import (
	"fmt"
	"strconv"
	"strings"
)

func queryString(S string, N int) bool {
	isSub := true
	for i := N; i >= 1; i-- {
		binary := strconv.FormatInt(int64(i), 2)
		isSub = isSub && strings.Contains(S, binary)
		if !isSub {
			return isSub
		}
	}

	return isSub
}

func main() {
	r := queryString("0110", 3)
	fmt.Println(r)
}
