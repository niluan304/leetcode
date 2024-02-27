package main

import (
	"math"

	. "github.com/niluan304/leetcode/copypasta/dp"
)

// dfs + 记忆化搜索
//
// - 时间复杂度：$\mathcal{O}(m \cdot n)$。
// - 空间复杂度：$\mathcal{O}(m \cdot n)$。
func minPathSum(grid [][]int) int {
	var dfs func(i, j int) int
	dfs = func(i, j int) int {
		if i < 0 || j < 0 {
			return math.MaxInt32
		}
		if i == 0 && j == 0 {
			return grid[0][0]
		}
		// 题意设定：每次只能向下或者向右移动一步。
		return min(dfs(i-1, j), dfs(i, j-1)) + grid[i][j]
	}

	MemorySearch2(&dfs)
	return dfs(len(grid)-1, len(grid[0])-1)
}

// 将 minPathSum 1:1 翻译为递推
//
// - 时间复杂度：$\mathcal{O}(m \cdot n)$。
// - 空间复杂度：$\mathcal{O}(m \cdot n)$。
func minPathSum2(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
		dp[i][0] = math.MaxInt32
	}
	for i := range dp[0] {
		dp[0][i] = math.MaxInt32
	}
	dp[0][1] = 0

	for i, row := range grid {
		for j, v := range row {
			dp[i+1][j+1] = min(dp[i+1][j], dp[i][j+1]) + v // 题意设定：每次只能向下或者向右移动一步。
		}
	}
	return dp[m][n]
}

// 如果移动方向改为上下左右，该如何操作？
