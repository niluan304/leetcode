package main

import (
	"math"
)

// dfs + 记忆化搜索
//
// - 时间复杂度：$\mathcal{O}(n \cdot target)$。
// - 空间复杂度：$\mathcal{O}(n \cdot target)$。
func lengthOfLongestSubsequence(nums []int, target int) int {
	var n = len(nums)
	var cache = make([][]*int, n+1)
	for i, _ := range cache {
		cache[i] = make([]*int, target+1)
	}

	var dfs func(i, target int) int
	dfs = func(i, target int) (res int) {
		if i == 0 {
			if target == 0 {
				return 0
			}
			return math.MinInt
		}

		ptr := &cache[i][target]
		if *ptr != nil {
			return **ptr
		}
		defer func() { *ptr = &res }()

		res = dfs(i-1, target)
		if num := nums[i-1]; num <= target {
			res = max(res, dfs(i-1, target-num)+1)
		}

		return res
	}

	if ans := dfs(n, target); ans > 0 {
		return ans
	}
	return -1
}

// dp
// 将 dfs + 记忆化搜索 1:1 翻译为 dp
//
// - 时间复杂度：$\mathcal{O}(n \cdot target)$。
// - 空间复杂度：$\mathcal{O}(n \cdot target)$。
func lengthOfLongestSubsequence2(nums []int, target int) int {
	var n = len(nums)
	var dp = make([][]int, n+1)
	for i, _ := range dp {
		dp[i] = make([]int, target+1)
	}
	for i := 1; i <= target; i++ {
		dp[0][i] = math.MinInt
	}

	for i, num := range nums {
		for j := 0; j <= target; j++ {
			dp[i+1][j] = dp[i][j]
			if j >= num {
				dp[i+1][j] = max(dp[i+1][j], dp[i][j-num]+1)
			}
		}
	}

	if ans := dp[n][target]; ans > 0 {
		return ans
	}
	return -1
}
