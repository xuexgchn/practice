package main

import (
	"fmt"
)

func main() {
	fmt.Println(test())
}

func test() bool {
	if isAnagram("anagram", "nagaram") != true {
		return false
	}
	if isAnagram("rat", "car") != false {
		return false
	}

	return true
}
func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	countS := make([]int, 30)
	countT := make([]int, 30)

	for i := 0; i < len(s); i++ {
		countS[int(s[i])-int('a')]++
		countT[int(t[i])-int('a')]++
	}

	for i := 0; i < 30; i++ {
		if countS[i] != countT[i] {
			return false
		}
	}

	return true
}
