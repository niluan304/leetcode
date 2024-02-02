package main

import . "github.com/niluan304/leetcode/copypasta/dp"

// dfs + 记忆化搜索
//
// 原题：
// [LC 2008. 出租车的最大盈利](https://leetcode.cn/problems/maximum-earnings-from-taxi/)
//
// - 时间复杂度：$\mathcal{O}(n + m)$。
// - 空间复杂度：$\mathcal{O}(n + m)$。
func maximizeTheProfit(n int, offers [][]int) int {
	type Pair struct{ start, gold int }
	var pairs = make([][]Pair, n)
	for _, offer := range offers {
		start, end, gold := offer[0], offer[1], offer[2]
		pairs[end] = append(pairs[end], Pair{start: start, gold: gold})
	}

	var cache = make([]*int, n)
	var dfs func(i int) int
	dfs = func(i int) (res int) {
		if i < 0 {
			return 0
		}

		ptr := &cache[i]
		if *ptr != nil {
			return **ptr
		}
		defer func() { *ptr = &res }()

		res = dfs(i - 1)

		for _, pair := range pairs[i] {
			res = max(res, dfs(pair.start-1)+pair.gold)
		}
		return res
	}

	ans := dfs(n - 1)
	return ans
}

// dfs + 记忆化搜索
//
// - 时间复杂度：$\mathcal{O}(n + m)$。
// - 空间复杂度：$\mathcal{O}(n + m)$。
func maximizeTheProfit2(n int, offers [][]int) int {
	type Pair struct{ start, gold int }
	var pairs = make([][]Pair, n)
	for _, offer := range offers {
		start, end, gold := offer[0], offer[1], offer[2]
		pairs[end] = append(pairs[end], Pair{start: start, gold: gold})
	}

	var dfs func(i int) int
	dfs = func(i int) (res int) {
		if i < 0 {
			return 0
		}

		res = dfs(i - 1)
		for _, pair := range pairs[i] {
			res = max(res, dfs(pair.start-1)+pair.gold)
		}
		return res
	}

	MemorySearch(&dfs)
	ans := dfs(n - 1)
	return ans
}

// 将 maximizeTheProfit 1:1 翻译为递推
//
// - 时间复杂度：$\mathcal{O}(n + m)$。
// - 空间复杂度：$\mathcal{O}(n + m)$。
func maximizeTheProfit3(n int, offers [][]int) int {
	type Pair struct{ start, gold int }
	var pairs = make([][]Pair, n)
	for _, offer := range offers {
		start, end, gold := offer[0], offer[1], offer[2]
		pairs[end] = append(pairs[end], Pair{start: start, gold: gold})
	}

	var dp = make([]int, n+1)
	for i := 0; i < n; i++ {
		dp[i+1] = dp[i]
		for _, pair := range pairs[i] {
			dp[i+1] = max(dp[i+1], dp[pair.start]+pair.gold)
		}
	}

	ans := dp[n]
	return ans
}
