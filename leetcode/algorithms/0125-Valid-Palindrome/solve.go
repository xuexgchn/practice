package main

import (
	"fmt"
)

func main() {
	fmt.Println(test())
}

func test() bool {
	if isPalindrome("A man, a plan, a canal: Panama") != true {
		return false
	}
	if isPalindrome("race a car") != false {
		return false
	}
	if isPalindrome("") != true {
		return false
	}

	return true
}

func isPalindrome(s string) bool {

	newString := toString(s)
	i, j := 0, len(newString)-1

	for i < j {
		if newString[i] != newString[j] {
			return false
		}
		i++
		j--
	}
	return true
}

func toString(s string) string {
	newString := ""

	for _, char := range s {
		if '0' <= char && char <= '9' {
			newString = newString + string(char)
		} else if 'a' <= char && char <= 'z' {
			newString = newString + string(char)
		} else if 'A' <= char && char <= 'Z' {
			newString = newString + string(char+'a'-'A')
		} else {
			continue
		}
	}

	return newString
}
