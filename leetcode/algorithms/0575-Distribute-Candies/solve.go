package main

import (
	"fmt"
)

func main() {
	fmt.Println(test())
}

func test() bool {
	if distributeCandies([]int{1, 1, 2, 2, 3, 3}) != 3 {
		return false
	}
	if distributeCandies([]int{1, 1, 2, 3}) != 2 {
		return false
	}

	return true
}

func distributeCandies(candies []int) int {
	count := map[int]int{}
	length := len(candies)

	for i := 0; i < length; i++ {
		count[candies[i]] = 1
	}

	return min(length/2, len(count))
}

func min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}
