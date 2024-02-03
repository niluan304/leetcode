package main

import "slices"

// LIS + 前后缀分解
// - 时间复杂度：$\mathcal{O}(n\log n)$。
// - 空间复杂度：$\mathcal{O}(n)$。
func minimumMountainRemovals(nums []int) int {
	var n = len(nums)
	var ans = 0 // math.MaxInt32 // math.MinInt32

	prefix, suffix := make([]int, n), make([]int, n)

	var greedy []int
	for i := 0; i < n; i++ {
		num := nums[i]
		j, _ := slices.BinarySearch(greedy, num) // 改成 num+1 为非严格递增
		if j == len(greedy) {
			greedy = append(greedy, num)
		} else {
			greedy[j] = num
		}
		prefix[i] = j
	}

	greedy = []int{}
	for i := n - 1; i >= 0; i-- {
		num := nums[i]
		j, _ := slices.BinarySearch(greedy, num) // 改成 num+1 为非严格递增
		if j == len(greedy) {
			greedy = append(greedy, num)
		} else {
			greedy[j] = num
		}
		suffix[i] = j
	}

	for i := 0; i < n; i++ {
		if prefix[i] == 0 || suffix[i] == 0 {
			continue
		}
		ans = max(ans, prefix[i]+suffix[i]+1)
	}

	return n - ans
}
