package main

import "strings"

func isPrefixOfWord(sentence string, searchWord string) int {
	if strings.HasPrefix(sentence, searchWord) {
		return 1
	}

	searchWord = string(' ') + searchWord
	i := strings.Index(sentence, searchWord)

	if i == -1 {
		return i
	}

	subStr := sentence[:i]
	count := 0
	for _, c := range subStr {
		if c == ' ' {
			count++
		}
	}

	return count + 2
}

func main() {}
