package main

import (
	"fmt"
)

func main() {
	fmt.Println(test())
}

func test() bool {
	if checkPerfectNumber(28) != true {
		return false
	}
	return true
}

func checkPerfectNumber(num int) bool {
	if num <= 2 {
		return false
	}
	sum := 0
	for i := 2; i <= num/i; i++ {
		if num%i == 0 {
			sum = sum + i + num/i
		}
	}
	return sum+1 == num
}
