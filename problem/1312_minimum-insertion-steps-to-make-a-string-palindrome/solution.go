package main

import "math"

// dfs + 记忆化搜搜
// 时间复杂度：O(n^2)
// 空间复杂度：O(n^2)
func minInsertions(s string) int {
	var n = len(s)
	var cache = make([][]*int, n)
	for i, _ := range cache {
		cache[i] = make([]*int, n)
	}

	var dfs func(i, j int) int
	dfs = func(i, j int) int {
		if i >= j {
			return 0 // i == j 表示单个字符，i > j 无意义
		}

		if cache[i][j] != nil {
			return *cache[i][j]
		}

		v := math.MaxInt32
		if s[i] == s[j] {
			v = dfs(i+1, j-1) // 首尾字母相同，无需插入字符
		} else {
			v = min(dfs(i+1, j), dfs(i, j-1)) + 1 // 首尾字母不相同，需要在首或尾插入一个字符
		}

		cache[i][j] = &v
		return v
	}
	return dfs(0, n-1)
}

// dp 动态规划，区间dp
// 时间复杂度：O(n^2)
// 空间复杂度：O(n^2)
func minInsertions2(s string) int {
	var n = len(s)
	var dp = make([][]int, n)
	for i, _ := range dp {
		dp[i] = make([]int, n)
	}

	for i := n - 1; i >= 0; i-- {
		for j := i + 1; j < n; j++ {
			if s[i] == s[j] {
				dp[i][j] = dp[i+1][j-1] // 首尾字母相同，无需插入字符
			} else {
				dp[i][j] = min(dp[i+1][j], dp[i][j-1]) + 1 // 首尾字母不相同，需要在首或尾插入一个字符
			}
		}
	}

	return dp[0][n-1]
}
