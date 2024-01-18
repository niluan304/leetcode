package main

import (
	"slices"
)

func countSubarrays(nums []int, k int) int64 {
	var n = len(nums)
	var ans = 0
	var left, cnt = 0, 0
	var mx = slices.Max(nums)

	for right, num := range nums {
		if num == mx {
			cnt++
		}
		for cnt == k {
			ans += n - right
			if nums[left] == mx {
				cnt--
			}
			left++
		}
	}
	return int64(ans)
}
