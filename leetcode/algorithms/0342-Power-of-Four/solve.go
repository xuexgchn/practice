package main

import (
	"fmt"
)

func main() {
	fmt.Println(test())
}

func test() bool {
	if isPowerOfFour(16) != true {
		return false
	}
	if isPowerOfFour(5) != false {
		return false
	}

	return true
}

func isPowerOfFour(n int) bool {
	if n < 1 {
		return false
	}

	for n > 1 {
		if n%4 != 0 {
			return false
		}
		n /= 4
	}

	// A better solve
	// return num > 0 && num&(num-1) == 0 && 0xaaaaaaaa&num == 0

	return true
}
