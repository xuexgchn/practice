package main

import "fmt"
import "math"

func main() {
	fmt.Println(test())
}

func test() bool {
	if reverse(123) != 321 {
		return false
	}
	if reverse(-123) != -321 {
		return false
	}
	if reverse(120) != 21 {
		return false
	}
	return true
}

func reverse(x int) int {
	var num int
	
	for x != 0 {
		num = num * 10 + x % 10
		x = x / 10
	}

	if num > math.MaxInt32 || num < math.MinInt32 {
		return 0
	}

    return num
}