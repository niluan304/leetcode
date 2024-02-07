package main

// dfs + 记忆化搜索
// 时间复杂度：O(mn)
// 空间复杂度：O(mn)
func longestCommonSubsequence(text1 string, text2 string) int {
	var n, m = len(text1), len(text2)
	var cache = make([][]int, n)
	const Empty = -1
	for i := 0; i < n; i++ {
		cache[i] = make([]int, m)
		for j := 0; j < m; j++ {
			cache[i][j] = Empty
		}
	}
	var dfs func(i, j int) int
	dfs = func(i, j int) int {
		if i < 0 || j < 0 {
			return 0
		}
		v := cache[i][j]
		if v > Empty {
			return v
		}
		if text1[i] == text2[j] {
			v = dfs(i-1, j-1) + 1
		} else {
			v = max(dfs(i, j-1), dfs(i-1, j))
		}
		cache[i][j] = v
		return v
	}
	return dfs(n-1, m-1)
}
