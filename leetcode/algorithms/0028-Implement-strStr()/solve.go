package main

import (
	"fmt"
)

func main() {
	fmt.Println(test())
}

func test() bool {
	if strStr("hello", "") != 0 {
		return false
	}
	if strStr("hello", "ll") != 2 {
		return false
	}
	if strStr("aaaaa", "bba") != -1 {
		return false
	}
	return true
}

func strStr(haystack string, needle string) int {
	if needle == "" {
		return 0
	}

	hlen := len(haystack)
	nlen := len(needle)

	for i := 0; i <= hlen-nlen; i++ {
		if haystack[i:i+nlen] == needle {
			return i
		}
	}
	return -1
}
