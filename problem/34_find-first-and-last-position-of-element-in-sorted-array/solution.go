package main

import "sort"

// 暴力解法
// 时间复杂度：O(n)
// 空间复杂度：O(1)
func searchRange(nums []int, target int) []int {
	var (
		left  = -1
		right = -1
	)

	for i, num := range nums {
		if num != target {
			if num > target {
				break
			}
			continue
		}

		if left == -1 {
			left = i
			right = i
		} else {
			right = i
		}
	}

	return []int{left, right}
}

// 二分查找
// 时间复杂度：O(log n)
// 空间复杂度：O(1)
func searchRange2(nums []int, target int) []int {
	var (
		left  = 0
		right = len(nums) - 1
	)

	for left <= right {
		if nums[left] == target && nums[right] == target {
			return []int{left, right}
		}

		mid := (right-left)/2 + left

		if nums[mid] == target {
			if nums[left] < target {
				left++
			}
			if nums[right] > target {
				right--
			}
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return []int{-1, -1}
}

// 二分查找
// 时间复杂度：O(log n)
// 空间复杂度：O(1)
func searchRange3(nums []int, target int) []int {
	left := binarySearch(nums, target)

	if left == len(nums) || // nums 中所有数都小于 target
		nums[left] != target { // nums 未找到 target
		return []int{-1, -1}
	}

	right := binarySearch(nums, target+1)
	return []int{left, right - 1}
}

// 二分查找
// 在数组中寻找第一个小于等于 target的数, 返回其索引
// 如果全都小于 target, 返回数组长度
func binarySearch(nums []int, target int) int {
	var (
		// 闭区间 [left, right]
		left  = 0
		right = len(nums) - 1
	)

	for left <= right { // 区间不为空
		mid := (right-left)/2 + left

		if nums[mid] < target {
			left = mid + 1 // 范围缩小到 [mid+1, right]
		} else {
			right = mid - 1 // 范围缩小到 [left, mid-1]
		}
	}
	return left
}

// 二分查找
// 时间复杂度：O(log n)
// 空间复杂度：O(1)
func searchRange4(nums []int, target int) []int {
	// 二分查找可以找到第一个 i, 使得 nums[i] <= target
	// 但如果 i == len(nums)，说明 max(nums) < target
	// 此外也不保证 nums[i] == target
	left := sort.SearchInts(nums, target)

	if left == len(nums) || // nums 中所有数都小于 target
		nums[left] != target { // nums 未找到 target
		return []int{-1, -1}
	}

	right := sort.SearchInts(nums, target+1)
	return []int{left, right - 1}
}
