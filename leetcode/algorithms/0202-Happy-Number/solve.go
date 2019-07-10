package main

import (
	"fmt"
)

func main() {
	fmt.Println(test())
}

func test() bool {
	if isHappy(1) != true {
		return false
	}

	if isHappy(19) != true {
		return false
	}

	return true
}

func isHappy(n int) bool {
	count := map[int]int{}

	for n != 1 {
		if _, ok := count[n]; ok {
			return false
		}
		count[n] = 1
		tmp := 0

		for n != 0 {
			tmp = tmp + (n%10)*(n%10)
			n /= 10
		}
		n = tmp
	}

	return true
}
