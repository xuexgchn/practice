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

func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	root.Left, root.Right = invertTree(root.Right), invertTree(root.Left)

	return root
}
