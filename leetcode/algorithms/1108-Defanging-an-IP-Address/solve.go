package main

import (
	"fmt"
)

func main() {
	fmt.Println(test())
}

func test() bool {
	if defangIPaddr("1.1.1.1") != "1[.]1[.]1[.]1" {
		return false
	}

	if defangIPaddr("255.100.50.0") != "255[.]100[.]50[.]0" {
		return false
	}

	return true
}

func defangIPaddr(address string) string {
	res := ""
	for _, char := range address {
		if char == '.' {
			res = res + "[.]"
		} else {
			res = res + string(char)
		}
	}
	return res
}
