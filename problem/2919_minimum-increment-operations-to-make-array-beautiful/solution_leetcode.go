package main

import . "github.com/niluan304/leetcode/copypasta/dp"

// [视频讲解](https://www.bilibili.com/video/BV1tw411q7VZ/) 第三题。
//
// ## 前置知识
//
// 请看视频讲解 [动态规划入门：从记忆化搜索到递推](https://b23.tv/72onpYq)
//
// ## 写法一：记忆化搜索
//
// 把大于 $k$ 的元素视作 $k$。
//
// 由于大于 $3$ 的子数组必然包含等于 $3$ 的子数组，问题转换成：每个长为 $3$ 的子数组都需要包含至少一个 $k$。
//
// 考虑最后一个元素「选或不选」，即是否增大：
//
// - 增大到 $k$：那么对于左边那个数来说，它右边就有一个 $k$ 了。
// - 不增大：那么对于左边那个数来说，它右边有一个没有增大的数。
//
// 进一步地，如果倒数第二个数也不增大，那么对于倒数第三个数，它右边就有两个没有增大的数，那么它一定要增大（不用管右边那两个数是否为 $k$，因为下面的 $\textit{dfs}$ 会考虑到所有的情况，取最小值）。
//
// 因此，用 $i$ 表示当前位置，$j$ 表示右边有几个没有增大的数。定义 $\textit{dfs}(i,j)$ 表示在该状态下对于前 $i$ 个数的最小递增运算数。
//
// 考虑 $\textit{nums}[i]$ 是否增大到 $k$：
//
// - 增大，即 $\textit{dfs}(i-1,0) + \max(k-\textit{nums}[i], 0)$。
// - 如果 $j<2$，则可以不增大，即 $\textit{dfs}(i-1,j+1)$。
//
// 这两种情况取最小值。
//
// 递归边界：当 $i<0$ 时返回 $0$。
//
// 递归入口：$\textit{dfs}(n-1,0)$，即答案。
// ### 复杂度分析
//
// - 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 为 $\textit{nums}$ 的长度。
// - 空间复杂度：$\mathcal{O}(1)$。
func endlessCheng(nums []int, k int) int64 {
	n := len(nums)
	var dfs func(int, int) int
	dfs = func(i, j int) int {
		if i < 0 {
			return 0
		}

		res := dfs(i-1, 0) + max(k-nums[i], 0) // nums[i] 增大
		if j < 2 {
			res = min(res, dfs(i-1, j+1)) // nums[i] 不增大
		}
		return res
	}

	MemorySearch2(&dfs)

	return int64(dfs(n-1, 0))
}
