package main

import (
	"math"

	. "github.com/niluan304/leetcode/container"
)

// Floyd + dp
//
// - 时间复杂度：O(n^3)。
// - 空间复杂度：O(n^2)。
func countPaths(n int, roads [][]int) int {
	const MOD = 1_000_000_007

	type Pair struct {
		Distance int
		Count    int
	}

	init := func(path [][]Pair) {
		// 初始化 path-1，两点直达距离
		for _, road := range roads {
			u, v, w := road[0], road[1], road[2]
			path[u][v] = Pair{Distance: w, Count: 1}
			path[v][u] = Pair{Distance: w, Count: 1}
		}
	}

	// 正常的 Floyd算法签名：func Floyd (n int, init func(path [][]int)) (path [][]int)
	path := make([][]Pair, n)
	for i := range path {
		path[i] = make([]Pair, n)
		for j := range path[i] {
			path[i][j] = Pair{
				Distance: math.MaxInt / 2, // 初始化最大距离,
				Count:    0,
			}
		}
		path[i][i] = Pair{
			Distance: 0,
			Count:    1,
		}
	}

	init(path)

	// 迭代更新以 [0, n-1] 为中转节点的最短距离
	for mid := 0; mid < n; mid++ {
		for i := 0; i < n; i++ {
			if mid == i {
				continue
			}
			for j := 0; j < n; j++ {
				if mid == j {
					continue
				}
				old := path[i][j]

				x, y := path[i][mid], path[mid][j]
				cur := Pair{
					Distance: x.Distance + y.Distance,
					Count:    x.Count * y.Count,
				}

				if cur.Distance < old.Distance {
					path[i][j] = cur
				} else if cur.Distance == old.Distance {
					path[i][j].Count = (old.Count + cur.Count) % MOD
				}
			}
		}
	}

	return path[0][n-1].Count
}

// Dijkstra + dp
//
// - 时间复杂度：O(n^2)。
// - 空间复杂度：O(n^2)。
func countPaths2(n int, roads [][]int) int {
	const MOD = 1_000_000_007
	type DijkstraEdge struct {
		To     int
		Weight int
	}

	start := 0
	init := func(graph [][]DijkstraEdge) {
		for _, road := range roads {
			u, v, w := road[0], road[1], road[2]
			graph[u] = append(graph[u], DijkstraEdge{To: v, Weight: w})
			graph[v] = append(graph[v], DijkstraEdge{To: u, Weight: w})
		}
	}

	// Dijkstra算法
	graph := make([][]DijkstraEdge, n)

	// 初始化 graph[i] 的邻居和边权
	init(graph)

	visit := make([]bool, n)
	distance := make([]int, n)
	dp := make([]int, n)
	for i := 0; i < n; i++ {
		distance[i] = math.MaxInt / 2 //  初始化最大距离
	}
	distance[start] = 0 // 初始化
	dp[start] = 1       // 初始化

	// 计算从起始节点到所有其他节点的最短距离
	for {
		cur := -1 // 初始化

		// 在未访问节点中，找到距离起始节点最近的节点
		for i, ok := range visit {
			if !ok && (cur == -1 || distance[i] < distance[cur]) {
				cur = i
			}
		}

		// 如果没有找到最近的未访问节点，则退出循环
		if cur == -1 {
			break
		}

		// 标记当前节点为已访问
		visit[cur] = true

		// 更新当前节点邻居到起点的最短距离
		for _, edge := range graph[cur] {
			i := edge.To
			d := edge.Weight + distance[cur]

			if distance[i] > d {
				distance[i] = d
				dp[i] = dp[cur]
			} else if distance[i] == d {
				dp[i] = (dp[i] + dp[cur]) % MOD
			}
		}
	}
	return dp[n-1]
}

// DijkstraPriorityQueue + dp
//
// - 时间复杂度：O(m logm)。
// - 空间复杂度：O(m)。
func countPaths3(n int, roads [][]int) int {
	const MOD = 1_000_000_007
	type DijkstraEdge struct {
		To     int
		Weight int
	}

	start := 0
	init := func(graph [][]DijkstraEdge) {
		for _, road := range roads {
			u, v, w := road[0], road[1], road[2]
			graph[u] = append(graph[u], DijkstraEdge{To: v, Weight: w})
			graph[v] = append(graph[v], DijkstraEdge{To: u, Weight: w})
		}
	}

	// Dijkstra算法
	graph := make([][]DijkstraEdge, n)

	// 初始化 graph[i] 的邻居和边权
	init(graph)

	distance := make([]int, n)
	dp := make([]int, n)
	for i := 0; i < n; i++ {
		distance[i] = math.MaxInt / 2 //  初始化最大距离
	}
	distance[start] = 0 // 初始化
	dp[start] = 1       // 初始化

	pq := &PriorityQueue[int]{} // 优先队列，最大堆
	pq.Insert(start, 0)
	// 计算从起始节点到所有其他节点的最短距离
	for pq.Len() > 0 {
		// 在未访问节点中，找到距离起始节点最近的节点
		head := pq.Remove()
		cur := head.Value()

		// 下面循环中的 d < distance[i] 可能会把重复的节点 i 入堆
		// 也就是说，堆中可能会包含多个相同节点，且这些相同节点的 distance 值互不相同
		// 那么这个节点第二次及其后面出堆的时候，由于 distance[cur] 已经更新成最短路了，可以直接跳过
		if d := -head.Priority(); d > distance[cur] {
			continue
		}

		for _, edge := range graph[cur] {
			i := edge.To
			d := edge.Weight + distance[cur]

			if distance[i] > d {
				distance[i] = d
				dp[i] = dp[cur]

				pq.Insert(i, -d) // 最大堆，取反后就是最小堆了
			} else if distance[i] == d {
				dp[i] = (dp[i] + dp[cur]) % MOD
			}
		}
	}
	return dp[n-1]
}
