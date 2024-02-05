package main

import (
	"math"

	. "github.com/niluan304/leetcode/copypasta/dp"
)

// DP
//
// 思路类似 LIS
//
// # 方法一：O(n^2) 划分型动态规划
//
// ### 如何思考
//
// 划分出**第一个**子数组，问题变成一个规模更小的子问题。
//
// 由于「划分出长为 $x$ 和 $y$ 的子数组」和「划分出长为 $y$ 和 $x$ 的子数组」之后，剩余的子问题是相同的，因此这题适合用动态规划解决。
//
// ### 复杂度分析
//
// - 时间复杂度：$O(n^2)$，其中 $n$ 为 $\textit{nums}$ 的长度。
// - 空间复杂度：$O(n^2)$。
// TODO 使用线段树优化
// [动态规划+线段树优化](https://leetcode.cn/problems/minimum-cost-to-split-an-array/solutions/2072725/by-endlesscheng-05s0)
func minCost(nums []int, k int) int {
	var ans = 0
	cost := trimmed(nums) // 预处理代价
	var dfs func(i int) int
	dfs = func(i int) int {
		if i == 0 {
			return 0
		}
		res := math.MaxInt32
		for j := 0; j < i; j++ {
			res = min(res, dfs(j)+cost[j][i-1]+k)
		}
		return res
	}

	cache := MemorySearch(&dfs)
	_ = cache

	ans = dfs(len(nums))
	return ans
}

func trimmed(nums []int) [][]int {
	n := len(nums)
	cost := make([][]int, n)

	for i := 0; i < n; i++ {
		cost[i] = make([]int, n)

		count := map[int]int{}
		mul := 0
		for j, num := range nums[i:] {
			count[num]++
			if count[num] == 2 {
				mul++
			}
			cost[i][i+j] = j + 1 - (len(count) - mul)
		}
	}
	return cost
}
