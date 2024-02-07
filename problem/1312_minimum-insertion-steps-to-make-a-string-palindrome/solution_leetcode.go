package main

// 题解
// 本题属于 516. 最长回文子序列 的子题，求的是将 s 变成回文串需要添加的最少字符数，所以我们只用求最长回文子序列长度即可，
// 然后字符串 s 中除去最长回文子序列，剩下的字符就是不构成回文子序列的字符，添加与其同等数量的字符，将 s 构成回文串。
//
// 作者：return up;
// 链接：https://leetcode.cn/problems/minimum-insertion-steps-to-make-a-string-palindrome/description/
// 来源：力扣（LeetCode）
// 著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
func leetcodeReturnUp(s string) int {
	return len(s) - longestPalindromeSubseq(s)
}

// 516. 最长回文子序列 https://leetcode.cn/problems/longest-palindromic-subsequence/
// dp 动态规划，区间dp
// 时间复杂度：O(n^2)
// 空间复杂度：O(n^2)
func longestPalindromeSubseq(s string) int {
	var n = len(s)
	var dp = make([][]int, n)
	for i := 0; i < n; i++ {
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

	return dp[0][n-1]
}
