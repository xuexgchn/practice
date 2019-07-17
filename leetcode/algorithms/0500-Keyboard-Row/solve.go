package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(test())
}

func test() bool {
	findWords([]string{"Hello", "Alaska", "Dad", "Peace"})
	return true
}

var count = map[rune]byte{
	'a': 'a',
	'b': 'z',
	'c': 'z',
	'd': 'a',
	'e': 'q',
	'f': 'a',
	'g': 'a',
	'h': 'a',
	'i': 'q',
	'j': 'a',
	'k': 'a',
	'l': 'a',
	'm': 'z',
	'n': 'z',
	'o': 'q',
	'p': 'q',
	'q': 'q',
	'r': 'q',
	's': 'a',
	't': 'q',
	'u': 'q',
	'v': 'z',
	'w': 'q',
	'x': 'z',
	'y': 'q',
	'z': 'z',
}

func findWords(words []string) []string {
	ans := make([]string, 0, len(words))

	for _, word := range words {
		newWord := strings.ToLower(word)
		if isAllIn(newWord) {
			ans = append(ans, word)
		}
	}

	// fmt.Println(ans)
	return ans
}

func isAllIn(s string) bool {
	for _, c := range s {
		if count[c] != count[rune(s[0])] {
			return false
		}
	}

	return true
}
