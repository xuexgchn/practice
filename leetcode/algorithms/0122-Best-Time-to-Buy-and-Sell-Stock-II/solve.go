package main

import "fmt"

func main() {
	fmt.Println(test())
}

func test() bool {
	if maxProfit([]int{7, 1, 5, 3, 6, 4}) != 7 {
		return false
	}
	if maxProfit([]int{1, 2, 3, 4, 5}) != 4 {
		return false
	}
	if maxProfit([]int{7, 6, 4, 3, 1}) != 0 {
		return false
	}

	return true
}

func maxProfit(prices []int) int {
	ans := 0

	if len(prices) <= 1 {
		return 0
	}

	for i := 0; i < len(prices)-1; i++ {
		tmp := prices[i+1] - prices[i]

		if tmp > 0 {
			ans += tmp
		}
	}

	return ans
}
