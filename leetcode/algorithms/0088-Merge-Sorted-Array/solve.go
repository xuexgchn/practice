package main

import "fmt"

func main() {
	fmt.Println(test())
}

func test() bool {
	merge([]int{1, 2, 3, 0, 0, 0}, 3, []int{2, 5, 6}, 3)

	return true
}

// merge from the tail
func merge(nums1 []int, m int, nums2 []int, n int) {
	bingoIndex := len(nums1) - 1

	for bingoIndex > 0 {
		if nums1[bingoIndex] != nums1[bingoIndex-1] {
			break
		}
		bingoIndex--
	}

	bingoIndex--
	index2 := len(nums2) - 1
	for i := bingoIndex + index2 + 1; i >= 0; i-- {
		if bingoIndex < 0 {
			nums1[i] = nums2[index2]
			index2--
		} else if index2 < 0 {
			nums1[i] = nums1[bingoIndex]
			bingoIndex--
		} else if nums1[bingoIndex] >= nums2[index2] {
			nums1[i] = nums1[bingoIndex]
			bingoIndex--
		} else {
			nums1[i] = nums2[index2]
			index2--
		}
	}
}
