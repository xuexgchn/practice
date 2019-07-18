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

func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	if root.Left == nil {
		return 1 + minDepth(root.Right)
	}
	if root.Right == nil {
		return 1 + minDepth(root.Left)
	}
	return 1 + min(minDepth(root.Left), minDepth(root.Right))
}

func min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}
