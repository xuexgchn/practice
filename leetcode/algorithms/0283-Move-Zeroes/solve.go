package main

import (
	"fmt"
)

func main() {
	fmt.Println(test())
}

func test() bool {
	moveZeroes([]int{0, 1, 0, 3, 12})
	return true
}

func moveZeroes(nums []int) {
	i, j := 0, 0
	length := len(nums)

	for ; i < length; i++ {
		if nums[i] != 0 {
			nums[j] = nums[i]
			j++
		}
	}

	for i = j; i < length; i++ {
		nums[i] = 0
	}

	// fmt.Println(nums)
}
