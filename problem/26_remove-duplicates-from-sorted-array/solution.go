package main

import "math"

func removeDuplicates(nums []int) int {
	idx := 1
	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[i-1] {
			nums[idx] = nums[i]
			idx++
		}
	}
	return idx
}

func removeDuplicates2(nums []int) int {
	x := math.MinInt
	idx := 0

	for _, num := range nums {
		if num != x {
			nums[idx] = num
			idx++
		}
		x = num
	}
	return idx
}
