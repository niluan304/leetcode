package main

import (
	"math"
)

// 暴力解法, 遍历数组找出最小值
// 时间复杂度：O(n)
// 空间复杂度：O(1)
func findMin(nums []int) int {
	min := math.MaxInt

	for _, num := range nums {
		if num < min {
			min = num
		}
	}

	return min
}

// 二分查找
// 时间复杂度：O(logn)
// 空间复杂度：O(1)
func findMin2(nums []int) int {
	// nums[0] < nums[n-1], 说明数组已复位, 则直接返回 nums[0]
	if nums[0] < nums[len(nums)-1] {
		return nums[0]
	}

	var (
		left  = 0
		right = len(nums) - 1
	)
	// 若旋转过, 则一定有 nums[left] > nums[right]
	for left != right-1 {
		mid := left + (right-left)/2
		if nums[mid] < nums[right] {
			right = mid // 缩减范围至 [left, mid]
		} else {
			left = mid // 缩减范围至 [mid, right]
		}
	}
	return nums[right]
}

// 二分查找
// 时间复杂度：O(logn)
// 空间复杂度：O(1)
func findMin3(nums []int) int {
	var (
		left  = 0
		right = len(nums) - 1
	)

	// 原数组为升序数组
	// 设 x 为旋转次数, 则有:
	// 在 [0, x-1] 单调递减
	// 在 [x, n-1] 单调递增
	// 且 [0, x-1] 的数都大于 [x, n-1] 的数
	for nums[left] > nums[right] {
		mid := left + (right-left)/2

		if nums[mid] < nums[right] {
			// 在 [x, n-1] 单调递增, nums[mid] < nums[right], 说明 mid 更接近 x
			// 缩减范围至 [left, mid]
			right = mid
		} else {
			// nums[mid] > nums[right] 说明 mid 不在 [x, n-1]
			// 缩减范围至 [mid+1, right],
			left = mid + 1
		}
	}
	return nums[left]
}

// 二分查找
// 时间复杂度：O(logn)
// 空间复杂度：O(1)
func findMin4(nums []int) int {
	var (
		last  = nums[len(nums)-1]
		left  = -1
		right = len(nums) - 1
	)
	// [0, n-2]
	// (-1, n-1)

	for left < right-1 {
		mid := left + (right-left)/2
		if nums[mid] < last {
			// 说明最小值在 mid 左侧
			// 缩减范围至 [left, mid]
			right = mid
		} else {
			// 说明最小值在 mid 右侧
			// 缩减范围至 [mid, right]
			left = mid
		}
	}
	return nums[right]
}

// 二分查找
// 时间复杂度：O(logn)
// 空间复杂度：O(1)
func findMin5(nums []int) int {
	var (
		left  = 0
		right = len(nums) - 1
	)

	for left < right {
		mid := left + (right-left)/2
		if nums[mid] < nums[right] {
			// 说明最小值在 mid 左侧
			// 缩减范围至 [left, mid]
			right = mid
		} else {
			// 说明最小值在 mid 右侧
			// 缩减范围至 [mid+1, right]
			left = mid + 1
		}
	}
	return nums[right]
}
