package main

import "fmt"
import "../../mytest"

func main() {
	fmt.Println(test())
}

func test() bool {
	if mytest.IntSliceEqual(intersection([]int{1, 2, 2, 1}, []int{2, 2}), []int{2}) != true {
		return false
	}
	if mytest.IntSliceEqual(intersection([]int{4, 9, 5}, []int{9, 4, 9, 8, 4}), []int{4, 9}) != true {
		return false
	}
	return true
}

func intersection(nums1 []int, nums2 []int) []int {
	count := map[int]int{}
	for i := 0; i < len(nums1); i++ {
		count[nums1[i]] = 0
	}
	for i := 0; i < len(nums2); i++ {
		if _, ok := count[nums2[i]]; ok {
			count[nums2[i]] = 1
		}
	}

	var ans []int
	for k, v := range count {
		if v != 0 {
			ans = append(ans, k)
		}
	}

	return ans
}
