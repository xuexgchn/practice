package main

import (
	"fmt"
)

func main() {
	fmt.Println(test())
}

func test() bool {
	if reverseVowels("hello") != "holle" {
		return false
	}
	if reverseVowels("leetcode") != "leotcede" {
		return false
	}

	return true
}

func reverseVowels(s string) string {
	newString := []byte(s)
	for i, j := 0, len(newString)-1; i < j; {
		for ; i < j && !isVowels(newString[i]); i++ {

		}
		for ; i < j && !isVowels(newString[j]); j-- {

		}
		if i < j {
			newString[i], newString[j] = newString[j], newString[i]
			i++
			j--
		}
	}

	return string(newString)
}

func isVowels(s byte) bool {
	return s == 'a' || s == 'e' || s == 'i' || s == 'o' || s == 'u' || s == 'A' || s == 'E' || s == 'I' || s == 'O' || s == 'U'
}
