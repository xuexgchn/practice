package main

import (
	"fmt"
)

func main() {
	fmt.Println(test())
}

func test() bool {
	if isPerfectSquare(16) != true {
		return false
	}
	if isPerfectSquare(14) != false {
		return false
	}
	return true
}

func isPerfectSquare(num int) bool {
	x := mySqrt(num)
	return x*x == num
}

// mySqrt problem 0069
func mySqrt(x int) int {
	low, high := 0, x
	mid := 0
	for low < high {
		mid = (low + high) / 2
		if mid*mid <= x && (mid+1)*(mid+1) > x {
			return mid
		} else if mid*mid < x {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return low
}
