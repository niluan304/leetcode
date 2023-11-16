package main

// dfs + 记忆化搜索
// 时间复杂度：O(n^2)
// 空间复杂度：O(n^2)
func longestPalindromeSubseq(s string) int {
	var n = len(s)
	var cache = make([][]*int, n)
	for i := 0; i < n; i++ {
		cache[i] = make([]*int, n)
	}

	var dfs func(i, j int) int
	dfs = func(i, j int) int {
		if i > j {
			return 0
		}
		if i == j {
			return 1
		}

		if cache[i][j] != nil {
			return *cache[i][j]
		}

		var v = 0
		if s[i] == s[j] {
			v = dfs(i+1, j-1) + 2
		} else {
			v = _max(dfs(i+1, j), dfs(i, j-1))
		}
		cache[i][j] = &v
		return v
	}
	return dfs(0, n-1)
}

func _max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func _min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func _abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// dp 动态规划，区间dp
// 时间复杂度：O(n^2)
// 空间复杂度：O(n^2)
func longestPalindromeSubseq2(s string) int {
	var n = len(s)
	var dp = make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, n)
	}

	// i倒序遍历，因为需要先得到i+1
	for i := n - 1; i >= 0; i-- {
		dp[i][i] = 1
		// j从i+1正向遍历，使得j-1规避了负数，同时还使得i+1规避了上限
		for j := i + 1; j < n; j++ {
			if s[i] == s[j] {
				dp[i][j] = dp[i+1][j-1] + 2
			} else {
				dp[i][j] = _max(dp[i+1][j], dp[i][j-1])
			}
		}
	}

	return dp[0][n-1]
}
