package main

import (
	"fmt"
)

func main() {
	fmt.Println(test())
}

func test() bool {
	if canConstruct("a", "b") != false {
		return false
	}
	if canConstruct("aa", "ab") != false {
		return false
	}
	if canConstruct("aa", "aab") != true {
		return false
	}

	return true
}

func canConstruct(ransomNote string, magazine string) bool {
	countA := make([]int, 30)
	countB := make([]int, 30)

	for _, char := range ransomNote {
		countA[char-'a']++
	}

	for _, char := range magazine {
		countB[char-'a']++
	}

	for i, x := range countB {
		if x < countA[i] {
			return false
		}
	}

	return true
}
