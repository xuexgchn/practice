package main

import "fmt"

func main() {
	fmt.Println(test())
}

func test() bool {
	if romanToInt("III") != 3 {
		return false
	}
	if romanToInt("IV") != 4 {
		return false
	}
	if romanToInt("IX") != 9 {
		return false
	}
	if romanToInt("LVIII") != 58 {
		return false
	}
	if romanToInt("MCMXCIV") != 1994 {
		return false
	}
	return true
}

func romanToInt(s string) int {
	dic := map[byte]int {
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}
	ans := 0
	tmp := 0
	for i := len(s) - 1; i >= 0; i-- {
		now := dic[s[i]]
		// 正常罗马数字，越大的字符，在前面
		if tmp <= now { // 左加
			ans = ans + now
		} else { // 右减
			ans = ans - now
		}
		tmp = now
	}
	
    return ans
}