package main

import "fmt"

func main() {
	fmt.Println(test())
}

func test() bool {
	if ! isPalindrome(121) {
		return false
	}
	if isPalindrome(-121) {
		return false
	}
	if isPalindrome(10) {
		return false
	}
	return true
}

func isPalindrome(x int) bool {
    if x < 0 {
		return false
	}
	var num , y int
	y = x
	for y != 0 {
		num = num * 10 + y % 10
		y = y / 10
	}
	if num == x {
		return true
	}
	return false
}