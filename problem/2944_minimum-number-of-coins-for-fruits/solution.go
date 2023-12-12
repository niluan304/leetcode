package main

import "math"

// dfs + 记忆化搜索
// 时间复杂度: O(n^2)
// 空间复杂度: O(n)
func minimumCoins(prices []int) int {
	var n = len(prices)
	var cache = make([]*int, n+1)
	var dfs func(i int) int
	dfs = func(i int) int {
		if i > n {
			return 0
		}
		ptr := &cache[i]
		if *ptr != nil {
			return **ptr
		}

		var res = math.MaxInt32
		for j := i + 1; j <= i+1+i; j++ {
			res = min(res, dfs(j))
		}
		res += prices[i-1]
		*ptr = &res
		return res
	}
	return dfs(1)
}

// dp
// 时间复杂度: O(n^2)
// 空间复杂度: O(n)
func minimumCoins2(prices []int) int {
	var n = len(prices)
	var dp = make([]int, n+1)
	for i := n; i >= 1; i-- {
		if i*2 >= n {
			dp[i] = prices[i-1] // 可以写为 dp = append([]int{0}, prices...)
			continue
		}
		var res = math.MaxInt32
		for j := i + 1; j <= i+1+i; j++ {
			res = min(res, dp[j])
		}
		dp[i] = res + prices[i-1]
	}

	return dp[1]
}

// dp
// 时间复杂度: O(n^2)
// 空间复杂度: O(n)
func minimumCoins3(prices []int) int {
	var n = len(prices)
	var dp = append([]int{0}, prices...)
	for i := (n+1)/2 - 1; i >= 1; i-- {
		var res = math.MaxInt32
		for j := i + 1; j <= i+1+i; j++ {
			res = min(res, dp[j])
		}
		dp[i] += res
	}

	return dp[1]
}

// dp + 状态机
// 时间复杂度: O(n^2)
// 空间复杂度: O(n)
func minimumCoins4(prices []int) int {
	var n = len(prices)
	type Item struct{ Buy, NotBuy int }
	var dp = make([]Item, n+1) // dp[i]：买和不买 第 i 个水果的最低花费

	for i := 1; i <= n; i++ {
		notBuy := math.MaxInt32
		// 赠送的范围：比如 i = 10, j ∈ [5,9]; i = 9, j ∈ [5,8].
		for j := (i + 1) / 2; j <= i-1; j++ {
			// 不用买第 i 个的最低花费：在赠送范围内选择花费最少的
			notBuy = min(notBuy, dp[j].Buy)
		}
		dp[i] = Item{
			// 需要买第 i 个的最低花费: min(dp[i-1]) +  prices[i]
			Buy:    min(dp[i-1].NotBuy, dp[i-1].Buy) + prices[i-1],
			NotBuy: notBuy,
		}
	}
	return min(dp[n].Buy, dp[n].NotBuy)
}

// dfs + 记忆化搜索 + 状态机
// 时间复杂度: O(n^2)
// 空间复杂度: O(n)
func minimumCoins5(prices []int) int {
	var n = len(prices)
	type Item struct{ Buy, NotBuy int }
	var cache = make([]*Item, n+1)

	var dfs func(i int) Item
	dfs = func(i int) Item {
		pp := &cache[i]
		if *pp != nil {
			return **pp
		}
		if i == 0 {
			return Item{Buy: 0, NotBuy: 0}
		}

		var notBuy = math.MaxInt32
		for j := i - 1; j*2 >= i; j-- {
			notBuy = min(notBuy, dfs(j).Buy)
		}

		last := dfs(i - 1)
		item := Item{
			Buy:    min(last.Buy, last.NotBuy) + prices[i-1],
			NotBuy: notBuy,
		}
		cache[i] = &item
		return item
	}
	ans := dfs(n)
	return min(ans.Buy, ans.NotBuy)
}
