package main

import (
	"fmt"
)

func main() {
	fmt.Println(test())
}

func test() bool {
	if findMaxConsecutiveOnes([]int{1, 1, 0, 1, 1, 1}) != 3 {
		return false
	}
	return true
}

func findMaxConsecutiveOnes(nums []int) int {
	count, ans := 0, 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == 1 {
			count++
			if count > ans {
				ans = count
			}
		} else {
			count = 0
		}
	}

	return ans
}
