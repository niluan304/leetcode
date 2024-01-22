package main

import (
	"math"
)

// 最短距离必定是 3种方式之一：
//
// 1. 直接由 i 到 j
// 2. i 先到 x ，转移至 y ，再到 j
// 3. i 先到 y ，转移至 x ，再到 j
//
// - 时间复杂度：$\mathcal{O}(n^2)$。
// - 空间复杂度：$\mathcal{O}(n^2)$。
func countOfPairs(n int, x int, y int) []int {
	var ans = make([]int, n)

	// 暴力枚举每对房屋最短距离
	for i := 1; i < n+1; i++ {
		for j := i + 1; j < n+1; j++ {
			// 最短距离必定是 3 种方式中的一种
			minDist := min(
				j-i,                 // i 直接去 j
				Abs(i-x)+1+Abs(j-y), // i 通过 x 去到 y 再到 j
				Abs(i-y)+1+Abs(j-x), // i 通过 y 去到 x 再到 j
			)
			ans[minDist-1] += 2 // 对称性：i 到 j 的最短距离，同样是 j 到 i 的最短距离，因此 +2
		}
	}

	return ans
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// Floyd算法
// 数据范围足够小：2 <= n <= 100，可以使用 Floyd算法 求解多点距离
//
// - 时间复杂度：$\mathcal{O}(n^3)$。
// - 空间复杂度：$\mathcal{O}(n^2)$。
func countOfPairs2(n int, x int, y int) []int {
	path := Floyd(n, func(path [][]int) {
		for i := 0; i < n-1; i++ {
			path[i][i+1] = 1
			path[i+1][i] = 1
		}

		if x == y {
			return
		}

		path[x-1][y-1] = 1
		path[y-1][x-1] = 1
	})

	var ans = make([]int, n+1)
	for i, _ := range path {
		for _, v := range path[i] {
			ans[v]++
		}
	}
	return ans[1:]
}

// Floyd 返回任意两点间的最短距离
// Floyd算法 又称为插点法，寻找给定的加权图中多源点之间最短路径的算法，与 Dijkstra算法 类似。
// Floyd算法利用动态规划的思想，Dijkstra算法则是贪心。
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

// BFS
// - 时间复杂度：$\mathcal{O}(n^2)$。
// - 空间复杂度：$\mathcal{O}(n^2)$。
func countOfPairs3(n int, x int, y int) []int {
	graph := make([][]int, n)
	for i := 0; i < n-1; i++ {
		graph[i] = append(graph[i], i+1)
		graph[i+1] = append(graph[i+1], i)
	}
	graph[x-1] = append(graph[x-1], y-1)
	graph[y-1] = append(graph[y-1], x-1)

	type pair struct{ index, distance int }
	var ans = make([]int, n)

	var bfs = func(i int) {
		var visit = make([]bool, n)
		visit[i] = true

		queue := []pair{{index: i, distance: 0}}
		for len(queue) > 0 {
			head := queue[0]
			queue = queue[1:]

			idx, dist := head.index, head.distance
			for _, j := range graph[idx] {
				if visit[j] {
					continue
				}
				visit[j] = true

				queue = append(queue, pair{index: j, distance: dist + 1})
				ans[dist]++
			}
		}
	}
	for i := 0; i < n; i++ {
		bfs(i)
	}
	return ans
}
