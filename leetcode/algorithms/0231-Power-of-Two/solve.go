package main

import (
	"fmt"
)

func main() {
	fmt.Println(test())
}

func test() bool {
	if isPowerOfTwo(1) != true {
		return false
	}
	if isPowerOfTwo(16) != true {
		return false
	}
	if isPowerOfTwo(218) != false {
		return false
	}
	return true
}

func isPowerOfTwo(n int) bool {
	if n < 1 {
		return false
	}

	for n > 1 {
		if n%2 == 1 {
			return false
		}
		n /= 2
	}

	return true
}
