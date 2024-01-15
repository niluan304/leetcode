package main

import "sort"

func maximumBeauty(nums []int, k int) int {
	sort.Ints(nums)

	var ans = 1
	var left = 0
	for right, num := range nums {
		for num-nums[left] > 2*k {
			left++
		}

		ans = max(ans, right-left+1)
	}
	return ans
}
