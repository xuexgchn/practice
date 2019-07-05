package main

import "fmt"

func main() {
	fmt.Println(test())
}

func test() bool {
	if toLowerCase("Hello") != "hello" {
		return false
	}
	if toLowerCase("here") != "here" {
		return false
	}
	if toLowerCase("LOVELY") != "lovely" {
		return false
	}
	return true
}

func toLowerCase(str string) string {
	ans := ""
	for i := 0; i < len(str); i++ {
		if str[i] >=  65 && str[i] <= 90 {
			ans = ans + string(str[i] + 32)
		} else {
			ans = ans + string(str[i])
		}
	}
	return ans
}