package main

import (
	"math"
)

// dp 动态规划，状态机
// 时间复杂度：O(n)
// 空间复杂度：O(n)
func maxProfit(prices []int, fee int) int {
	var n = len(prices)
	var dp = make([]struct{ Hold, NotHold int }, n+1)
	dp[0].Hold = math.MinInt32

	for i := 0; i < n; i++ {
		price := prices[i]

		// 持有股票的状态
		dp[i+1].Hold = _max(
			dp[i].Hold,          // 保持不变：昨天也持有股票
			dp[i].NotHold-price, // 没有股票 -> 买入股票
		)

		// 未持有股票的状态
		dp[i+1].NotHold = _max(
			dp[i].NotHold,        // 保持不变：昨天也未持有股票
			dp[i].Hold+price-fee, // 持有股票 -> 卖出股票
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

// dp 动态规划，状态机 + 压缩空间复杂度
// 时间复杂度：O(n)
// 空间复杂度：O(1)
func maxProfit2(prices []int, fee int) int {
	var dp struct{ Hold, NotHold int }
	dp.Hold = math.MinInt32

	for _, price := range prices {
		last := dp

		// 持有股票的状态
		dp.Hold = _max(
			last.Hold,          // 保持不变：昨天也持有股票
			last.NotHold-price, // 没有股票 -> 买入股票
		)

		// 未持有股票的状态
		dp.NotHold = _max(
			last.NotHold,        // 保持不变：昨天也未持有股票
			last.Hold+price-fee, // 持有股票 -> 卖出股票
		)
	}
	return dp.NotHold
}
