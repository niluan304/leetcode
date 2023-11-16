package main

// 暴力解法
// 时间复杂度：O(n)
// 空间复杂度：O(1)
func findPeakElement(nums []int) int {
	var n = len(nums)

	var peakLeft = true
	for i := 0; i < n-1; i++ {
		peakRight := nums[i] > nums[i+1]
		if peakLeft && peakRight {
			return i
		}
		peakLeft = !peakRight
	}
	return n - 1
}

// 二分查找
// 时间复杂度：O(logn)
// 空间复杂度：O(1)
func findPeakElement2(nums []int) int {
	var (
		left  = 0
		right = len(nums) - 1
	)

	for left < right {
		mid := left + (right-left)/2

		// 题目保证 nums[i] != nums[i + 1]
		if nums[mid] > nums[mid+1] {
			right = mid // [left, mid] 中有峰值
		} else {
			left = mid + 1 // [mid+1, right] 中有峰值
		}
	}

	return left
}
