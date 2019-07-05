package main

import "fmt"

func main() {
	fmt.Println(test())
}

func test() bool {
	return true
}

func getRow(rowIndex int) []int {
	var ans []int
	ans = append(ans, 1)
	for i := 1; i <= rowIndex; i++ {
		tmp := ans[i - 1] * (rowIndex - i + 1) / i
		ans = append(ans, tmp)
	}
	return ans
}