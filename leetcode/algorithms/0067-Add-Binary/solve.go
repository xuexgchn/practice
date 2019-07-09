package main

import "fmt"

func main() {
	fmt.Println(test())
}

func test() bool {
	if addBinary("", "") != "0" {
		return false
	}
	if addBinary("0", "0") != "0" {
		return false
	}
	if addBinary("11", "1") != "100" {
		return false
	}
	if addBinary("1010", "1011") != "10101" {
		return false
	}
	return true
}

func addBinary(a string, b string) string {
	length := max(len(a), len(b))
	return makeString(add(toInt(a, length), toInt(b, length)))
}

func makeString(nums []int) string {
	bytes := make([]byte, len(nums))
	for i := range bytes {
		bytes[i] = byte(nums[i]) + '0'
	}

	return string(bytes)
}

func toInt(s string, length int) []int {
	res := make([]int, length)

	lenS := len(s)
	for i, char := range s {
		// reverse
		res[lenS-i-1] = int(char - '0')
	}

	return res
}

func add(a []int, b []int) []int {
	length := len(a) + 1
	res := make([]int, length)

	for i := 0; i < length-1; i++ {
		tmp := a[i] + b[i] + res[i]
		res[i+1] = tmp / 2
		res[i] = tmp % 2
	}

	finalRes := []int{}
	for i := length - 1; i >= 0; i-- {
		if res[i] == 0 && len(finalRes) == 0 {
			continue
		}
		finalRes = append(finalRes, res[i])
	}
	if len(finalRes) == 0 {
		finalRes = append(finalRes, 0)
	}
	return finalRes
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
