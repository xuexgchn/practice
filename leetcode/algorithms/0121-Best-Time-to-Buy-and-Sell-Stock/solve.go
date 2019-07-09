package main

import "fmt"

func main() {
	fmt.Println(test())
}

func test() bool {
	if maxProfit([]int{7, 1, 5, 3, 6, 4}) != 5 {
		return false
	}
	if maxProfit([]int{7, 6, 4, 3, 1}) != 0 {
		return false
	}

	return true
}

func maxProfit(prices []int) int {
	length := len(prices)
	if length == 0 {
		return 0
	}

	minCost := prices[0]
	maxProfit := 0

	for i := 0; i < length; i++ {
		minCost = min(prices[i], minCost)
		maxProfit = max(prices[i]-minCost, maxProfit)
	}

	return maxProfit
}

func min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
