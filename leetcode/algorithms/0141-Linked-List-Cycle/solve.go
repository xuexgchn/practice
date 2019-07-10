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

func hasCycle(head *ListNode) bool {
	if head == nil {
		return false
	}

	slowOne, fastOne := head, head.Next
	for fastOne != nil && fastOne.Next != nil && slowOne != fastOne {
		slowOne = slowOne.Next
		fastOne = fastOne.Next.Next
	}

	return slowOne == fastOne
}
