package main

import (
	"fmt"
)

func main() {
	fmt.Println(test())
}

func test() bool {
	if countPrimes(10) != 4 {
		return false
	}
	if countPrimes(2) != 0 {
		return false
	}

	return true
}

func countPrimes(n int) int {
	if n < 2 {
		return 0
	}

	primesList := make([]int, n)
	primesList[0], primesList[1] = 1, 1

	for i := 2; i < n; i++ {
		if primesList[i] == 1 {
			continue
		}
		for j := 2; i*j < n; j++ {
			primesList[i*j] = 1
		}
	}
	ans := 0
	for i := 0; i < n; i++ {
		if primesList[i] == 0 {
			ans++
		}
	}

	return ans
}
