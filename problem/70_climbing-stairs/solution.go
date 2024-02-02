package main

import (
	"slices"

	. "github.com/niluan304/leetcode/copypasta/dp"
)

func climbStairs(n int) int {
	var dp = make([]int, n+2)
	dp[0], dp[1] = 0, 1 // 设置 dp[-1], dp[0]，避免数组越界，增加向右的偏移量

	for i := 0; i < n; i++ {
		dp[i+2] = dp[i+1] + dp[i]
	}
	return dp[n+1]
}

func climbStairs2(n int) int {
	return waysToTarget(n, []int{1, 2})
}

// 爬楼梯
// 需要 target 阶你才能到达楼顶。你有 ways[:] 方式向上爬
// 你从 0 开始出发，有多少种方式爬到楼顶？
//
// - 时间复杂度：$\mathcal{O}(n \codt m)$。
// - 空间复杂度：$\mathcal{O}(n \codt m)$。
//
// 题目保证 ways[i] >= 1
// 但如果 ways[i] 可以为负数，该怎么处理？
func waysToTarget(target int, ways []int) int {
	slices.Sort(ways[:])

	var dfs func(n int) int
	dfs = func(n int) int {
		if n == 0 {
			return 1
		}

		res := 0
		for _, way := range ways {
			if way > n {
				break
			}
			res += dfs(n - way)
		}
		return res
	}

	MemorySearch(&dfs)
	ans := dfs(target)
	return ans
}
