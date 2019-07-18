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

func isBalanced(root *TreeNode) bool {
	_, res := checkBalanced(root)
	return res
}

func checkBalanced(root *TreeNode) (int, bool) {
	if root == nil {
		return 0, true
	}
	leftDepth, leftRes := checkBalanced(root.Left)
	rightDepth, rightRes := checkBalanced(root.Right)

	if leftRes && rightRes && abs(leftDepth, rightDepth) <= 1 {
		return max(leftDepth, rightDepth) + 1, true
	}

	return 0, false
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func abs(a int, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}
