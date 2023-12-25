package main

import (
	"math"
)

// 暴力穷举
// 时间复杂度: O(n^2)
// 空间复杂度: O(1)
// Deprecated: 超时
func minSubArrayLen(target int, nums []int) int {

	n := len(nums)
	for i := 0; i < n; i++ {
		sum := 0
		for _, num := range nums[:i+1] {
			sum += num
		}
		if sum >= target {
			return i + 1
		}

		for j := i + 1; j < n; j++ {
			sum += nums[j] - nums[j-i-1]
			if sum >= target {
				return i + 1
			}
		}
	}

	return 0
}

func _min(x, y int) int {
	if x > y {
		return y
	}

	return x
}

// 滑动窗口
// 时间复杂度: O(n)
// 空间复杂度: O(1)
func minSubArrayLen2(target int, nums []int) int {
	sum := 0
	count := math.MaxInt
	left := 0

	for right, num := range nums {
		sum += num
		for sum >= target {
			sum -= nums[left]
			count = _min(count, right-left)
			left++
		}
	}

	if count == math.MaxInt {
		return 0
	}
	return count
}
