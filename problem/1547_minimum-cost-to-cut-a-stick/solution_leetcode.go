package main

import (
	"math"
	"sort"
)

// 4ms 代码示例
// 木棍长度, 切割点, 如何找到切割点所在木棍长度
// 切一刀就变成左右两边的子问题
func leetcodeMinTime(n int, cuts []int) int {
	cuts = append(cuts, 0, n)
	sort.Ints(cuts)

	// 创建 dp：存在可计算因子
	m := len(cuts)
	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, m)
	}
	for i := m - 1; i >= 0; i-- {
		for j := i + 2; j < m; j++ {
			dp[i][j] = math.MaxInt
			// 枚举分割点
			for k := i + 1; k < j; k++ {
				dp[i][j] = min(dp[i][j], dp[i][k]+dp[k][j]+cuts[j]-cuts[i])
			}
		}
	}
	return dp[0][m-1]
}
