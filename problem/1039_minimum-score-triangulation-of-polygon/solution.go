package main

import "math"

// dfs + 记忆化搜索
// 时间复杂度：O(n^3)
// 空间复杂度：O(n^2)
func minScoreTriangulation(values []int) int {
	var n = len(values)
	var cache = make([][]*int, n)
	for i := 0; i < n; i++ {
		// nil 表示还没有计算过
		cache[i] = make([]*int, n)
	}
	var dfs func(i, j int) int
	dfs = func(i, j int) int {
		if i+1 == j { // 只有两个点，无法组成三角形
			return 0
		}

		if cache[i][j] != nil {
			return *cache[i][j]
		}
		var v = math.MaxInt32
		for k := i + 1; k < j; k++ { // 枚举顶点 k
			v = min(v, dfs(i, k)+values[i]*values[k]*values[j]+dfs(k, j))
		}
		cache[i][j] = &v
		return v
	}
	return dfs(0, n-1)
}

// dp 动态规划，区间dp
// 时间复杂度：O(n^3)
// 空间复杂度：O(n^2
func minScoreTriangulation2(values []int) int {
	var n = len(values)
	var dp = make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, n)
	}

	for i := n - 1; i >= 0; i-- {
		for j := i + 2; j < n; j++ {
			v := math.MaxInt32
			for k := i + 1; k < j; k++ {
				v = min(v, values[i]*values[k]*values[j]+dp[i][k]+dp[k][j])
			}
			dp[i][j] = v
		}
	}

	return dp[0][n-1]
}
