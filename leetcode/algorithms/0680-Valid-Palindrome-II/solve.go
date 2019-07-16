package main

import (
	"fmt"
)

func main() {
	fmt.Println(test())
}

func test() bool {
	if validPalindrome("aba") != true {
		return false
	}
	if validPalindrome("abca") != true {
		return false
	}

	return true
}

func validPalindrome(s string) bool {
	left, right := 0, len(s)-1

	for left < right {
		if s[left] != s[right] {
			return isPalind(s, left, right-1) || isPalind(s, left+1, right)
		}
		left++
		right--
	}
	return true
}

func isPalind(s string, left int, right int) bool {
	for left < right {
		if s[left] != s[right] {
			return false
		}
		left++
		right--
	}
	return true
}
