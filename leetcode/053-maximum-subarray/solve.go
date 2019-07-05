package main

import "fmt"

func main() {
	fmt.Println(test())
}

func test() bool {
	if maxSubArray([]int{-2,1,-3,4,-1,2,1,-5,4}) != 6 {
		return false
	}
	
	return true
}

func maxSubArray(nums []int) int {
	length := len(nums)
	if length == 0 {
		return 0
	}
	if length == 1 {
		return nums[0]
	}
	dp := nums
	ans := nums[0]
	// dp[i] 表示：以索引i结尾的最大子序列和
	for i := 1; i < length; i++ {
		if dp[i - 1] > 0 {
			dp[i] = dp[i - 1] + nums[i]
		}
		if dp[i] > ans {
			ans = dp[i]
		}
	}
	return ans
}