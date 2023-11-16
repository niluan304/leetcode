package main

import (
	"math"
	"strings"
)

func wordBreak(s string, wordDict []string) bool {
	type Set struct {
		set      map[string]struct{}
		min, max int
	}

	var set = Set{
		set: make(map[string]struct{}, len(wordDict)),
		min: math.MaxInt,
		max: 0,
	}

	for _, item := range wordDict {
		set.set[item] = struct{}{}
		set.min = min(set.min, len(item))
		set.max = max(set.max, len(item))
	}

	var dfs func(s string, set Set, idx int) bool
	dfs = func(s string, set Set, idx int) bool {
		if len(s) == idx {
			return true
		}

		left, right := idx+set.min, min(len(s), idx+set.max)
		for i := left; i <= right; i++ {
			if _, ok := set.set[s[idx:i]]; !ok {
				continue
			}

			if dfs(s, set, i) {
				return true
			}
		}

		return false
	}

	return dfs(s, set, 0)
}

func max(x, y int) int {
	if x > y {
		return x
	}

	return y
}

func min(x, y int) int {
	if x < y {
		return x
	}

	return y
}

func wordBreak2(s string, wordDict []string) bool {
	var bound = map[int][]int{}
	for _, word := range wordDict {
		n := len(word)
		idx := 0
		for {
			i := strings.Index(s[idx:], word)
			if i == -1 {
				break
			}
			bound[idx+i] = append(bound[idx+i], idx+i+n)
			idx += i + n
		}
	}

	var n = len(s)
	var cache = make([]bool, n)
	var dfs func(i int) bool
	dfs = func(i int) bool {
		if i == n {
			return true
		}
		if cache[i] {
			return false
		}
		for _, idx := range bound[i] {
			if dfs(idx) {
				return true
			}
		}

		cache[i] = true
		return false
	}

	return dfs(0)
}

// dfs + 记忆化搜索
// 时间复杂度：O(m*n)
// 空间复杂度：O(n)
func wordBreak3(s string, wordDict []string) bool {
	var n = len(s)
	var cache = make([]bool, n)

	var dfs func(i int) bool
	dfs = func(i int) bool {
		if i == n {
			return true
		}
		if cache[i] {
			return false
		}

		for _, word := range wordDict {
			if !strings.HasPrefix(s[i:], word) {
				continue
			}
			if dfs(i + len(word)) {
				return true
			}
		}
		cache[i] = true
		return false
	}
	return dfs(0)
}

// dfs + 记忆化搜索 翻译为 动态规划
// 时间复杂度：O(m*n)
// 空间复杂度：O(n)
func wordBreak4(s string, wordDict []string) bool {
	var (
		n  = len(s)
		dp = make([]bool, n+1)
	)
	dp[0] = true

	for i := 0; i < n; i++ {
		for _, word := range wordDict {
			j := i + len(word)
			if j > n || !strings.HasPrefix(s[i:], word) {
				continue
			}
			dp[j] = dp[j] || dp[i]
		}
	}

	return dp[n]
}

// dfs + 记忆化搜索
// 时间复杂度：O(m*n)
// 空间复杂度：O(n)
// 死循环
func wordBreak5(s string, wordDict []string) bool {
	var (
		n     = len(s)
		cache = make([]bool, n+1)
	)
	cache[0] = true

	var dfs func(i int) bool
	dfs = func(i int) bool {
		if cache[i] {
			return true
		}

		for _, word := range wordDict {
			j := i - len(word)
			// 自顶向下：strings.HasSuffix
			// 自底向上：strings.HasPrefix
			// 一个只能判断后缀，一个只能判断前缀
			if j < 0 || !strings.HasSuffix(s[:i], word) {
				continue
			}

			cache[i] = cache[i] || dfs(j)
		}
		return cache[i]
	}

	return dfs(n)
}
