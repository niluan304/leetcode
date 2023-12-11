package main

import (
	"math"
	"sort"
)

// 回溯
// 时间复杂度: O(2^(m*n))
// 空间复杂度: O(n)
// Deprecated: 超时
func minimumEffortPath(heights [][]int) int {
	m, n := len(heights), len(heights[0])
	var ans = math.MaxInt32
	var visit = make([][]bool, m)
	for i, _ := range visit {
		visit[i] = make([]bool, n)
	}
	visit[0][0] = true
	type Point struct {
		x, y int
	}
	var path = [][2]int{{heights[0][0], 0}}

	var dfs func(i, j int)
	dfs = func(i, j int) {
		if i == m-1 && j == n-1 {
			ans = min(ans, path[len(path)-1][1])
		}

		visit[i][j] = true
		next := []Point{
			{x: i - 1, y: j},
			{x: i + 1, y: j},
			{x: i, y: j - 1},
			{x: i, y: j + 1},
		}
		for _, point := range next {
			x, y := point.x, point.y
			if x < 0 || y < 0 || x >= m || y >= n || visit[x][y] {
				continue
			}

			tail := path[len(path)-1]
			path = append(path, [2]int{
				heights[x][y],
				max(tail[1], abs(tail[0]-heights[x][y])),
			})
			visit[x][y] = true

			dfs(x, y)

			visit[x][y] = false
			path = path[:len(path)-1]
		}
	}
	dfs(0, 0)
	return ans
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// 二分查找 + dfs
// 时间复杂度: O(mnlogC)
// 空间复杂度: O(mn)
func minimumEffortPath2(heights [][]int) int {
	type Value struct{ x, y int }
	var dirs = []Value{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	m, n := len(heights), len(heights[0])

	return sort.Search(1e6, func(maxHeightDiff int) bool {
		visit := make([][]bool, m)
		for i := range visit {
			visit[i] = make([]bool, n)
		}
		visit[0][0] = true

		var dfs func(i, j int) bool
		dfs = func(i, j int) bool {
			if i == m-1 && j == n-1 {
				return true
			}
			for _, d := range dirs {
				x, y := i+d.x, j+d.y
				if x < 0 || y < 0 || x >= m || y >= n ||
					visit[x][y] || abs(heights[x][y]-heights[i][j]) > maxHeightDiff {
					continue
				}
				visit[x][y] = true
				if dfs(x, y) {
					return true
				}
			}
			return false
		}
		return dfs(0, 0)
	})
}
