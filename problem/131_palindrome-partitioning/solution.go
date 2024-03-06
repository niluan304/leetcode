package main

import "slices"

// 回溯模板2
// 答案的视角（枚举子串结束位置）
// - 时间复杂度：$\mathcal{O}(n \cdot 2^n)$。
// - 空间复杂度：$\mathcal{O}(n)$。
func partition(s string) (ans [][]string) {
	var path []string
	n := len(s)

	var dfs func(i int)
	dfs = func(i int) {
		if i == n {
			ans = append(ans, slices.Clone(path))
			return
		}

		for j := i; j < n; j++ {
			sub := s[i : j+1]
			if !isPartition(sub) {
				continue
			}

			path = append(path, sub)
			dfs(j + 1)
			path = path[:len(path)-1]
		}
	}

	dfs(0)
	slices.SortFunc(ans, func(a, b []string) int {
		return len(b) - len(a)
	})
	return ans
}

func isPartition(s string) bool {
	n := len(s)
	for i := 0; i < n/2; i++ {
		if s[i] != s[n-1-i] {
			return false
		}
	}
	return true
}

// 回溯模板1
// 输入的视角（逗号选或不选）
// - 时间复杂度：$\mathcal{O}(n \cdot 2^n)$。
// - 空间复杂度：$\mathcal{O}(n)$。
func partition2(s string) (ans [][]string) {
	var path []string
	n := len(s)

	var dfs func(start, end int)
	dfs = func(start, end int) {
		if end == n {
			ans = append(ans, slices.Clone(path))
			return
		}

		if end < n-1 {
			dfs(start, end+1) // 不选
		}

		sub := s[start : end+1]
		if !isPartition(sub) {
			return
		}
		path = append(path, sub)
		dfs(end+1, end+1) // 选
		path = path[:len(path)-1]
	}

	dfs(0, 0)
	slices.SortFunc(ans, func(a, b []string) int {
		return len(b) - len(a)
	})
	return ans
}
