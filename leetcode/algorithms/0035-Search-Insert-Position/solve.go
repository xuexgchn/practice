package main

import "fmt"

func main() {
	fmt.Println(test())
}

func test() bool {
	if searchInsert([]int{1, 3, 5, 6}, 5) != 2 {
		return false
	}
	if searchInsert([]int{1, 3, 5, 6}, 2) != 1 {
		return false
	}
	if searchInsert([]int{1, 3, 5, 6}, 7) != 4 {
		return false
	}
	if searchInsert([]int{1, 3, 5, 6}, 0) != 0 {
		return false
	}
	return true
}

func searchInsert(nums []int, target int) int {
	i := 0
	for ; i < len(nums); i++ {
		if target <= nums[i] {
			return i
		}
	}

	return i
}
