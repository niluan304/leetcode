package main

import (
	"math"

	. "github.com/niluan304/leetcode/copypasta/dp"
)

// nums = [1, 4, 9 ..., x]，x 为小于等于 n 的最大完全平方数
// 现在有 N 件物品，花费为 nums，价值都为 1，求恰好装满容量为 n 的背包时，价值值的最小值
//
// dfs + 记忆化搜索
// - 时间复杂度：$\mathcal{O}(n \cdot m)$。
// - 空间复杂度：$\mathcal{O}(n \cdot m)$。
func numSquares(n int) int {
	var ans = 0 // math.MaxInt32 // math.MinInt32
	var nums []int
	for i := 1; i <= n; i++ {
		x := i * i
		if x > n {
			break
		}
		nums = append(nums, x)
	}

	var dfs func(i int, j int) int
	dfs = func(i int, j int) int {
		if i == 0 {
			if j == 0 {
				return 0
			}

			return math.MaxInt32
		}

		v := dfs(i-1, j)
		if k := j - nums[i-1]; k >= 0 {
			v = min(v, dfs(i, k)+1)
		}
		return v
	}

	MemorySearch2(&dfs)

	ans = dfs(len(nums), n)
	return ans
}

// nums = [1, 4, 9 ..., x]，x 为小于等于 n 的最大值
// 现在有 N 件物品，花费为 nums，价值都为 1，求恰好装满容量为 n 的背包时，价值值的最小值
//
// 完全背包
// - 时间复杂度：$\mathcal{O}(n \cdot m)$。
// - 空间复杂度：$\mathcal{O}(n)$。
func numSquares2(n int) int {
	var nums, values []int
	for i := 1; i <= n; i++ {
		x := i * i
		if x > n {
			break
		}
		nums = append(nums, x)
		values = append(values, 1)
	}

	var dp = make([]int, n+1)
	for i, _ := range dp {
		dp[i] = math.MaxInt32
	}
	dp[0] = 0

	for i, num := range nums {
		for j := num; j <= n; j++ {
			dp[j] = min(dp[j], dp[j-num]+values[i])
		}
	}
	return dp[n]
}
