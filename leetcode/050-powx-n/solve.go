package main

import "fmt"

func main() {
	fmt.Println(test())
}

func test() bool {
	if myPow(2.00000, 10) != 1024.00000 {
		return false
	}
	if myPow(2.10000, 3) != 9.26100 {
		return false
	}
	if myPow(2.00000, -2) != 0.25000 {
		return false
	}
	if myPow(0.44528, 0) != 1.0 {
		return false
	}
	return true
}

func myPow(x float64, n int) float64 {
	if n == 0 {
		return 1.0
	}
	if n == 1 {
		return x
	} 
	if n == -1 {
		return 1 / x
	}
	if n % 2 == 0 {
		tmp := myPow(x, n / 2)
		return tmp * tmp
	} else {
		tmp := myPow(x, (n - 1) / 2)
		return tmp * tmp * x
	}
}