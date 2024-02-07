package main

import (
	"math"
)

// dfs + 记忆化搜索
// 时间复杂度：O(n)
// 空间复杂度：O(n)
func maxProfit(k int, prices []int) int {
	type Profit struct {
		Hold    *int
		NotHold *int
	}
	var n = len(prices)
	var cache = make([][]Profit, n+1)
	for i := range cache {
		cache[i] = make([]Profit, k+1)
	}

	var dfs func(i int, hold bool, k int) int
	dfs = func(i int, hold bool, k int) int {
		if i < 0 {
			if hold {
				return math.MinInt
			} else {
				return 0
			}
		}
		if k < 0 {
			return math.MinInt
		}

		v := cache[i][k].NotHold
		if hold {
			v = cache[i][k].Hold
		}
		if v != nil {
			return *v
		}

		// 状态机
		if hold {
			v := max(
				dfs(i-1, true, k),            // 保持不变：昨天也持有股票
				dfs(i-1, false, k)-prices[i], // 买入股票
			)
			cache[i][k].Hold = &v
			return v
		} else {
			v := max(
				dfs(i-1, false, k),            // 保持不变：昨天也未持有股票
				dfs(i-1, true, k-1)+prices[i], // 卖出股票, 交易次数上限 -1
			)
			cache[i][k].NotHold = &v
			return v
		}
	}

	// 最后一天未持有股票，且交易次数达到上限：k次。
	return dfs(n-1, false, k)
}

// dp 动态规划，状态机
// 时间复杂度：O(k*n)
// 空间复杂度：O(k*n)
func maxProfit2(k int, prices []int) int {
	type Profit struct {
		Hold    int
		NotHold int
	}

	var n = len(prices)
	var dp = make([][]Profit, n+1)
	for i := range dp {
		dp[i] = make([]Profit, k+1)
		dp[i][0].Hold = math.MinInt32 // 交易次数为 0 时，无法持有股票
	}
	for i := 1; i <= k; i++ {
		dp[0][i] = Profit{
			Hold:    math.MinInt32,
			NotHold: math.MinInt32,
		}
	}

	for i, price := range prices {
		for j := 1; j <= k; j++ {
			// 持有股票的状态
			dp[i+1][j].Hold = max(
				dp[i][j].Hold,            // 保持不变：昨天也持有股票
				dp[i][j-1].NotHold-price, // 没有股票 -> 买入股票
			)

			// 未持有股票的状态
			dp[i+1][j].NotHold = max(
				dp[i][j].NotHold,    // 保持不变：昨天也未持有股票
				dp[i][j].Hold+price, // 持有股票 -> 卖出股票
			)
		}
	}

	//fmt.Println("k, prices", k, prices)
	//for i := 0; i <= n; i++ {
	//	fmt.Println(dp[i])
	//}

	var ans = 0
	for _, profit := range dp[n] {
		ans = max(ans, profit.NotHold)
	}
	return ans
}

// dp 动态规划，状态机 + 压缩空间复杂度
// 时间复杂度：O(k*n)
// 空间复杂度：O(k)
func maxProfit3(k int, prices []int) int {
	type Profit struct {
		Hold    int
		NotHold int
	}

	var dp = make([]Profit, k+1)
	for i := 0; i <= k; i++ {
		dp[i].Hold = math.MinInt32 // 交易次数为 0 时，无法持有股票
	}

	for _, price := range prices {
		for j := k; j >= 1; j-- {
			last := dp[j]
			// 持有股票的状态
			dp[j].Hold = max(
				last.Hold,             // 保持不变：昨天也持有股票
				dp[j-1].NotHold-price, // 没有股票 -> 买入股票
			)

			// 未持有股票的状态
			dp[j].NotHold = max(
				last.NotHold,    // 保持不变：昨天也未持有股票
				last.Hold+price, // 持有股票 -> 卖出股票
			)
		}
	}

	return dp[k].NotHold
}
