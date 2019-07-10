package main

import (
	"fmt"
)

func main() {
	fmt.Println(test())
}

func test() bool {
	if canWinNim(4) != false {
		return false
	}
	return true
}

func canWinNim(n int) bool {
	return n%4 != 0
}
