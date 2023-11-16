package main

import "sort"

func moveZeroes(nums []int) {
	var count = 0
	for _, num := range nums {
		if num == 0 {
			count++
		}
	}

	sort.SliceStable(nums, func(i, j int) bool {
		if nums[i] == 0 {
			return false
		}
		if nums[j] == 0 {
			return true
		}

		return i < j
	})

	for i := 0; i < count; i++ {
		nums[len(nums)-1-i] = 0
	}
}

func moveZeroes2(nums []int) {
	var count = 0
	for i, num := range nums {
		// 可优化为 if num != 0 { nums[count] = num; count++ }
		if num == 0 {
			count++
		} else {
			nums[i-count] = num
		}
	}

	for i := 0; i < count; i++ {
		nums[len(nums)-1-i] = 0
	}
}
