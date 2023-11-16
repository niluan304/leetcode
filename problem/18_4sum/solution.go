package main

import "sort"

// 相向双指针
// 时间复杂度：O(n^3)
// 空间复杂度：O(1)
// 先排序，然后降阶为三数之和，再降阶为两数之和
func fourSum(nums []int, target int) [][]int {
	// 快速排序, 时间复杂度 O(nlogn)
	sort.Ints(nums)

	var (
		ans [][]int
		n   = len(nums)
	)

	for i := 0; i < n-3; i++ {
		// 去重
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		for j := i + 1; j < n-2; j++ {
			// 去重
			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}

			var (
				left  = j + 1
				right = n - 1
			)

			for left < right {
				sum := nums[i] + nums[j] + nums[left] + nums[right]
				if sum == target {
					ans = append(ans, []int{nums[i], nums[j], nums[left], nums[right]})
					left++
					// 跳过与左边界相同的元素
					for left < right && nums[left] == nums[left-1] {
						left++
					}

					right--
					// 跳过与右边界相同的元素
					for left < right && nums[right] == nums[right+1] {
						right--
					}
				} else if sum < target {
					left++
				} else {
					right--
				}
			}
		}
	}

	return ans
}
