package main

import (
	"fmt"
)

func main() {
	fmt.Println(test())
}

func test() bool {
	if rob([]int{1, 2, 3, 1}) != 4 {
		return false
	}
	if rob([]int{2, 7, 9, 3, 1}) != 12 {
		return false
	}
	if rob([]int{2, 1, 1, 2}) != 4 {
		return false
	}

	return true
}

func rob(nums []int) int {
	length := len(nums)

	if length == 0 {
		return 0
	}
	if length == 1 {
		return nums[0]
	}
	if length == 2 {
		return max(nums[0], nums[1])
	}

	// dp fun: dp[i] = max(dp[i - 1], dp[i - 2] + nums[i-1])
	dp := make([]int, length+1)
	dp[0] = 0
	dp[1] = nums[0]
	for i := 2; i <= length; i++ {
		dp[i] = max(dp[i-1], dp[i-2]+nums[i-1])
	}
	return dp[length]
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
