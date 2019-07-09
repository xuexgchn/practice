package main

import "fmt"

func main() {
	fmt.Println(test())
}

func test() bool {
	if mySqrt(0) != 0 {
		return false
	}
	if mySqrt(1) != 1 {
		return false
	}
	if mySqrt(2) != 1 {
		return false
	}
	if mySqrt(4) != 2 {
		return false
	}
	if mySqrt(8) != 2 {
		return false
	}
	return true
}

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
