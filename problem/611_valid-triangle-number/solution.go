package main

import (
	"sort"
)

// 暴力解法
// 时间复杂度：O(n^3)
// 空间复杂度：O(1)
func triangleNumber(nums []int) int {
	// 排序
	sort.Ints(nums)
	var (
		ans = 0
		n   = len(nums)
	)

	for i := 0; i < n-2; i++ {
		for j := i + 1; j < n-1; j++ {
			for k := j + 1; k < n; k++ {
				// 三角形的两边之和大于最长边
				if nums[i]+nums[j] > nums[k] {
					ans++
				}
			}
		}
	}

	return ans
}

// 暴力解法
// 时间复杂度：O(n^3)
// 空间复杂度：O(1)
func triangleNumber2(nums []int) int {
	// 排序
	sort.Ints(nums)
	var (
		ans = 0
		n   = len(nums)
	)

	for i := 0; i < n-2; i++ {
		for j := i + 1; j < n-1; j++ {
			for k := n - 1; k > j; k-- {
				// 三角形的两边之和大于最长边
				if nums[i]+nums[j] > nums[k] {
					ans += k - j
					break
				}
			}
		}
	}

	return ans
}

// 排序 + 双指针
// 时间复杂度：O(n^2)
// 空间复杂度：O(1)
func triangleNumber3(nums []int) int {
	// 排序
	sort.Ints(nums)
	var (
		ans = 0
		n   = len(nums)
	)

	for i := n - 1; i >= 0; i-- {
		var (
			left  = 0
			right = i - 1
		)
		for left < right {
			sum := nums[left] + nums[right]
			if sum <= nums[i] {
				left++
			} else {
				// nums[left], nums[i] 与 nums[left+1:right]的数都满足条件
				ans += right - left
				right--
			}
		}
	}
	return ans
}
