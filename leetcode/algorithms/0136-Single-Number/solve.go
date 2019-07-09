package main

import (
	"fmt"
)

func main() {
	fmt.Println(test())
}

func test() bool {
	if singleNumber([]int{2, 2, 1}) != 1 {
		return false
	}
	if singleNumber([]int{4, 1, 2, 1, 2}) != 4 {
		return false
	}

	return true
}

func singleNumber(nums []int) int {
	res := 0
	for _, n := range nums {
		res ^= n
	}

	return res
}
