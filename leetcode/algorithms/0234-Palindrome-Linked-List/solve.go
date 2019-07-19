package main

import (
	"fmt"
)

func main() {
	fmt.Println(test())
}

func test() bool {
	return true
}

// ListNode define
type ListNode struct {
	Val  int
	Next *ListNode
}

func isPalindrome(head *ListNode) bool {
	nums := []int{}

	for head != nil {
		nums = append(nums, head.Val)
		head = head.Next
	}

	i, j := 0, len(nums)-1
	for i < j {
		if nums[i] != nums[j] {
			return false
		}
		i++
		j--
	}

	return true
}
