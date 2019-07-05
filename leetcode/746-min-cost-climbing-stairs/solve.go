package main

import "fmt"

func main() {
	fmt.Println(test())
}

func test() bool {
	if minCostClimbingStairs([]int{10, 15, 20}) != 15 {
		return false
	}
	if minCostClimbingStairs([]int{1, 100, 1, 1, 1, 100, 1, 1, 100, 1}) != 6 {
		return false
	}
	return true
}

func minCostClimbingStairs(cost []int) int {
	min_int := func(a, b int) int {
		if a > b {
			return b
		} 
		return a
	}

	lenCost := len(cost)
	if lenCost == 0 {
		return 0
	}
	if lenCost == 1 {
		return cost[1]
	}
	if lenCost == 2 {
		return min_int(cost[0], cost[1])
	}

	dp := []int{0, 0}
	for i := 2; i <= lenCost; i++ {
		tmp := min_int(dp[i - 1] + cost[i - 1], dp[i - 2] + cost[i - 2])
		dp = append(dp, tmp)
	}
	return dp[lenCost]
}