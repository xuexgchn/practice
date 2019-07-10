package main

import (
	"fmt"
)

func main() {
	fmt.Println(test())
}

func test() bool {
	if containsNearbyDuplicate([]int{1, 2, 3, 1}, 3) != true {
		return false
	}
	if containsNearbyDuplicate([]int{1, 0, 1, 1}, 3) != true {
		return false
	}
	if containsNearbyDuplicate([]int{1, 2, 3, 1, 2, 3}, 2) != false {
		return false
	}
	if containsNearbyDuplicate([]int{99, 99}, 2) != true {
		return false
	}
	return true
}

// Accept but not the best
// func containsNearbyDuplicate(nums []int, k int) bool {
// 	length := len(nums)
// 	for i := 0; i < length; i++ {
// 		max := min(i+k, length-1)
// 		for j := i + 1; j <= max; j++ {
// 			if nums[i] == nums[j] {
// 				return true
// 			}
// 		}
// 	}

// 	return false
// }

// func min(a int, b int) int {
// 	if a < b {
// 		return a
// 	}
// 	return b
// }

func containsNearbyDuplicate(nums []int, k int) bool {
	if k <= 0 {
		return false
	}
	count := make(map[int]int, len(nums))

	for i, x := range nums {
		value := count[x]
		if value != 0 && i-value+1 <= k {
			return true
		}
		count[x] = i + 1
	}

	return false
}
