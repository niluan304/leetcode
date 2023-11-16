package main

// 代码示例
//
// 执行用时：4ms
// 消耗内存：1.895mb
func leetcodeMinTime(s string) int {
	dp := make([]int, len(s)+1)

	for i := range dp {
		dp[i] = len(s)
	}

	for i := 0; i <= len(s); i++ {
		expand(i, i, s, dp)
		expand(i, i+1, s, dp)
	}
	return dp[len(s)]
}

func expand(l, r int, s string, dp []int) {
	for l >= 0 && r < len(s) && s[l] == s[r] {
		if l == 0 {
			dp[r+1] = 0
		} else {
			dp[r+1] = min(dp[r+1], dp[l]+1)
		}
		l--
		r++
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
