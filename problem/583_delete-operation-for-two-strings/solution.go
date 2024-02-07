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
		// word1 为空时, 所需要的操作数
		if i < 0 {
			return j + 1
		}
		// word2 为空时, 所需要的操作数
		if j < 0 {
			return i + 1
		}

		v := cache[i][j]
		if v > 0 {
			return v
		}

		if word1[i] == word2[j] {
			v = dfs(i-1, j-1)
		} else {
			v = min(
				dfs(i-1, j), // 删除 word1 的字符
				dfs(i, j-1), // 删除 word2 的字符
			) + 1
		}

		cache[i][j] = v
		return v
	}
	return dfs(n-1, m-1)
}

// dp 动态规划
// 时间复杂度：O(mn)
// 空间复杂度：O(mn)
func minDistance2(word1 string, word2 string) int {
	var n, m = len(word1), len(word2)
	var dp = make([][]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]int, m+1)
	}

	// 初始化
	// word1 为空时, 所需要的操作数
	for i := 1; i <= n; i++ {
		dp[i][0] = i
	}
	// word2 为空时, 所需要的操作数
	for j := 1; j <= m; j++ {
		dp[0][j] = j
	}

	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			// 递推公式
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = min(dp[i][j-1], dp[i-1][j]) + 1
			}
		}
	}

	//// debug: 打印dp数组
	//for _, ints := range dp {
	//	fmt.Println(ints)
	//}

	return dp[n][m]
}

// dp 动态规划 + 压缩空间
// 时间复杂度：O(mn)
// 空间复杂度：O(n)
func minDistance3(word1 string, word2 string) int {
	var n, m = len(word1), len(word2)
	if n > m {
		m, n = n, m
		word1, word2 = word2, word1
		// or
		//minDistance3(word2, word1)
	}
	var dp = make([]int, n+1)

	// 初始化
	// word1 为空时, 所需要的操作数
	for i := 1; i <= n; i++ {
		dp[i] = i
	}

	for i := 1; i <= m; i++ {
		last := dp[0]
		// word2 为空时, 所需要的操作数
		dp[0]++
		for j := 1; j <= n; j++ {
			// 递推公式
			tmp := dp[j] // 缓存左上角的值
			if word1[j-1] == word2[i-1] {
				dp[j] = last
			} else {
				dp[j] = min(dp[j], dp[j-1]) + 1
			}
			last = tmp
		}

		//// debug: 打印dp数组
		//fmt.Println(dp)
	}

	return dp[n]
}
