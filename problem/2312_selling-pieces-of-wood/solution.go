package main

// dfs + 记忆化搜索
//
// 根据「横切」「竖切」「不切」的操作，将原问题分解为子问题
//
// - 时间复杂度：O(nm(m+n))。
// - 空间复杂度：O(nm)。
func sellingWood(m int, n int, prices [][]int) int64 {
	type Pair struct{ h, w int }

	costs := make(map[Pair]int, len(prices))
	for _, price := range prices {
		h, w, p := price[0], price[1], price[2]
		costs[Pair{h: h, w: w}] = p // 题目保证 key 唯一，那么就不用取 max
	}

	// 使用 map 可能会超时
	cache := make([][]*int, m+1)
	for i := range cache {
		cache[i] = make([]*int, n+1)
	}

	var dfs func(i, j int) int
	dfs = func(i, j int) (res int) {
		if i == 0 || j == 0 {
			return 0
		}

		ptr := &cache[i][j]
		if *ptr != nil {
			return **ptr
		}
		defer func() { *ptr = &res }()

		// 不切
		res = costs[Pair{h: i, w: j}]

		// 横切，枚举高度 k
		for k := 1; k <= i/2; k++ {
			res = max(res, dfs(k, j)+dfs(i-k, j))
		}

		// 竖切，枚举宽度 k
		for k := 1; k <= j/2; k++ {
			res = max(res, dfs(i, k)+dfs(i, j-k))
		}
		return res
	}
	ans := dfs(m, n)
	return int64(ans)
}

// 将 sellingWood 1:1 翻译为递推
//
// - 时间复杂度：O(nm(m+n))。
// - 空间复杂度：O(nm)。
func sellingWood2(m int, n int, prices [][]int) int64 {
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}

	// 在计算递推之前，直接将 prices 记录到 dp 数组中，也就是不切时的售卖价格
	for _, price := range prices {
		h, w, p := price[0], price[1], price[2]
		dp[h][w] = p // 题目保证 key 唯一，那么就不用取 max
	}

	for i := 0; i <= m; i++ {
		for j := 0; j <= n; j++ {
			// 横切，枚举高度 k
			for k := 1; k <= i/2; k++ {
				dp[i][j] = max(dp[i][j], dp[k][j]+dp[i-k][j])
			}

			// 竖切，枚举宽度 k
			for k := 1; k <= j/2; k++ {
				dp[i][j] = max(dp[i][j], dp[i][k]+dp[i][j-k])
			}
		}
	}

	ans := dp[m][n]
	return int64(ans)
}
