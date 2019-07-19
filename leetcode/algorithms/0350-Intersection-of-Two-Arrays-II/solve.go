package main

import "fmt"
import "../../mytest"

func main() {
	fmt.Println(test())
}

func test() bool {
	if mytest.IntSliceEqual(intersect([]int{1, 2, 2, 1}, []int{2, 2}), []int{2, 2}) != true {
		return false
	}
	if mytest.IntSliceEqual(intersect([]int{4, 9, 5}, []int{9, 4, 9, 8, 4}), []int{4, 9}) != true {
		return false
	}
	return true
}

func intersect(nums1 []int, nums2 []int) []int {
	count1 := convertToMap(nums1)
	count2 := convertToMap(nums2)

	if len(count1) > len(count2) {
		count1, count2 = count2, count1
	}

	for x := range count1 {
		count1[x] = min(count1[x], count2[x])
	}

	res := []int{}
	for x, n := range count1 {
		for i := 0; i < n; i++ {
			res = append(res, x)
		}
	}

	return res
}

func convertToMap(nums []int) map[int]int {
	res := make(map[int]int, len(nums))

	for i := range nums {
		res[nums[i]]++
	}

	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
