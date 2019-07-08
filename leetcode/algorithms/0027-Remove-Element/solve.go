package main

import (
	"fmt"
)

func main() {
	fmt.Println(test())
}

func test() bool {
	if removeElement([]int{3, 2, 2, 3}, 3) != 2 {
		return false
	}
	if removeElement([]int{0, 1, 2, 2, 3, 0, 4, 2}, 2) != 5 {
		return false
	}
	return true
}

func removeElement(nums []int, val int) int {
	index := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != val {
			nums[index] = nums[i]
			index++
		}
	}

	return index
}
