package main

import (
	"fmt"

	"../../mytest"
)

func main() {
	fmt.Println(test())
}

func test() bool {
	if mytest.IntSliceEqual(twoSum([]int{2, 7, 11, 15}, 9), []int{1, 2}) != true {
		return false
	}

	return true
}

func twoSum(numbers []int, target int) []int {
	length := len(numbers)
	if length <= 1 {
		return []int{}
	}

	head, tail := 0, length-1
	for head < tail {
		sum := numbers[head] + numbers[tail]
		if sum > target {
			tail--
		} else if sum == target {
			return []int{head + 1, tail + 1}
		} else {
			head++
		}
	}

	return []int{}
}
