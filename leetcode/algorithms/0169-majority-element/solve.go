package main

import "fmt"

func main() {
	fmt.Println(test())
}

func test() bool {
	if majorityElement([]int{3, 2, 3}) != 3 {
		return false
	}
	return true
}

func majorityElement(nums []int) int {
    return find(nums, 0, len(nums) - 1)
}

func find(nums []int, left int, right int) int {
	if left == right {
		return nums[left]
	}
	mid := (left + right) / 2

	left_ans := find(nums, left, mid)
	right_ans := find(nums, mid + 1, right)

	if left_ans == right_ans {
		return left_ans
	} else {
		left_count := 0
		right_count := 0
		for i := left; i <= right; i++ {
			if nums[i] == left_ans {
				left_count = left_count + 1
			} else if nums[i] == right_ans {
				right_count = right_count + 1
			}
		}
		if left_count > right_count {
			return left_ans
		} else {
			return right_ans
		}
	}
}