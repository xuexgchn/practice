package main

import (
	"fmt"
)

func main() {
	fmt.Println(test())
}

func test() bool {
	if findTheDifference("abcd", "abcde") != 'e' {
		return false
	}

	return true
}

func findTheDifference(s string, t string) byte {
	countS := make([]int, 30)
	countT := make([]int, 30)

	for _, char := range s {
		countS[int(char-'a')]++
	}
	for _, char := range t {
		countT[int(char-'a')]++
	}

	for i, x := range countT {
		if x != countS[i] {
			return byte(i + 'a')
		}
	}
	return 'a'
}
