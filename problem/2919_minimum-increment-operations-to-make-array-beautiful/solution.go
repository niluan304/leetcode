package main

import (
	"math"
	"slices"

	. "github.com/niluan304/leetcode/copypasta/dp"
)

const Len = 3 // 题目要求：如果数组中任何长度 大于或等于 3 的子数组，因此 Len = 3

// dfs + 记忆化搜索
//
// - 时间复杂度：$\mathcal{O}(n \cdot Σ)$。$Σ = 3$。
// - 空间复杂度：$\mathcal{O}(n)$。
func minIncrementOperations(nums []int, k int) int64 {
	var dfs func(i int) int
	dfs = func(i int) (res int) {
		if i < 0 {
			return 0
		}

		// 如果把当前数字调整为不小于 k 的数需要的操作数量为 max(0, k - num)，而此前需要的操作数量为 min(dp1, dp2, dp3)
		// 原来距离为 1 变为距离为 2
		// 原来距离为 2 变为距离为 3
		return max(0, k-nums[i]) + min(dfs(i-1), dfs(i-2), dfs(i-3))
	}

	cache := MemorySearch(&dfs)
	_ = cache

	n := len(nums)
	ans := min(dfs(n-1), dfs(n-2), dfs(n-3))

	return int64(ans)
}

// DP
//
// 将 minIncrementOperations 1:1 翻译为递推
//
// - 时间复杂度：$\mathcal{O}(n \cdot Σ)$。$Σ = 3$。
// - 空间复杂度：$\mathcal{O}(n)$。
func minIncrementOperations2(nums []int, k int) int64 {
	var n = len(nums)
	var dp = make([]int, n+Len)
	for i, num := range nums {
		// 如果把当前数字调整为不小于 k 的数需要的操作数量为 max(0, k - num)，而此前需要的操作数量为 min(dp1, dp2, dp3)
		// 原来距离为 1 变为距离为 2
		// 原来距离为 2 变为距离为 3

		dp[i+Len] = math.MaxInt32
		for j := 0; j < Len; j++ {
			dp[i+Len] = min(dp[i+Len], dp[i+j])
		}
		dp[i+Len] += max(k-num, 0)
	}
	return int64(slices.Min(dp[n:]))
}

// 选与不选
//
// - 时间复杂度：$\mathcal{O}(n \cdot Σ)$。$Σ = 3$。
// - 空间复杂度：$\mathcal{O}(n \cdot Σ)$。$Σ = 3$。
func minIncrementOperations3(nums []int, k int) int64 {
	var n = len(nums)

	// dfs(i ,j) 表示 nums[i] 的右边最近的大于等于 k 的数的距离，1<= j <= 3
	var dfs func(i, j int) int
	dfs = func(i, j int) (res int) {
		if i < 0 {
			return 0
		}

		res = max(0, k-nums[i]) + dfs(i-1, 1)
		if j == Len {
			res = res // 距离达到最大值 3，只能选择 nums[i] 变成 k
		} else {
			res = min( // 距离未达到最大值 3，可以选与不选
				res,           // 选择 nums[i] 变成 k
				dfs(i-1, j+1), // 不选择 nums[i] 变成 k
			)
		}
		return res
	}

	cache := MemorySearch2(&dfs)
	_ = cache

	ans := dfs(n-1, 1)
	return int64(ans)
}
