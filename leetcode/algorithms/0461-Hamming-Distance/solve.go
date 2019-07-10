package main

import (
	"fmt"
)

func main() {
	fmt.Println(test())
}

func test() bool {
	if hammingDistance(1, 4) != 2 {
		return false
	}
	return true
}

func hammingDistance(x int, y int) int {
	x = x ^ y

	res := 0
	for x > 0 {
		res += x & 1
		x >>= 1
	}

	return res
}
