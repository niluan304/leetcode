package main

import (
	"math"
)

// dp 动态规划，状态机
// 时间复杂度：O(n)
// 空间复杂度：O(n)
func maxProfit(prices []int) int {
	type Key struct {
		N    int
		Hold bool
		K    int
	}

	var k = 2

	var n = len(prices)
	var cache = make(map[Key]int, n*(2*k+1))

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

		key := Key{N: i, Hold: hold, K: k}
		v, ok := cache[key]
		if ok {
			return v
		}

		// 状态机
		if hold {
			v = max(
				dfs(i-1, true, k),            // 保持不变
				dfs(i-1, false, k)-prices[i], //买入股票
			)
		} else {
			v = max(
				dfs(i-1, false, k),            // 保持不变
				dfs(i-1, true, k-1)+prices[i], // 卖出股票
			)
		}

		cache[key] = v
		return v
	}

	// 最后一天未持有股票，且交易次数达到上限：k次。
	return dfs(n-1, false, k)
}

// 贪心?
// 时间复杂度：O(n^2)
// 空间复杂度：O(1)
func maxProfit2(prices []int) int {
	profit := 0
	for i := 0; i < len(prices)-2; i++ {
		profit = max(
			profit,
			maxProfit1(prices[:i])+maxProfit1(prices[i:]), // 两次交易
		)
	}
	return profit
}

func maxProfit1(prices []int) int {
	low := math.MaxInt
	profit := 0
	for _, price := range prices {
		low = min(low, price)
		profit = max(profit, price-low)
	}

	return profit
}

// dfs + 记忆化搜索
// 时间复杂度：O(n)
// 空间复杂度：O(n)
func maxProfit3(prices []int) int {
	type Profit struct {
		Hold    *int
		NotHold *int
	}
	var n = len(prices)
	var cache = make([][3]Profit, n)

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

		v = new(int)
		// 状态机
		if hold {
			*v = max(
				dfs(i-1, true, k),            // 保持不变：昨天也持有股票
				dfs(i-1, false, k)-prices[i], // 买入股票
			)
			cache[i][k].Hold = v
		} else {
			*v = max(
				dfs(i-1, false, k),            // 保持不变：昨天也未持有股票
				dfs(i-1, true, k-1)+prices[i], // 卖出股票, 交易次数上限 -1
			)
			cache[i][k].NotHold = v
		}
		return *v
	}

	// 最后一天未持有股票，且交易次数达到上限：k次。
	return dfs(n-1, false, 2)
}
