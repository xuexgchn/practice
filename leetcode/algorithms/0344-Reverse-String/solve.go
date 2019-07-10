package main

import (
	"fmt"
)

func main() {
	fmt.Println(test())
}

func test() bool {
	return true
}

func reverseString(s []byte) {
	i, j := 0, len(s)-1

	for i < j {
		s[i], s[j] = s[j], s[i]
		i++
		j--
	}
}
