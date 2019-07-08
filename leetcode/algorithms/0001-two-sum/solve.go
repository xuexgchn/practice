package main

import (
	"fmt"

	"../../mytest"
)

func main() {
	fmt.Print(test())
}

func test() bool {
	if mytest.IntSliceEqual(twoSum([]int{2, 7, 11, 15}, 9), []int{0, 1}) != true {
		return false
	}
	return true
}

func twoSum(nums []int, target int) []int {
	var numsLen = len(nums)
	for i := 0; i < numsLen-1; i++ {
		for j := i + 1; j < numsLen; j++ {
			sum := nums[i] + nums[j]
			if target == sum {
				return []int{i, j}
			}
		}
	}
	return []int{}
}
