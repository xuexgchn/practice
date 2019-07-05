package main

import "fmt"

func main() {
	fmt.Println(test())
}

func test() bool {
	if climbStairs(2) != 2 {
		return false
	}
	if climbStairs(3) != 3 {
		return false
	}

	return true
}

func climbStairs(n int) int {
	fibs := make([]int, n)
    if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}
	fibs[0] = 1
	fibs[1] = 2
	for i := 2; i < n; i++ {
		fibs[i] = fibs[i - 1] + fibs[i - 2]
		
	}
	return fibs[n - 1]
}