package main

import (
	"fmt"
)

func main() {
	fmt.Println(test())
}

func test() bool {
	if removeDuplicates([]int{1, 1, 2}) != 2 {
		return false
	}
	if removeDuplicates([]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}) != 5 {
		return false
	}
	return true
}

func removeDuplicates(nums []int) int {
	length := len(nums)
	if length == 0 {
		return 0
	}

	index := 0
	ans := 0

	for i := 0; i < length; i++ {
		if nums[i] != nums[index] {
			index = i
			ans++
		}
		// According to the problem description, we need to modify the nums
		nums[ans] = nums[i]
	}

	return ans + 1
}
