package main

import (
	"fmt"

	"../../mytest"
)

func main() {
	fmt.Println(test())
}

func test() bool {
	if !mytest.IntSliceEqual(distributeCandies(7, 4), []int{1, 2, 3, 1}) {
		return false
	}
	if !mytest.IntSliceEqual(distributeCandies(10, 3), []int{5, 2, 3}) {
		return false
	}
	return true
}

func distributeCandies(candies int, n int) []int {
	ans := make([]int, n)

	for count := 1; candies >= 0; count++ {
		index := (count - 1) % n

		ans[index] = ans[index] + min(count, candies)
		candies = candies - count
	}
	return ans
}

func min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}
