package main

import (
	"math"
	"sort"
)

// dfs + 记忆化搜索
// 时间复杂度：O(n^2)
// 空间复杂度：O(n)
func coinChange(coins []int, amount int) int {
	sort.Ints(coins)

	// Empty 如果为 math.MaxInt,由于 dfs(n)+1, 会导致溢出为 math.MinInt
	// 因此设置个非常大的数，  0 <= amount <= 10^4
	const Empty = 1_0000_0000
	var cache = make([]int, amount+1)

	var dfs func(target int) int
	dfs = func(target int) int {
		//// 已经预处理了 coin > target 的情况, 不需要额外判断 target < 0
		//if target < 0 {
		//	return Empty
		//}
		if target == 0 {
			return 0
		}

		v := cache[target]
		if v > 0 {
			return v
		}

		v = Empty
		idx := sort.Search(len(coins), func(j int) bool {
			return coins[j] > target
		})
		for _, coin := range coins[:idx] {
			// 逐个替换硬币
			v = _min(v, dfs(target-coin)+1)
		}

		cache[target] = v
		return v
	}
	var ans = dfs(amount)
	if ans == Empty {
		return -1
	}
	return ans
}

func _min(list ...int) int {
	var ans = math.MaxInt
	for _, n := range list {
		if ans > n {
			ans = n
		}
	}

	return ans
}

// dp 动态规划
// 时间复杂度：O(n^2)
// 空间复杂度：O(n)
func coinChange2(coins []int, amount int) int {
	sort.Ints(coins)

	// Empty 如果为 math.MaxInt,由于 dfs(n)+1, 会导致溢出为 math.MinInt
	// 因此设置个非常大的数，  0 <= amount <= 10^4
	const Empty = 1_0000_0000

	var dp = make([]int, amount+1)
	for i := 1; i <= amount; i++ {
		dp[i] = Empty
		idx := sort.Search(len(coins), func(j int) bool {
			return coins[j] > i
		})
		for _, coin := range coins[:idx] {
			// 逐个替换硬币
			dp[i] = _min(dp[i], dp[i-coin]+1)
		}
	}

	var ans = dp[amount]
	if ans == Empty {
		return -1
	}
	return ans
}

// dp 动态规划
// 时间复杂度：O(n^2)
// 空间复杂度：O(n)
func coinChange3(coins []int, amount int) int {
	// Empty 如果为 math.MaxInt,由于 dfs(n)+1, 会导致溢出为 math.MinInt
	// 因此设置个非常大的数，  0 <= amount <= 10^4
	const Empty = 1_0000_0000

	var dp = make([]int, amount+1)
	for i := 1; i <= amount; i++ {
		dp[i] = Empty
	}
	for _, coin := range coins {
		for j := coin; j <= amount; j++ {
			dp[j] = _min(dp[j], dp[j-coin]+1)
		}
	}

	if dp[amount] == Empty {
		return -1
	}
	return dp[amount]
}
