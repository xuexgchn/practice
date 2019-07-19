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

// TreeNode define
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if p.Val > q.Val {
		p, q = q, p
	}
	for root != nil {
		if root.Val < p.Val {
			root = root.Right
		} else if root.Val > q.Val {
			root = root.Left
		} else {
			return root
		}
	}
	return root
}
