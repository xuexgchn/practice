package main

import "fmt"

func main() {
	fmt.Println(test())
}

func test() bool {
	if lengthOfLastWord("Hello World") != 5 {
		return false
	}
	if lengthOfLastWord("a ") != 1 {
		return false
	}
	if lengthOfLastWord("a") != 1 {
		return false
	}

	return true
}

func lengthOfLastWord(s string) int {
	count := 0
	ans := 0
	for _, char := range s {
		if string(char) == " " {
			if count != 0 {
				ans = count
			}
			count = 0
		} else {
			count++
		}
	}

	if count != 0 {
		ans = count
	}

	return ans
}
