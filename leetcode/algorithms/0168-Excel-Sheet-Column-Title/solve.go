package main

import (
	"fmt"
)

func main() {
	fmt.Println(test())
}

func test() bool {
	if convertToTitle(1) != "A" {
		return false
	}
	if convertToTitle(26) != "Z" {
		return false
	}
	if convertToTitle(27) != "AA" {
		return false
	}
	if convertToTitle(28) != "AB" {
		return false
	}
	if convertToTitle(701) != "ZY" {
		return false
	}
	if convertToTitle(52) != "AZ" {
		return false
	}

	return true
}

func convertToTitle(n int) string {
	res := ""

	for n > 0 {
		n--
		res = string(n%26+'A') + res
		n /= 26
	}

	return res
}
