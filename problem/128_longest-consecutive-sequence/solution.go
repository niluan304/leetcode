package main

import (
	"sort"
)

// 排序
// - 时间复杂度：$\mathcal{O}(n\log n)$。
// - 空间复杂度：$\mathcal{O}(1)$。
// Deprecated: 超时
func sortInt(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	sort.Ints(nums)

	var mx = 1
	var cur = 1

	for i := 1; i < len(nums); i++ {
		switch nums[i] - nums[i-1] {
		case 0:
			continue
		case 1:
			cur++
		default:
			cur = 1
		}

		if mx < cur {
			mx = cur
		}
	}

	return mx
}

// 1e-9 <= nums[i] <= 1e9
// len(nums) <= 1e5
// 哈希表
//
// - 时间复杂度：$\mathcal{O}(n)$。
// - 空间复杂度：$\mathcal{O}(n)$。
func longestConsecutive(nums []int) int {
	set := map[int]bool{}
	for _, num := range nums {
		set[num] = true
	}
	ans := 0
	for num, _ := range set {
		if set[num-1] {
			continue
		}
		count := 1
		for set[num+count] {
			count++
		}
		ans = max(ans, count)
	}
	return ans
}
