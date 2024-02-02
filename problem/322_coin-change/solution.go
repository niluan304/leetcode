package main

import (
	"math"

	. "github.com/niluan304/leetcode/copypasta/dp"
)

// dfs + 记忆化搜索
// - 时间复杂度：$\mathcal{O}(n \cdot m)$。
// - 空间复杂度：$\mathcal{O}(n \cdot m)$。
func coinChange(coins []int, amount int) int {
	var dfs func(i, j int) int
	dfs = func(i, j int) int {
		if i == 0 {
			if j == 0 {
				return 0
			}
			return math.MaxInt32 // 无法达成的方案
		}

		v := dfs(i-1, j) // 不选 coins[i] 硬币
		if k := j - coins[i-1]; k >= 0 {
			v = min(v, dfs(i, k)+1) // 选择 coins[i] 硬币
		}
		return v
	}

	MemorySearch2(&dfs)
	ans := dfs(len(coins), amount)
	if ans >= math.MaxInt32 {
		return -1
	}
	return ans
}

// 将 dfs + 记忆化搜索 1:1　翻译为递推
//
// - 时间复杂度：$\mathcal{O}(n \cdot m)$。
// - 空间复杂度：$\mathcal{O}(n \cdot m)$。
//
// 完全背包解释：
// 有容量为 amount 的背包，你有数量无限的 N 种物品，价值都是 1，费用为 coins，求恰好装满背包时的最小价值和。
func coinChange2(coins []int, amount int) int {
	var n = len(coins)
	var dp = make([][]int, n+1)
	for i, _ := range dp {
		dp[i] = make([]int, amount+1)
	}
	for i, _ := range dp[0] {
		dp[0][i] = math.MaxInt32
	}
	dp[0][0] = 0

	for i, coin := range coins {
		for j := 0; j <= amount; j++ {
			dp[i+1][j] = dp[i][j]

			if j >= coin {
				dp[i+1][j] = min(dp[i+1][j], dp[i+1][j-coin]+1)
			}
		}
	}
	var ans = dp[n][amount]
	if ans >= math.MaxInt32 {
		return -1
	}
	return ans
}

// coinChange2 的空间优化
// - 时间复杂度：$\mathcal{O}(n \cdot m)$。
// - 空间复杂度：$\mathcal{O}(m)$。
func coinChange3(coins []int, amount int) int {
	var dp = make([]int, amount+1)
	for i, _ := range dp {
		dp[i] = math.MaxInt32
	}
	dp[0] = 0

	for _, coin := range coins {
		for j := coin; j <= amount; j++ {
			dp[j] = min(dp[j], dp[j-coin]+1)
		}
	}
	var ans = dp[amount]
	if ans >= math.MaxInt32 {
		return -1
	}
	return ans
}

func coinChange4(coins []int, amount int) int {
	values := make([]int, len(coins))
	for i, _ := range coins {
		values[i] = 1
	}
	return unboundedKnapsackExactlyFull(values, coins, amount)
}

// 有数量无限的 N 种物品，价值为 values, 容量消耗为 weights。
// 求恰好装满容量为 maxW 的背包时，其最小价值和。
func unboundedKnapsackExactlyFull(values, weights []int, maxW int) int {
	n := len(values)
	dp := make([][]int, n+1) // int64  fill
	for i, _ := range dp {
		dp[i] = make([]int, maxW+1)
	}
	for i, _ := range dp[0] {
		dp[0][i] = math.MaxInt32
	}

	dp[0][0] = 0
	for i, value := range values {
		w := weights[i]
		for j := 0; j <= maxW; j++ {
			dp[i+1][j] = dp[i][j]
			if j >= w {
				dp[i+1][j] = min(dp[i+1][j], dp[i+1][j-w]+value)
			}
		}
	}
	ans := dp[n][maxW]
	if ans >= math.MaxInt32 { // 无法恰好装满背包
		return -1
	}
	return ans
}
