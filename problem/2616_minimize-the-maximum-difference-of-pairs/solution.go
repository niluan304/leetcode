package main

import "sort"

// 排序后贪心 + 二分答案
// 时间复杂度: O(nlogn)
// 空间复杂度: O(1)
func minimizeMax(nums []int, p int) int {
	var n = len(nums)
	sort.Ints(nums)

	ans := sort.Search(nums[n-1]-nums[0], func(maxD int) bool {
		var cnt = 0
		for i := 0; i < n-1; i++ {
			if nums[i]+maxD >= nums[i+1] {
				cnt++
				i++
			}
		}
		return cnt >= p
	})
	return ans
}
