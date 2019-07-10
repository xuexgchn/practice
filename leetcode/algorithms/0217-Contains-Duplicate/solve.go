package main

import (
	"fmt"
)

func main() {
	fmt.Println(test())
}

func test() bool {
	if containsDuplicate([]int{1, 2, 3, 1}) != true {
		return false
	}
	if containsDuplicate([]int{1, 2, 3, 4}) != false {
		return false
	}

	if containsDuplicate([]int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2}) != true {
		return false
	}
	return true
}

func containsDuplicate(nums []int) bool {
	count := make(map[int]bool, len(nums))

	for _, x := range nums {
		if count[x] {
			return true
		}
		count[x] = true
	}

	return false
}
