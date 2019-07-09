package main

import "fmt"

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

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	tmp := head

	for tmp.Next != nil {
		if tmp.Val == tmp.Next.Val {
			tmp.Next = tmp.Next.Next
		} else {
			tmp = tmp.Next
		}
	}

	return head
}
