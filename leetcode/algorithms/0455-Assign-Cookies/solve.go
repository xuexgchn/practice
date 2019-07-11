package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(test())
}

func test() bool {
	if findContentChildren([]int{1, 2, 3}, []int{1, 1}) != 1 {
		return false
	}
	if findContentChildren([]int{1, 2}, []int{1, 2, 3}) != 2 {
		return false
	}
	if findContentChildren([]int{10, 9, 8, 7}, []int{5, 6, 7, 8}) != 2 {
		return false
	}

	return true
}

func findContentChildren(g []int, s []int) int {
	sort.Ints(g)
	sort.Ints(s)

	res := 0
	for i, j := 0, 0; i < len(g) && j < len(s); {
		if g[i] <= s[j] {
			res++
			i++
		}
		j++
	}

	return res
}
