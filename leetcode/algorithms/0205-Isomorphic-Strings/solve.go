package main

import (
	"fmt"
)

func main() {
	fmt.Println(test())
}

func test() bool {
	if isIsomorphic("egg", "add") != true {
		return false
	}
	if isIsomorphic("foo", "bar") != false {
		return false
	}
	if isIsomorphic("paper", "title") != true {
		return false
	}
	if isIsomorphic("ab", "aa") != false {
		return false
	}

	return true
}

func isIsomorphic(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	countS := make([]int, 300)
	countT := make([]int, 300)

	for i := 0; i < len(s); i++ {
		if countS[int(s[i])] != countT[int(t[i])] {
			return false
		}
		countS[int(s[i])] = i + 1
		countT[int(t[i])] = i + 1
	}

	return true
}
