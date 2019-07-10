package main

import (
	"fmt"
)

func main() {
	fmt.Println(test())
}

func test() bool {
	if missingNumber([]int{3, 0, 1}) != 2 {
		return false
	}
	if missingNumber([]int{9, 6, 4, 2, 3, 5, 7, 0, 1}) != 8 {
		return false
	}
	return true
}

func missingNumber(nums []int) int {
	length := len(nums)
	totalSum := length * (length + 1) / 2
	for _, x := range nums {
		totalSum = totalSum - x
	}

	return totalSum
}
