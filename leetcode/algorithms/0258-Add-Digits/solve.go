package main

import (
	"fmt"
)

func main() {
	fmt.Println(test())
}

func test() bool {
	if addDigits(38) != 2 {
		return false
	}

	return true
}

func addDigits(num int) int {
	if num == 0 {
		return 0
	}
	ans := num % 9
	if ans == 0 {
		return 9
	}
	return ans
}
