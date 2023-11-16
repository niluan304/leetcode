package main

import (
	"math"
)

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

// Deprecated: O(n^3) timeout when n <= 2000.
// dfs + 记忆化搜索
// 时间复杂度：O(n^3)
// 空间复杂度：O(n^2)
func minCut(s string) int {

	var n = len(s)
	var cache = make([][]*int, n+1)
	for i := 0; i <= n; i++ {
		cache[i] = make([]*int, n+1)
	}

	var dfs func(i, j int) int
	dfs = func(i, j int) int {
		if cache[i][j] != nil {
			return *cache[i][j]
		}

		if isPalindrome(s[i:j]) {
			cache[i][j] = new(int)
			return 0
		}

		var v = math.MaxInt32
		for k := i + 1; k <= j-1; k++ {
			v = _min(
				v,
				dfs(i, k)+dfs(k, j)+1,
			)
		}
		cache[i][j] = &v
		return v
	}
	return dfs(0, n)
}

func isPalindrome(s string) bool {
	var n = len(s)
	for i := 0; i < n/2; i++ {
		if s[i] != s[n-1-i] {
			return false
		}
	}
	return true
}

// Deprecated: O(n^3) timeout when n <= 2000.
// dp 动态规划，区间dp
// 时间复杂度：O(n^3)，timeout，n <= 2000
// 空间复杂度：O(n^2)
func minCut2(s string) int {
	var n = len(s)
	var dp = make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}

	for i := n; i >= 0; i-- {
		for j := i + 1; j <= n; j++ {
			if isPalindrome(s[i:j]) {
				continue
			}

			v := math.MaxInt32
			for k := i + 1; k <= j-1; k++ {
				v = _min(v, dp[i][k]+dp[k][j]+1)
			}
			dp[i][j] = v
		}
	}
	return dp[0][n]
}

// dp 动态规划
// 时间复杂度：O(n^2)
// 空间复杂度：O(n^2)
func minCut3(s string) int {
	var n = len(s)
	var notPalindrome = make([][]bool, n)
	for i := range notPalindrome {
		notPalindrome[i] = make([]bool, n)
	}

	for i := n - 1; i >= 0; i-- {
		for j := i + 1; j < n; j++ {
			notPalindrome[i][j] = notPalindrome[i+1][j-1] || s[i] != s[j]
		}
	}

	var dp = make([]int, n)
	for i := 0; i < n; i++ {
		// s[0:i+1] 本身就是一个回文串，不需要分割
		if !notPalindrome[0][i] {
			continue
		}

		dp[i] = math.MaxInt32
		for j := 0; j < i; j++ {
			// 如果 s[j+1:i+1] 是 回文字符串，那么就可以切割为 dp[i-1] + s[j+1:i+1]，因此 dp[i] = dp[i-1] + 1
			if !notPalindrome[j+1][i] && dp[j]+1 < dp[i] {
				dp[i] = dp[j] + 1
			}
		}
	}

	return dp[n-1]
}
