package main

import (
	"math"
)

// dp 动态规划，状态机
// 时间复杂度：O(n)
// 空间复杂度：O(n)
func maxProfit(prices []int) int {
	var n = len(prices)
	var dp = make([]struct{ Hold, NotHold int }, n+1)
	dp[0].Hold = math.MinInt32
	dp[1].Hold = -prices[0]

	// 避免 dp[1] 时数组越界, 因此直接初始化 dp[1]，且循环从 dp[2] 开始
	for i := 1; i < n; i++ {
		price := prices[i]

		// 持有股票的状态
		dp[i+1].Hold = _max(
			dp[i].Hold,            // 保持不变：昨天也持有股票
			dp[i-1].NotHold-price, // 没有股票 -> 买入股票
		)

		// 未持有股票的状态
		dp[i+1].NotHold = _max(
			dp[i].NotHold,    // 保持不变：昨天也未持有股票
			dp[i].Hold+price, // 持有股票 -> 卖出股票
		)
	}

	//// debug: 打印 dp 数组
	//fmt.Println(dp)

	return dp[n].NotHold
}

func _max(x, y int) int {
	if x < y {
		x = y
	}
	return x
}
