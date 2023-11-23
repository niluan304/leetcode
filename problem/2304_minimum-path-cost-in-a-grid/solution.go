package main

import (
	"math"
	"slices"
)

// Deprecated: Fail to pass
// 解题思路未错误，但代码实现时，为优化空间复杂度，导致编码解答错误。
func minPathCost(grid [][]int, moveCost [][]int) int {
	var n, m = len(grid), len(grid[0])
	var costs = make([]int, m)
	for _, row := range grid[:n-1] {
		for i := 0; i < m; i++ {
			var minCost = math.MaxInt32
			for _, v := range row {
				minCost = min(minCost, v+moveCost[v][i])
			}
			costs[i] += minCost
		}
	}

	var ans = math.MaxInt32
	for i, v := range grid[n-1] {
		ans = min(ans, costs[i]+v)
		costs[i] += v
	}
	return ans
}

func minPathCost2(grid [][]int, moveCost [][]int) int {
	var m, n = len(grid), len(grid[0])
	var cache = make([][]*int, m)
	for i, _ := range cache {
		cache[i] = make([]*int, n)
	}

	var dfs func(i, j int) int
	dfs = func(i, j int) int {
		if i == m-1 { // 递归边界
			return grid[i][j]
		}
		ptr := &cache[i][j]
		if *ptr != nil { // 之前计算过
			return **ptr
		}
		res := math.MaxInt
		for k, c := range moveCost[grid[i][j]] { // 移动到下一行的第 k 列
			res = min(res, dfs(i+1, k)+c+grid[i][j])
		}
		*ptr = &res // 记忆化
		return **ptr
	}

	var ans = math.MaxInt32
	for j := 0; j < n; j++ { // 枚举起点
		ans = min(ans, dfs(0, j))
	}
	return ans
}

func minPathCost3(grid [][]int, moveCost [][]int) int {
	var m, n = len(grid), len(grid[0])
	var dp = make([][]int, m)
	for i, _ := range dp {
		dp[i] = make([]int, n)
	}

	for j, v := range grid[0] {
		dp[0][j] = v
	}

	for i, row := range grid[:m-1] {
		for x, v := range row {
			for j, cost := range moveCost[v] {
				cost += dp[i][x] + grid[i+1][j]
				if x == 0 {
					dp[i+1][j] = cost
				} else {
					dp[i+1][j] = min(dp[i+1][j], cost)
				}
			}
		}
	}
	return slices.Min(dp[m-1])
}
