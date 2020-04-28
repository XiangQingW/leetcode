package main

import (
	"fmt"
	"math"
)

type Pos struct {
	row int
	col int
}

func getRepeatedBytes(count int, b byte) []byte {
	r := []byte{}
	for i := 0; i < count; i++ {
		r = append(r, b)
	}
	return r
}

func alphabetBoardPath(target string) string {
	mp := make(map[byte]Pos)

	count := 0
	row := 0
	for i := byte('a'); i <= byte('z'); i++ {
		mp[i] = Pos{
			row,
			count}

		count++
		if count == 5 {
			row++
			count = 0
		}
	}

	ans := []byte{}
	cur_row := 0
	cur_col := 0

	for i := 0; i < len(target); i++ {
		next_char := byte(target[i])

		if pos, ok := mp[next_char]; ok {
			dif_row := cur_row - pos.row
			dif_col := cur_col - pos.col

			bytes := []byte{}
			shouldAppendDown := next_char == 'z' && dif_row != 0
			if dif_row < 0 {
				count := math.Abs(float64(dif_row))
				if shouldAppendDown {
					count--
				}
				bytes = getRepeatedBytes(int(count), 'D')
			} else if 0 < dif_row {
				bytes = getRepeatedBytes(dif_row, 'U')
			}
			ans = append(ans, bytes...)

			bytes = []byte{}
			if dif_col < 0 {
				count := math.Abs(float64(dif_col))
				bytes = getRepeatedBytes(int(count), 'R')
			} else if 0 < dif_col {
				bytes = getRepeatedBytes(dif_col, 'L')
			}
			ans = append(ans, bytes...)

			cur_row = pos.row
			cur_col = pos.col

			if shouldAppendDown {
				ans = append(ans, 'D')
			}
			ans = append(ans, '!')
		}
	}

	return string(ans)
}

func main() {
	r := alphabetBoardPath("zdz")
	fmt.Println(r)
}
