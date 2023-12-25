package main

import (
	"math"
)

func findTheCity(n int, edges [][]int, distanceThreshold int) int {
	path := Floyd(n, func(path [][]int) {
		for _, edge := range edges {
			i, j, v := edge[0], edge[1], edge[2]
			path[i][j] = v
			path[j][i] = v
		}
	})

	var node, count = -1, math.MaxInt32
	for i, distance := range path {
		cnt := -1 // 用于比较 i == j, d = 0
		for _, d := range distance {
			if d <= distanceThreshold {
				cnt++
			}
		}
		if cnt <= count {
			count = cnt
			node = i
		}
	}
	return node
}

func Floyd(n int, initPath func(path [][]int)) [][]int {
	var path = make([][]int, n)
	for i, _ := range path {
		path[i] = make([]int, n)
		for j, _ := range path[i] {
			path[i][j] = math.MaxInt32 // 初始化最大距离
		}
		path[i][i] = 0
	}

	// 初始化 path-1，两点直达距离
	initPath(path)

	// 迭代更新以 [0, n-1] 为中转节点的最短距离
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if i == j {
				continue
			}
			for k := 0; k < n; k++ {
				if i == k {
					continue
				}
				path[j][k] = min(path[j][k], path[j][i]+path[i][k])
			}
		}
	}
	return path
}
