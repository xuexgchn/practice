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

	for i := 0; i < length-1; i++ {
		for j := length - 1; j > i; j-- {
			sum := numbers[i] + numbers[j]
			if sum > target {
				continue
			} else if sum == target {
				return []int{i + 1, j + 1}
			} else {
				break
			}
		}
	}

	return []int{}
}
