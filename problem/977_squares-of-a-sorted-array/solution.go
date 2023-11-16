package main

import (
	"sort"
)

func sortedSquares(nums []int) []int {
	for i, num := range nums {
		nums[i] = num * num
	}

	sort.Ints(nums)
	return nums
}

func sortedSquares2(nums []int) []int {
	for i, num := range nums {
		nums[i] = num * num
	}

	var i, j = 0, len(nums) - 1
	var ans = make([]int, len(nums))

	for k := len(nums) - 1; k >= 0; k-- {
		if nums[i] > nums[j] {
			ans[k] = nums[i]
			i++
		} else {
			ans[k] = nums[j]
			j--
		}
	}

	return ans
}

func sortedSquares3(nums []int) []int {
	var i, j = 0, len(nums) - 1
	var ans = make([]int, len(nums))
	for k := len(nums) - 1; k >= 0; k-- {
		var vi, vj = nums[i], nums[j]
		if vi+vj < 0 {
			ans[k] = vi * vi
			i++
		} else {
			ans[k] = vj * vj
			j--
		}
	}

	return ans
}
