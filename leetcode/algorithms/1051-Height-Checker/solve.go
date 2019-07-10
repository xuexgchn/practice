package main

import (
	"fmt"
)

func main() {
	fmt.Println(test())
}

func test() bool {
	if heightChecker([]int{1, 1, 4, 2, 1, 3}) != 3 {
		return false
	}
	if heightChecker([]int{2, 1, 2, 1, 1, 2, 2, 1}) != 4 {
		return false
	}

	return true
}

func heightChecker(heights []int) int {
	//sort:According to the description
	count := make([]int, 105)
	for _, h := range heights {
		count[h]++
	}
	res := 0
	length := len(heights)

	for i, j := 0, 0; i < len(count); i++ {
		for count[i] > 0 {
			if j < length && heights[j] != i {
				res++
			}
			count[i]--
			j++
		}
	}

	return res
}
