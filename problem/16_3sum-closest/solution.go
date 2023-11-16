package main

import (
	"math"
	"sort"
)

// 暴力解法
// 时间复杂度：O(n^3)
// 空间复杂度：O(1)
func threeSumClosest(nums []int, target int) int {
	var (
		closest = math.MaxInt
		sum     = 0
	)
	for i := range nums {
		for j := i + 1; j < len(nums)-1; j++ {
			for k := j + 1; k < len(nums); k++ {
				_sum := nums[i] + nums[j] + nums[k]
				_closest := abs(_sum - target)

				if _closest < closest {
					closest = _closest
					sum = _sum
				}
			}
		}
	}

	return sum
}

func abs(n int) int {
	if n < 0 {
		return -n
	}

	return n
}

// 相向双指针
// 时间复杂度：O(n^2)
// 空间复杂度：O(1)
func threeSumClosest2(nums []int, target int) int {
	var (
		closest = math.MaxInt

		n = len(nums)
	)

	// 相向双指针需要先排序
	sort.Ints(nums)

	for i := 0; i < n-2; i++ {
		var (
			left  = i + 1
			right = n - 1
		)

		for left < right {
			sum := nums[i] + nums[left] + nums[right]
			if sum == target {
				return target
			}

			// 更新最接近的值
			if abs(sum-target) < abs(closest-target) {
				closest = sum
			}

			// 移动指针, 使得 sum 更接近 target
			if sum < target {
				left++
			} else {
				right--
			}
		}
	}

	return closest
}

// 相向双指针
// 时间复杂度：O(n^2)
// 空间复杂度：O(1)
// 优化：去重
func threeSumClosest3(nums []int, target int) int {
	var (
		closest = math.MaxInt

		n = len(nums)
	)

	// 相向双指针需要先排序
	sort.Ints(nums)

	for i := 0; i < n-2; i++ {
		var (
			left  = i + 1
			right = n - 1
		)

		// 去重，如果当前值和上一个值相同，则跳过
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		for left < right {
			sum := nums[i] + nums[left] + nums[right]
			if sum == target {
				return target
			}

			// 更新最接近的值
			if abs(sum-target) < abs(closest-target) {
				closest = sum
			}

			// 移动指针, 使得 sum 更接近 target
			if sum < target {
				left++
			} else {
				right--
			}
		}
	}

	return closest
}
