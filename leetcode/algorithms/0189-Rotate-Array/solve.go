package main

import (
	"fmt"
)

func main() {
	fmt.Println(test())
}

func test() bool {
	rotate([]int{1, 2, 3, 4, 5, 6, 7}, 3)
	return true
}

func rotate(nums []int, k int) {
	n := len(nums)

	if k == 0 || k == n {
		return
	}

	k %= n

	// reverse all
	reverse(nums, 0, n-1)
	// reverse head
	reverse(nums, 0, k-1)
	// reverse tail
	reverse(nums, k, n-1)

	// fmt.Println(nums)
}

func reverse(nums []int, i int, j int) {
	for i < j {
		nums[i], nums[j] = nums[j], nums[i]
		i++
		j--
	}
}
