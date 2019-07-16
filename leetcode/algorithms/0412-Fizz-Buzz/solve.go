package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(test())
}

func test() bool {
	return true
}

func fizzBuzz(n int) []string {
	if n < 1 {
		return []string{}
	}
	ans := make([]string, n)

	for i := 1; i <= n; i++ {
		if i%3 == 0 && i%5 == 0 {
			ans[i-1] = "FizzBuzz"
		} else if i%3 == 0 {
			ans[i-1] = "Fizz"
		} else if i%5 == 0 {
			ans[i-1] = "Buzz"
		} else {
			ans[i-1] = strconv.Itoa(i)
		}
	}

	return ans
}
