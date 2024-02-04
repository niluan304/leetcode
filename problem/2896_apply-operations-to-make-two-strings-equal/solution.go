package main

import (
	"math"
	"strings"

	. "github.com/niluan304/leetcode/copypasta/dp"
)

// dfs + 记忆化搜索
//
// - 时间复杂度：$\mathcal{O}(n^2 \cdot ∑)$。 $∑ = 2$。
// - 空间复杂度：$\mathcal{O}(n^2 \cdot ∑)$。 $∑ = 2$。
func minOperations(s1 string, s2 string, x int) int {
	// 反转操作不回改变 "0"，"1" 个数的奇偶性，
	// 如果 s1,s2 的不同，那么返回 -1，
	// 否则，哪怕只用第二种操作，都可以实现 s1 = s2。
	if strings.Count(s1, "0")%2 != strings.Count(s2, "0")%2 {
		return -1
	}

	// count 表示执行反转非相邻字符数，偶数时合法
	// change 表示当前字符是否被右边字符反转
	var dfs func(i int, count int, change bool) int
	dfs = func(i int, count int, change bool) (res int) {
		if i == -1 {
			// 不合法的方案：
			// 1. s1[0] 左边没有字符供反转
			// 2.反转的次数必须是偶数次
			if change || count%2 != 0 {
				return math.MaxInt32
			}
			return count / 2 * x
		}

		if (s1[i] == s2[i]) == !change {
			return dfs(i-1, count, false) // 字符一致，无需反转
		}
		return min(
			dfs(i-1, count, true)+1,  // 和相邻字符反转
			dfs(i-1, count+1, false), // 和非相邻字符反转
		)
	}

	MemorySearch3(&dfs)
	ans := dfs(len(s1)-1, 0, false)
	return ans
}

// dfs + 记忆化搜索
//
// - 时间复杂度：$\mathcal{O}(n)$。
// - 空间复杂度：$\mathcal{O}(n)$。
func minOperations2(s1 string, s2 string, x int) int {
	var index []int
	for i := 0; i < len(s1); i++ {
		if s1[i] != s2[i] {
			index = append(index, i)
		}
	}
	m := len(index)
	if m%2 != 0 {
		return -1
	}

	var dfs func(i int) int
	dfs = func(i int) int { // 写法一
		if i == -1 {
			return 0
		}

		res := dfs(i-1) + x // 第一种操作：反转非相邻字符
		if i >= 1 {         // i == 0 时，无法进行反转相邻字符串
			res = min(res, dfs(i-2)+(index[i]-index[i-1])*2) // 第二种操作：反转相邻字符串
		}
		return res
	}

	dfs = func(i int) int { // 写法二
		if i == 0 {
			return dfs(i-1) + x // i == 0 时，无法进行反转相邻字符串
		}
		if i == -1 {
			return 0
		}

		return min(
			dfs(i-1)+x,                       // 第一种操作：反转非相邻字符
			dfs(i-2)+(index[i]-index[i-1])*2, // 第二种操作：反转相邻字符
		)
	}
	MemorySearch(&dfs)
	ans := dfs(m-1) / 2
	return ans
}

// 将 minOperations2 1:1 翻译为 dp
//
// - 时间复杂度：$\mathcal{O}(n)$。
// - 空间复杂度：$\mathcal{O}(n)$。
func minOperations3(s1 string, s2 string, x int) int {
	var index []int
	for i := 0; i < len(s1); i++ {
		if s1[i] != s2[i] {
			index = append(index, i)
		}
	}
	m := len(index)
	if m%2 != 0 {
		return -1
	}
	if m == 0 {
		return 0
	}

	dp := make([]int, m+1) // 也可以改成 m+2，用于去掉 m == 0 的判断
	dp[0] = 0
	dp[1] = x
	for i := 1; i < m; i++ {
		dp[i+1] = min(
			dp[i]+x,                         // 第一种操作：反转非相邻字符
			dp[i-1]+(index[i]-index[i-1])*2, // 第二种操作：反转相邻字符
		)
	}
	return dp[m] / 2
}
