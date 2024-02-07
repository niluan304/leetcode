package main

// dfs + 记忆化搜索
// 时间复杂度：O(mn)
// 空间复杂度：O(mn)
func minimumDeleteSum(s1 string, s2 string) int {
	var n, m = len(s1), len(s2)
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
		// s1 为空时, 所需要的操作数
		if i < 0 {
			return CountString(s2[:j+1])
		}
		// s2 为空时, 所需要的操作数
		if j < 0 {
			return CountString(s1[:i+1])
		}

		v := cache[i][j]
		if v > 0 {
			return v
		}

		if s1[i] == s2[j] {
			v = dfs(i-1, j-1)
		} else {
			v = min(
				dfs(i-1, j)+int(s1[i]), // 删除 s1 的字符
				dfs(i, j-1)+int(s2[j]), // 删除 s2 的字符
			)
		}

		cache[i][j] = v
		return v
	}
	return dfs(n-1, m-1)
}

func CountString(s string) int {
	var count = 0
	for i := 0; i < len(s); i++ {
		count += int(s[i])
	}
	return count
}
