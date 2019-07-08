package main

import "fmt"

func main() {
	fmt.Println(test())
}

func test() bool {
	fmt.Println(plusOne([]int{1, 2, 3, 4}))
	fmt.Println(plusOne([]int{4, 3, 2, 1}))
	fmt.Println(plusOne([]int{4, 3, 2, 9}))
	fmt.Println(plusOne([]int{9, 9, 9, 9}))
	return true
}

func plusOne(digits []int) []int {
	length := len(digits)

	if length == 0 {
		return []int{1}
	}

	digits[length-1]++

	for i := length - 1; i > 0; i-- {
		if digits[i] < 10 {
			break
		}
		digits[i] -= 10
		digits[i-1]++
	}

	// head
	if digits[0] > 9 {
		digits[0] -= 10
		digits = append([]int{1}, digits...)
	}

	return digits
}
