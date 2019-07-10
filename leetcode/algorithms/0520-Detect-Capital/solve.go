package main

import (
	"fmt"
)

func main() {
	fmt.Println(test())
}

func test() bool {
	if detectCapitalUse("USA") != true {
		return false
	}
	if detectCapitalUse("FlaG") != false {
		return false
	}
	if detectCapitalUse("mL") != false {
		return false
	}
	if detectCapitalUse("cVe") != false {
		return false
	}

	return true
}

func detectCapitalUse(word string) bool {
	length := len(word)
	if length < 2 {
		return true
	}

	if length == 2 {
		return isUpper(word[0]) || (!isUpper(word[0]) && !isUpper(word[1]))
	}

	if isUpper(word[0]) {
		nextFlag := isUpper(word[1])
		for i := 2; i < length; i++ {
			if isUpper(word[i]) != nextFlag {
				return false
			}
		}
	} else {
		for i := 1; i < length; i++ {
			if isUpper(word[i]) {
				return false
			}
		}
	}
	return true
}

func isUpper(c byte) bool {
	if c >= 'A' && c <= 'Z' {
		return true
	}
	return false
}
