package main

import "fmt"

func main() {
	fmt.Println(test())
}

func test() bool {
	if longestCommonPrefix([]string{}) != "" {
		return false
	}
	if longestCommonPrefix([]string{"aa", "a"}) != "a" {
		return false
	}
	if longestCommonPrefix([]string{"a", "a"}) != "a" {
		return false
	}
	if longestCommonPrefix([]string{"flower","flow","flight"}) != "fl" {
		return false
	}
	if longestCommonPrefix([]string{"dog","racecar","car"}) != "" {
		return false
	}
	return true
}


func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	first_one := strs[0]
	if len(strs) == 1 {
		return first_one
	}
	i := 0	
	for i = 0; i < len(first_one); i++ {
		for j := 1; j < len(strs); j++ {
			if i >= len(strs[j]) {
				return first_one[:i]
			}
			if strs[j][i] != first_one[i] {
				return first_one[:i]
			}

		}
	}
    return first_one[:i]
}