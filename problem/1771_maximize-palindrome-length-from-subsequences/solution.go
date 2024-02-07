package main

// dp 动态规划，区间dp
// 时间复杂度：O(n^2)
// 空间复杂度：O(n^2)
func longestPalindrome(word1 string, word2 string) int {
	var n, m = len(word1), len(word2)

	var cache = make([][]*int, n+1)
	for i, _ := range cache {
		cache[i] = make([]*int, m+1)
	}

	dp := longestPalindromeSubseq(word1 + word2)

	var dfs func(i, j int) int
	dfs = func(i, j int) int {
		if i >= n || j < 0 {
			return 0
		}

		var ptr = &cache[i][j]
		if *ptr != nil {
			return **ptr
		}

		*ptr = new(int)
		if word1[i] == word2[j] {
			**ptr = dp[i+1][n+j-1] + 2
		} else {
			**ptr = max(dfs(i+1, j), dfs(i, j-1))
		}
		return **ptr
	}
	return dfs(0, m-1)
}

// dfs + 记忆化搜索
// 时间复杂度：O(n^2)
// 空间复杂度：O(n^2)
func longestPalindrome2(word1 string, word2 string) int {
	var (
		s     = word1 + word2
		n     = len(s)
		bound = len(word1)
	)

	var cache = make([][]*int, n)
	for i, _ := range cache {
		cache[i] = make([]*int, n)
	}

	var ans = 0
	var dfs func(i, j int) int
	dfs = func(i, j int) int {
		if i > j {
			return 0
		}
		if i == j {
			return 1
		}

		var ptr = &cache[i][j]
		if *ptr != nil {
			return **ptr
		}

		*ptr = new(int)
		if s[i] == s[j] {
			**ptr = dfs(i+1, j-1) + 2
			if i < bound && bound <= j {
				ans = max(ans, **ptr)
			}
		} else {
			**ptr = max(dfs(i+1, j), dfs(i, j-1))
		}
		return **ptr
	}
	dfs(0, n-1)
	return ans
}

func longestPalindrome3(word1 string, word2 string) int {
	var (
		s     = word1 + word2
		n     = len(s)
		bound = len(word1)
	)

	const Use = 1

	var cache = make([][][2]*int, n)
	for i, _ := range cache {
		cache[i] = make([][2]*int, n)
	}

	var dfs func(i, j int, k int) int
	dfs = func(i, j int, k int) int {
		if (i > j) || (i >= bound && k != Use) || (j < bound && k != Use) {
			return 0
		}
		if i == j {
			return 1
		}

		var ptr = &cache[i][j][k]
		if *ptr != nil {
			return **ptr
		}

		*ptr = new(int)
		if s[i] == s[j] {
			**ptr = dfs(i+1, j-1, Use) + 2
		} else {
			**ptr = max(dfs(i+1, j, k), dfs(i, j-1, k))
		}
		return **ptr
	}

	return dfs(0, n-1, 0)
}

func longestPalindrome4(word1 string, word2 string) int {
	var (
		s     = word1 + word2
		n     = len(s)
		bound = len(word1)
	)

	var cache = make([][][2]*int, n)
	for i, _ := range cache {
		cache[i] = make([][2]*int, n)
	}

	var dfs func(i, j int, use bool) int
	dfs = func(i, j int, use bool) int {
		if (i > j) || (i >= bound && !use) || (j < bound && !use) {
			return 0
		}
		if i == j {
			return 1
		}

		k := 0
		if use {
			k = 1
		}

		var ptr = &cache[i][j][k]
		if *ptr != nil {
			return **ptr
		}

		*ptr = new(int)
		if s[i] == s[j] {
			**ptr = dfs(i+1, j-1, true) + 2
		} else {
			**ptr = max(dfs(i+1, j, use), dfs(i, j-1, use))
		}
		return **ptr
	}

	return dfs(0, n-1, false)
}

// 516. 最长回文子序列 https://leetcode.cn/problems/longest-palindromic-subsequence/
// dp 动态规划，区间dp
// 时间复杂度：O(n^2)
// 空间复杂度：O(n^2)
func longestPalindromeSubseq(s string) [][]int {
	var n = len(s)
	var dp = make([][]int, n)
	for i, _ := range dp {
		dp[i] = make([]int, n)
	}

	for i := n - 1; i >= 0; i-- {
		dp[i][i] = 1
		for j := i + 1; j < n; j++ {
			if s[i] == s[j] {
				dp[i][j] = dp[i+1][j-1] + 2
			} else {
				dp[i][j] = max(dp[i+1][j], dp[i][j-1])
			}
		}
	}
	return dp
}
