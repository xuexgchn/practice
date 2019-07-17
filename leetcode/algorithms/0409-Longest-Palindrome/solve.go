package main

import (
	"fmt"
)

func main() {
	fmt.Println(test())
}

func test() bool {
	if longestPalindrome("abccccdd") != 7 {
		return false
	}
	return true
}

func longestPalindrome(s string) int {
	length := len(s)

	if length < 2 {
		return length
	}

	count := make([]int, 300)
	flag := false

	for i := 0; i < length; i++ {
		count[s[i]]++
	}

	ans := 0
	for i := 0; i < len(count); i++ {
		if count[i]%2 == 1 {
			flag = true
		}
		ans += count[i] / 2 * 2
	}
	if flag {
		return ans + 1
	}
	return ans
}
