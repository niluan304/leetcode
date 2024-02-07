package main

import . "github.com/niluan304/leetcode/container"

func maximumScoreAfterOperations(edges [][]int, values []int) int64 {
	var n = len(values)
	var graph = make([][]int, n)
	for _, edge := range edges {
		a, b := edge[0], edge[1]

		graph[a] = append(graph[a], b)
		graph[b] = append(graph[b], a)
	}

	// dfs(i, fa) 计算以 i 为根的子树是健康时，失去的最小分数
	var dfs func(i, fa int) int
	dfs = func(i, fa int) int {
		cost := 0
		for _, j := range graph[i] {
			if j == fa {
				continue
			}
			cost += dfs(j, i)
		}
		if cost == 0 {
			return values[i] // 叶子节点
		}
		return min(cost, values[i])
	}

	total := Sum(values)
	minCost := dfs(0, -1)
	return int64(total - minCost)
}
