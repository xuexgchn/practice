package main

import "fmt"

func main() {
	fmt.Println(test())
}

func test() bool {
	if fib(2) != 1 {
		return false
	}
	if fib(3) != 2 {
		return false
	}
	if fib(4) != 3 {
		return false
	}
	return true
}

func fib(N int) int {
    if N == 0 {
		return 0
	}
	if N == 1 {
		return 1
	}
	return fib(N - 1) + fib(N - 2)
}