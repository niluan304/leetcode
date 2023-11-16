package main

import (
	"math"
	"sort"
)

// 0 ms 的代码示例
func leetcodeMinTime(nums []int, target int) int {
	n := len(nums)
	sort.Ints(nums)
	minDiff := math.MaxInt
	var ans int
	for i := 0; i < n-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		sum := nums[i] + nums[i+1] + nums[i+2]
		if sum > target {
			if sum-target < minDiff {
				ans = sum
			}
			break
		}

		sum = nums[i] + nums[n-2] + nums[n-1]
		if sum < target {
			if target-sum < minDiff {
				minDiff = target - sum
				ans = sum
			}
			continue
		}

		j, k := i+1, n-1
		for j < k {
			sum = nums[i] + nums[j] + nums[k]
			if sum == target {
				return target
			}
			if sum > target {
				if sum-target < minDiff {
					minDiff = sum - target
					ans = sum
				}
				k--
			} else {
				if target-sum < minDiff {
					minDiff = target - sum
					ans = sum
				}
				j++
			}
		}
	}
	return ans
}

// 2.93 MB 的代码示例
func leetcodeMinMemory(nums []int, target int) int {
	sort.Ints(nums)
	res := nums[0] + nums[1] + nums[2]
	for i := range nums {
		l, r := i+1, len(nums)-1
		for l < r {
			sum := nums[i] + nums[l] + nums[r]
			distance := target - sum
			if distance == 0 {
				return sum
			}
			if distance < 0 {
				r--
				distance = -distance
			} else {
				l++
			}
			if distance < abs(res-target) {
				res = sum
			}
		}
	}
	return res
}
