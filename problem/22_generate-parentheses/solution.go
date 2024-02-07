package main

import (
	"slices"

	. "github.com/niluan304/leetcode/copypasta/dp"
)

func generateParenthesis(n int) []string {
	switch n {
	case 0:
		return []string{}
	case 1:
		return []string{"()"}
	case 2:
		return []string{"(())", "()()"}
	}

	lastSum := generateParenthesis(n - 1)
	var m = map[string]struct{}{}
	var curSum = make([]string, 0, len(lastSum)*3)
	for _, last := range lastSum {
		for i := 0; i <= len(last); i += 2 {
			for j := i; j <= len(last); j += 2 {
				cur := last[:i] + "(" + last[i:j] + ")" + last[j:]
				if _, ok := m[cur]; !ok {
					curSum = append(curSum, cur)
					m[cur] = struct{}{}
				}
			}
		}
	}
	return curSum
}

// leetcode 2 动态规划 https://leetcode.cn/problems/generate-parentheses/solutions/9251/zui-jian-dan-yi-dong-de-dong-tai-gui-hua-bu-lun-da/
func leetcode2(n int) []string {
	return dp(n)[n]
}

func dp(n int) map[int][]string {
	if n == 0 {
		return map[int][]string{0: {""}}
	}
	if n == 1 {
		return map[int][]string{0: {""}, 1: {"()"}}
	}
	lastMap := dp(n - 1)
	var oneRes []string
	for i := 0; i < n; i++ {
		inners := lastMap[i]
		outers := lastMap[n-1-i]
		for _, inner := range inners {
			for _, outer := range outers {
				oneRes = append(oneRes, "("+inner+")"+outer)
			}
		}
	}
	lastMap[n] = oneRes
	return lastMap
}

// 递归版本的有效括号定义
// 定义的 dfs 函数为纯函数
func generateParenthesis2(n int) []string {
	// 1.空字符串 "" 是一个合法的括号
	// 2. 如果 "x" 是合法的括号，那么 "(x)" 也是合法的括号
	// 3. 如果 "x","y" 都是合法的括号，那么 "xy" 也是合法的括号

	var dfs func(n int) []string
	dfs = func(n int) []string {
		if n == 0 {
			return []string{""} // 1. 空字符串 "" 是一个合法的括号
		}

		var set = map[string]struct{}{}
		for _, value := range dfs(n - 1) {
			set["("+value+")"] = struct{}{} // 2. 如果 "x" 是合法的括号，那么 "(x)" 也是合法的括号
		}

		for i := 1; i < n/2+1; i++ {
			for _, x := range dfs(i) {
				for _, y := range dfs(n - i) {
					// 3. 如果 "x","y" 都是合法的括号，那么 "xy" 也是合法的括号
					set[y+x] = struct{}{}
					set[x+y] = struct{}{}
				}
			}
		}
		return MapKeys(set)
	}

	MemorySearch(&dfs)

	ans := dfs(n)
	slices.Sort(ans[:]) // 额外做升序排列，仅用于 debug
	return ans
}

func MapKeys[M ~map[K]V, K comparable, V any](m M) (keys []K) {
	for k, _ := range m {
		keys = append(keys, k)
	}
	return keys
}
