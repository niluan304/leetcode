package main

import (
	"math"

	. "github.com/niluan304/leetcode/copypasta/graph"
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

func findTheCity2(n int, edges [][]int, distanceThreshold int) int {
	path := FloydDFS(n, func(path [][]int) {
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
