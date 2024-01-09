package main

import (
	"strings"
)

// ### 记忆化搜索
//
// #### 复杂度分析
//
// - 时间复杂度：$\mathcal{O}(n^2\cdotL)$，其中 $n$ 为 $s$ 的长度，$L$ 为 $\textit{dictionary}$ 所有字符串的长度之和。
// 记忆化搜索的时间复杂度 $=$ 状态个数 $\times$ 单个状态的计算时间。本题中状态个数等于 $\mathcal{O}(n)$，单个状态的计算时间为 $\mathcal{O}(n\cdotL)$，因此时间
// 复杂度为 $\mathcal{O}(n^2\cdotL)$。
// - 空间复杂度：$\mathcal{O}(n)$。
func minExtraChar(s string, dictionary []string) int {
	var n = len(s)
	var cache = make([]*int, n)
	var dfs func(i int) int
	dfs = func(i int) int {
		if i == n {
			return 0
		}

		ptr := &cache[i]
		if *ptr != nil {
			return **ptr
		}

		v := dfs(i+1) + 1 // 不选：最差的情况就是这个字符无法被使用
		for _, dict := range dictionary {
			if strings.HasPrefix(s[i:], dict) { // 字符串比较的时间复杂度：$\mathcal{O}(n)$
				v = min(v, dfs(i+len(dict))) // 枚举选哪个：如果字典是字符串s[i:]的前缀，那就可以转换为 s[i+len(dict):] 下的最优选
			}
		}

		*ptr = &v
		return v
	}

	ans := dfs(0)
	return ans
}

// ### 将记忆化搜索翻译成动态规划
//
// #### 复杂度分析
//
// - 时间复杂度：$\mathcal{O}(n^2\cdotL)$，其中 $n$ 为 $s$ 的长度，$L$ 为 $\textit{dictionary}$ 所有字符串的长度之和。
// 动态规划的时间复杂度 $=$ 状态个数 $\times$ 单个状态的计算时间。本题中状态个数等于 $\mathcal{O}(n)$，单个状态的计算时间为 $\mathcal{O}(n\cdotL)$，因此时间
// 复杂度为 $\mathcal{O}(n^2\cdotL)$。
// - 空间复杂度：$\mathcal{O}(n)$。
func minExtraChar2(s string, dictionary []string) int {
	var n = len(s)
	var dp = make([]int, n+1)
	for i := n - 1; i >= 0; i-- {
		dp[i] = dp[i+1] + 1 // 不选：最差的情况就是这个字符无法被使用

		for _, dict := range dictionary {
			if strings.HasPrefix(s[i:], dict) { // 字符串比较的时间复杂度：$\mathcal{O}(n)$
				dp[i] = min(dp[i], dp[i+len(dict)]) // 枚举选哪个：如果字典是字符串s[i:]的前缀，那就可以转换为 s[i+len(dict):] 下的最优选
			}
		}
	}
	return dp[0]
}
