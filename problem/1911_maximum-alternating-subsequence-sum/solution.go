package main

import (
	"math"
)

// dp 动态规划, 状态机, 股票问题
// 时间复杂度：O(n)
// 空间复杂度：O(n)
func maxAlternatingSum(nums []int) int64 {
	type State struct {
		Add int
		Sub int
	}
	var n = len(nums)
	var dp = make([]State, n+1)
	dp[0].Add = math.MinInt32

	for i, num := range nums {
		// 持有
		dp[i+1].Add = _max(
			dp[i].Add,     // 保持不变
			dp[i].Sub+num, // 上一步 Sub -> 这一步 Add
		)

		// 未持有
		dp[i+1].Sub = _max(
			dp[i].Sub,     // 保持不变
			dp[i].Add-num, // 上一步 Add -> 这一步 Sub
		)
	}

	return int64(dp[n].Add)
}

func _max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

// 贪心
// 时间复杂度：O(n)
// 空间复杂度：O(1)
func maxAlternatingSum2(nums []int) int64 {
	profit, last := 0, 0 // 第一支股票购入价格是0
	for _, price := range nums {
		if last < price {
			profit += price - last // 卖出股票
		}

		last = price
	}
	return int64(profit)
}

// [122. 买卖股票的最佳时机 II](https://leetcode.cn/problems/best-time-to-buy-and-sell-stock-ii/)
// 时间复杂度：O(n)
// 空间复杂度：O(1)
func maxProfit(prices []int) int {
	profit, last := 0, prices[0]
	for _, price := range prices[1:] {
		if last < price {
			profit += price - last // 卖出股票
		}

		last = price
	}
	return profit
}
