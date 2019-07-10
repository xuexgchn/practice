package main

import (
	"fmt"
)

func main() {
	fmt.Println(test())
}

func test() bool {
	if trailingZeroes(3) != 0 {
		return false
	}
	if trailingZeroes(5) != 1 {
		return false
	}

	return true
}

func trailingZeroes(n int) int {
	res := 0

	for n >= 5 {
		n /= 5
		res += n
	}

	return res
}
