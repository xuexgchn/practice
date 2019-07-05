package main

import "fmt"

func main() {
	fmt.Println(test())
}

func test() bool {
	return true
}

func generate(numRows int) [][]int {
	if numRows == 0 {
		return nil
	}
	var ans [][]int
	ans = append(ans, []int{1})
	for i := 1; i < numRows; i++ {
		tmp := []int{1}
		for j := 1; j <= i - 1; j++ {
			tmp = append(tmp, ans[i - 1][j - 1] + ans[i - 1][j])
		}
		tmp = append(tmp, 1)
		ans = append(ans, tmp)
	}
    return ans
}