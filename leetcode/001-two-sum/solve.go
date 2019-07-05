package main

import (
	"fmt"
)

func twoSum(nums []int, target int) []int {
	var numsLen = len(nums)
	for i := 0; i < numsLen - 1; i++ {
		for j := i + 1; j < numsLen; j++ {
			sum := nums[i] + nums[j]
			if target == sum {
				return []int{i, j}
			}
		}
	}
	return []int{}
}

func equal(x, y []int) bool {
    if len(x) != len(y) {
        return false
    }
    for i := range x {
        if x[i] != y[i] {
            return false
        }
    }
    return true
}

func test() bool {
	if ! equal(twoSum([]int{2, 7, 11, 15}, 9), []int{0, 1}) {
		return false
	}
	return true
}

func main() {
	fmt.Print(test())
}