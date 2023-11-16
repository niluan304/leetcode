package main

import "sort"

func threeSum(nums []int) [][]int {
	var (
		counter = make(map[int]int)

		res [][]int
	)

	for _, v := range nums {
		counter[v]++
	}

	for i, v1 := range counter {
		for j, v2 := range counter {
			k := 0 - i - j
			if !(i <= j && j <= k) {
				continue
			}

			if i == j || i == k {
				if v1 < 2 {
					continue
				} else if j == k && v1 < 3 {
					continue
				}
			}

			if j == k && v2 < 2 {
				continue
			}
			if _, ok := counter[k]; ok {
				res = append(res, []int{i, j, k})
			}
		}
	}

	return res
}

// 相向双指针
// 时间复杂度：O(n^2)
// 空间复杂度：O(1)
// 先排序, 然后将问题分解为两数之和
func threeSum2(nums []int) [][]int {
	// 快速排序, 时间复杂度 O(nlogn)
	sort.Ints(nums)

	var (
		ans [][]int
		n   = len(nums)
	)

	// 枚举nums[i], 将问题转化为 nums[i+1:] 中两数之和为 -nums[i]
	for i := 0; i < n-2; i++ {
		target := -nums[i]
		// 去重
		if i > 0 && -target == nums[i-1] {
			continue
		}

		var (
			left  = i + 1
			right = n - 1
		)
		for left < right {
			// 去重
			if left > i+1 && nums[left] == nums[left-1] {
				left++
				continue
			}

			sum := nums[left] + nums[right]
			if sum == target {
				ans = append(ans, []int{nums[i], nums[left], nums[right]})
				left++
				right--
			} else if sum < target {
				left++
			} else {
				right--
			}
		}
	}

	return ans
}
