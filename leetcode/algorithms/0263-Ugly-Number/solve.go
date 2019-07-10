package main

import (
	"fmt"
)

func main() {
	fmt.Println(test())
}

func test() bool {
	if isUgly(6) != true {
		return false
	}
	if isUgly(8) != true {
		return false
	}
	if isUgly(14) != false {
		return false
	}

	return true
}

func isUgly(num int) bool {
	if num == 0 {
		return false
	}

	for num != 1 {
		if num%2 == 0 {
			num /= 2
		} else if num%3 == 0 {
			num /= 3
		} else if num%5 == 0 {
			num /= 5
		} else {
			return false
		}
	}

	return true
}
