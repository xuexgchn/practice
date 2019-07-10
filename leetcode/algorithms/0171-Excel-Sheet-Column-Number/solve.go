package main

import (
	"fmt"
)

func main() {
	fmt.Println(test())
}

func test() bool {
	if titleToNumber("A") != 1 {
		return false
	}
	if titleToNumber("AB") != 28 {
		return false
	}
	if titleToNumber("ZY") != 701 {
		return false
	}

	return true
}

func titleToNumber(s string) int {
	res := 0

	for i := 0; i < len(s); i++ {
		tmp := int(s[i] - 'A' + 1)
		res = res*26 + tmp
	}

	return res
}
