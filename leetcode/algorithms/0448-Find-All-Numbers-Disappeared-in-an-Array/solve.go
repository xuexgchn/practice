package main

import (
	"fmt"

	"../../mytest"
)

func main() {
	fmt.Println(test())
}

func test() bool {
	if !mytest.IntSliceEqual(findDisappearedNumbers([]int{4, 3, 2, 7, 8, 2, 3, 1}), []int{5, 6}) {
		return false
	}
	return true
}

func findDisappearedNumbers(nums []int) []int {
	for i := 0; i < len(nums); i++ {
		for nums[i] != nums[nums[i]-1] {
			nums[i], nums[nums[i]-1] = nums[nums[i]-1], nums[i]
		}
	}

	res := make([]int, 0, len(nums))

	for i, x := range nums {
		if x != i+1 {
			res = append(res, i+1)
		}
	}

	return res
}
