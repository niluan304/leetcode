package main

import "sort"

// 28 ms 的代码示例
func leetcodeMinTime(nums []int) [][]int {
	sort.Ints(nums)

	var result [][]int
	size := len(nums)
	for i := 0; i < size-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		if nums[i] > 0 {
			break
		}
		if nums[i]+nums[size-2]+nums[size-1] < 0 {
			continue
		}
		j, k := i+1, size-1
		for j < k {
			sum := nums[i] + nums[j] + nums[k]
			if sum < 0 {
				j++
			} else if sum > 0 {
				k--
			} else {
				result = append(result, []int{nums[i], nums[j], nums[k]})
				j++
				for j < k && nums[j] == nums[j-1] {
					j++
				}
				k--
				for j < k && nums[k] == nums[k+1] {
					k--
				}
			}
		}
	}
	return result
}
