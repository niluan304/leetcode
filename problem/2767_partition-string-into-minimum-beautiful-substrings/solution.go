package main

import (
	"math"
	"strconv"
	"strings"

	. "github.com/niluan304/leetcode/copypasta/dp"
)

// dfs + 记忆化搜索
//
// #### 复杂度分析
//
// - 时间复杂度：$\mathcal{O}(n^2)$，其中 $n$ 为 $s$ 的长度。动态规划的时间复杂度 $=$ 状态个数 $\times$ 单个状态的计算时间。
// 本题中状态个数等于 $\mathcal{O}(n)$，单个状态的计算时间为 $\mathcal{O}(n)$，所以动态规划的时间复杂度为 $\mathcal{O}(n^2)$。
// 注意在比较字符串时，由于 $\textit{pow}_5$ 中字符串的公共前缀很短，很容易就失配了，所以除了匹配上的那次是 $\mathcal{O}(n)$ 的，
// 其余匹配都可以视作是 $\mathcal{O}(1)$ 的。
// - 空间复杂度：$\mathcal{O}(n)$。
func minimumBeautifulSubstrings(s string) int {
	n := len(s)
	var pow5 []string
	for x := 1; ; x *= 5 {
		p5 := strconv.FormatInt(int64(x), 2)
		if len(p5) > n {
			break
		}
		pow5 = append(pow5, p5)
	}

	var dfs func(i int) int
	dfs = func(i int) int {
		if i == 0 {
			return 0
		}

		res := math.MaxInt32
		for _, num := range pow5 {
			if strings.HasSuffix(s[:i], num) {
				res = min(res, dfs(i-len(num))+1)
			}
		}
		return res
	}

	MemorySearch(&dfs)
	ans := dfs(n)
	if ans >= math.MaxInt32 {
		return -1
	}
	return ans
}

// 将 minimumBeautifulSubstrings 1:1 翻译为递推
//
// #### 复杂度分析
//
// - 时间复杂度：$\mathcal{O}(n^2)$，其中 $n$ 为 $s$ 的长度。
// - 空间复杂度：$\mathcal{O}(n)$。
func minimumBeautifulSubstrings2(s string) int {
	n := len(s)
	var pow5 []string
	for x := 1; ; x *= 5 {
		p5 := strconv.FormatInt(int64(x), 2)
		if len(p5) > n {
			break
		}
		pow5 = append(pow5, p5)
	}

	var dp = make([]int, n+1)
	for i := 1; i <= n; i++ {
		dp[i] = math.MaxInt32

		for _, p5 := range pow5 {
			if strings.HasSuffix(s[:i], p5) {
				dp[i] = min(dp[i], dp[i-len(p5)]+1)
			}
		}
	}

	ans := dp[n]
	if ans >= math.MaxInt32 {
		return -1
	}
	return ans
}
