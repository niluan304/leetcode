package main

import (
	"math"
	_ "net/http/pprof" // Blank import to pprof
)

// 贪心
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

// dfs + 记忆化搜索
// 时间复杂度：O(n^3)
// 空间复杂度：O(n^2)
func maxProfit2(prices []int) int {
	type Border struct{ Left, Right int }
	var cache = make(map[Border]int)

	var dfs func(i, j int) int
	dfs = func(i, j int) int {
		key := Border{Left: i, Right: j}
		v, ok := cache[key]
		if ok {
			return v
		}

		if i+1 == j {
			return max(0, prices[j]-prices[i])
		}
		for k := i + 1; k < j; k++ {
			v = max(v, dfs(i, k)+dfs(k, j))
		}
		cache[key] = v
		return v
	}

	//defer func() {
	//	// debug: 打印 cache
	//	//fmt.Println(cache)
	//	fmt.Println(len(prices), len(cache))
	//}()

	return dfs(0, len(prices)-1)
}

// dfs + 记忆化搜索
// 时间复杂度：O(n)
// 空间复杂度：O(n)
func maxProfit3(prices []int) int {
	type State struct {
		N    int
		Hold bool
	}

	var n = len(prices)
	var cache = make(map[State]int, n*2)

	var dfs func(i int, hold bool) int
	dfs = func(i int, hold bool) int {
		if i < 0 {
			if hold {
				return math.MinInt32 // 使用 math.MinInt - x 会导致溢出
			} else {
				return 0
			}
		}

		key := State{N: i, Hold: hold}
		v, ok := cache[key]
		if ok {
			return v
		}

		// 状态机
		if hold {
			v = max(
				dfs(i-1, true),            // 保持不变：昨天也持有股票
				dfs(i-1, false)-prices[i], // 没有股票 -> 买入股票
			)
		} else {
			v = max(
				dfs(i-1, false),          // 保持不变：昨天也未持有股票
				dfs(i-1, true)+prices[i], // 持有股票 -> 卖出股票
			)
		}

		cache[key] = v
		return v
	}

	return dfs(n-1, false)
}

// dp 动态规划，状态机
// 时间复杂度：O(n)
// 空间复杂度：O(n)
func maxProfit4(prices []int) int {
	var n = len(prices)
	var dp = make([]struct{ Hold, NotHold int }, n+1)
	dp[0].Hold = math.MinInt32

	for i, price := range prices {
		// 持有股票的状态
		dp[i+1].Hold = max(
			dp[i].Hold,          // 保持不变：昨天也持有股票
			dp[i].NotHold-price, // 没有股票 -> 买入股票
		)

		// 未持有股票的状态
		dp[i+1].NotHold = max(
			dp[i].NotHold,    // 保持不变：昨天也未持有股票
			dp[i].Hold+price, // 持有股票 -> 卖出股票
		)
	}

	return dp[n].NotHold
}

// dp 动态规划，状态机 + 压缩空间复杂度
// 时间复杂度：O(n)
// 空间复杂度：O(1)
func maxProfit5(prices []int) int {
	var dp struct{ Hold, NotHold int }
	dp.Hold = math.MinInt32

	for _, price := range prices {
		last := dp

		// 持有股票的状态
		dp.Hold = max(
			last.Hold,          // 保持不变：昨天也持有股票
			last.NotHold-price, // 没有股票 -> 买入股票
		)

		// 未持有股票的状态
		dp.NotHold = max(
			last.NotHold,    // 保持不变：昨天也未持有股票
			last.Hold+price, // 持有股票 -> 卖出股票
		)
	}
	return dp.NotHold
}
