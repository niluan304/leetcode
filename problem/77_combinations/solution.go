package main

import "slices"

func combine(n int, k int) [][]int {
	var ans [][]int
	var path []int
	var dfs func(i int)
	dfs = func(i int) {
		if len(path) == k {
			ans = append(ans, slices.Clone(path))
			return
		}
		if len(path)+(n-i+1) < k { // 长度不足，剪枝
			return
		}

		for j := i; j <= n; j++ {
			path = append(path, j)
			dfs(j + 1)
			path = path[:len(path)-1]
		}
	}

	dfs(1)
	return ans
}
