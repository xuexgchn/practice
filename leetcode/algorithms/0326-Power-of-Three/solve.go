package main

import (
	"fmt"
)

func main() {
	fmt.Println(test())
}

func test() bool {
	if isPowerOfThree(27) != true {
		return false
	}
	if isPowerOfThree(0) != false {
		return false
	}
	if isPowerOfThree(9) != true {
		return false
	}
	if isPowerOfThree(45) != false {
		return false
	}

	return true
}

func isPowerOfThree(n int) bool {
	if n < 1 {
		return false
	}

	for n > 1 {
		if n%3 != 0 {
			return false
		}
		n /= 3
	}

	return true
}
