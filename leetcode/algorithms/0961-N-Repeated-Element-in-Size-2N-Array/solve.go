package main

import (
	"fmt"
)

func main() {
	fmt.Println(test())
}

func test() bool {
	if repeatedNTimes([]int{1, 2, 3, 3}) != 3 {
		return false
	}
	if repeatedNTimes([]int{2, 1, 2, 5, 3, 2}) != 2 {
		return false
	}
	if repeatedNTimes([]int{5, 1, 5, 2, 5, 3, 5, 4}) != 5 {
		return false
	}

	return true
}

func repeatedNTimes(A []int) int {
	count := make([]int, 10005)

	for _, x := range A {
		if count[x] > 0 {
			return x
		}
		count[x]++
	}

	return 0
}
