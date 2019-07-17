package main

import (
	"fmt"
)

func main() {
	fmt.Println(test())
}

func test() bool {
	if minMoves([]int{1, 2, 3}) != 3 {
		return false
	}
	return true
}

func minMoves(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	// 可以理解成，每次让一个值减小1，使每个值最终等于最小值
	min := nums[0]
	for i := 0; i < len(nums); i++ {
		if nums[i] < min {
			min = nums[i]
		}
	}
	ans := 0
	for i := 0; i < len(nums); i++ {
		ans += nums[i] - min
	}
	return ans
}
