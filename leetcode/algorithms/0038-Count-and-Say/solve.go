package main

import "fmt"
import "strconv"

func main() {
	fmt.Println(test())
}

func test() bool {
	if countAndSay(1) != "1" {
		return false
	}
	if countAndSay(2) != "11" {
		return false
	}
	if countAndSay(3) != "21" {
		return false
	}
	if countAndSay(4) != "1211" {
		return false
	}
	if countAndSay(5) != "111221" {
		return false
	}

	return true
}

func countAndSay(n int) string {
	if n == 1 {
		return "1"
	}

	return count(countAndSay(n - 1))
}

func count(s string) string {
	countChar := string(s[0])
	count := 1
	result := ""
	for _, char := range s[1:] {
		if string(char) == countChar {
			count++
		} else {
			result = result + strconv.Itoa(count) + string(countChar)
			countChar = string(char)
			count = 1
		}
	}
	result = result + strconv.Itoa(count) + string(countChar)
	return result
}
