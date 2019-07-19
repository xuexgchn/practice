package main

import (
	"fmt"
	"strconv"
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

func binaryTreePaths(root *TreeNode) []string {
	if root == nil {
		return nil
	}

	res := make([]string, 0, 16)

	var dfs func(string, *TreeNode)
	dfs = func(tmp string, root *TreeNode) {
		if tmp == "" {
			tmp = strconv.Itoa(root.Val)
		} else {
			tmp += "->" + strconv.Itoa(root.Val)
		}

		if root.Left != nil {
			dfs(tmp, root.Left)
		}

		if root.Right != nil {
			dfs(tmp, root.Right)
		}

		if root.Left == nil && root.Right == nil {
			res = append(res, tmp)
		}
	}

	dfs("", root)

	return res
}
