package main

import "sort"

// 暴力解法
// 时间复杂度：O(n^3)
// 空间复杂度：O(1)
func threeSumSmaller(nums []int, target int) int {
	var (
		n   = len(nums)
		ans = 0
	)

	for i := 0; i < n-2; i++ {
		for j := i + 1; j < n-1; j++ {
			for k := j + 1; k < n; k++ {
				if nums[i]+nums[j]+nums[k] < target {
					ans++
				}
			}
		}
	}

	return ans
}

// 相向双指针
// 时间复杂度：O(n^2)
// 空间复杂度：O(1)
func threeSumSmaller2(nums []int, target int) int {
	// 需要先排序
	sort.Ints(nums)

	var (
		n   = len(nums)
		ans = 0
	)

	for i := n - 1; i >= 0; i-- {
		var (
			left  = 0
			right = i - 1
			over  = target - nums[i]
		)

		// 相向双指针
		for left < right {
			sum := nums[left] + nums[right]
			if sum >= over {
				right--
			} else {
				// nums[i] + nums[left] + nums[k] < target,  其中 k 属于 [left+1, right]
				ans += right - left
				left++
			}
		}
	}

	return ans
}
