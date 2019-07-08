package main

import (
	"fmt"

	"../../mytest"
)

func main() {
	fmt.Println(test())
}

func test() bool {
	if mytest.IntSliceEqual(plusOne([]int{1, 2, 3, 4}), []int{1, 2, 3, 5}) != true {
		return false
	}
	if mytest.IntSliceEqual(plusOne([]int{4, 3, 2, 1}), []int{4, 3, 2, 2}) != true {
		return false
	}
	if mytest.IntSliceEqual(plusOne([]int{1, 2, 3, 9}), []int{1, 2, 4, 0}) != true {
		return false
	}
	if mytest.IntSliceEqual(plusOne([]int{9, 9, 9, 9}), []int{1, 0, 0, 0, 0}) != true {
		return false
	}
	return true
}

func plusOne(digits []int) []int {
	length := len(digits)

	if length == 0 {
		return []int{1}
	}

	digits[length-1]++

	for i := length - 1; i > 0; i-- {
		if digits[i] < 10 {
			break
		}
		digits[i] -= 10
		digits[i-1]++
	}

	// head
	if digits[0] > 9 {
		digits[0] -= 10
		digits = append([]int{1}, digits...)
	}

	return digits
}
