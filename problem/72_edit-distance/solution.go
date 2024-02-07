package main

// dfs + 记忆化搜索
// 时间复杂度：O(mn)
// 空间复杂度：O(mn)
func minDistance(word1 string, word2 string) int {
	var n, m = len(word1), len(word2)
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
		if i < 0 {
			return j + 1
		}
		if j < 0 {
			return i + 1
		}

		v := cache[i][j]
		if v > Empty {
			return v
		}

		if word1[i] == word2[j] {
			v = dfs(i-1, j-1)
		} else {
			v = min(
				dfs(i, j-1),   // 插入
				dfs(i-1, j),   // 删除
				dfs(i-1, j-1), // 替换
			) + 1
		}
		cache[i][j] = v
		return v
	}
	return dfs(n-1, m-1)
}
