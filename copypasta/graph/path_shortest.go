package graph

import (
	"math"

	. "github.com/niluan304/leetcode/container"
)

// Floyd
// 是一种多源点最短路径算法，返回任意两点间的最短距离。
//
// 可以处理有向图或负数边权（但不支持负权回路）
//
// Floyd算法 又称为插点法，寻找给定的加权图中多源点之间最短路径的算法。
// Floyd算法的本质是动态规划
// 最大距离初始化为 math.MaxInt/2
//
// 复杂度
// - 时间复杂度：O(n^3)。
// - 空间复杂度：O(n^2)。
func Floyd(n int, init func(path [][]int)) [][]int {
	path := make([][]int, n)
	for i := range path {
		path[i] = make([]int, n)
		for j := range path[i] {
			path[i][j] = math.MaxInt / 2 // 初始化最大距离， /2 防止加法溢出
		}
		path[i][i] = 0
	}

	// 设置 path[i][j] 两点的边权
	init(path)

	// 迭代更新以 [0, n-1] 为中转节点的最短距离
	for mid := 0; mid < n; mid++ { // mid 为中转节点
		for i := 0; i < n; i++ {
			if mid == i {
				continue
			}
			for j := 0; j < n; j++ {
				if mid == j {
					continue
				}
				path[i][j] = min(path[i][j], path[i][mid]+path[mid][j])
			}
		}
	}
	return path
}

// FloydDFS
// Floyd 的本质是动态规划
// [带你发明 Floyd 算法：从记忆化搜索到递推（Python/Java/C++/Go/JS/Rust）](https://leetcode.cn/problems/find-the-city-with-the-smallest-number-of-neighbors-at-a-threshold-distance/solutions/2525946/dai-ni-fa-ming-floyd-suan-fa-cong-ji-yi-m8s51)
//
// ## 题意解读
//
// 对于城市 $i$，求出 $i$ 到其余城市的最短路长度，统计有多少个最短路长度 $\le \textit{distanceThreshold}$，把个数记作 $\textit{cnt}_i$。
//
// 例如示例 1，$\textit{cnt}_0=2,\textit{cnt}_1=3,\textit{cnt}_2=3,\textit{cnt}_3=2$。
//
// 其中 $\textit{cnt}_i$ 最小的 $i$ 就是答案。示例 1 中的城市 $0$ 和城市 $3$ 对应的 $\textit{cnt}_i$ 都是最小的，这种情况返回 $0$ 和 $3$ 的最大值，即 $3$。
//
// 为了解决本题，我们需要求出**任意两个城市之间的最短路长度**。
//
// ## 前置知识：动态规划入门
//
// 请看视频讲解 `b23.tv/pc522x3`
//
// ## 一、启发思考：寻找子问题
//
// ![lc1334-c.png](https://pic.leetcode.cn/1699881511-DbyZrc-lc1334-c.png)
//
// ## 二、递归怎么写：状态定义与状态转移方程
//
// 定义 $\textit{dfs}(k,i,j)$ 表示从 $i$ 到 $j$ 的最短路长度，并且这条最短路的中间节点编号都 $\le k$。注意中间节点不包含 $i$ 和 $j$。
//
// 根据上面的讨论：
//
// - 不选 $k$，那么中间节点的编号都 $\le k-1$，即 $\textit{dfs}(k,i,j) = \textit{dfs}(k-1,i,j)$。
// - 选 $k$，问题分解成从 $i$ 到 $k$ 的最短路，以及从 $k$ 到 $j$ 的最短路。由于这两条最短路的中间节点都不包含 $k$，所以中间节点的编号都 $\le k-1$，故得到 $\textit{dfs}(k,i,j)=\textit{dfs}(k-1,i,k)+\textit{dfs}(k-1,k,j)$。
//
// 这两种情况取最小值，就得到了 $\textit{dfs}(k,i,j)$。写成式子就是
//
// $$
// \textit{dfs}(k,i,j) = \min (\textit{dfs}(k-1,i,j), \textit{dfs}(k-1,i,k)+\textit{dfs}(k-1,k,j))
// $$
//
// 递归边界：$\textit{dfs}(-1,i,j) = w[i][j]$。$k=-1$ 表示 $i$ 和 $j$ 之间没有任何中间节点，此时最短路长度只能是连接 $i$ 和 $j$ 的边的边权，即 $w[i][j]$。如果没有连接 $i$ 和 $j$ 的边，则 $w[i][j]=\infty$。
//
// 递归入口：$\textit{dfs}(n-1,i,j)$，表示从 $i$ 到 $j$ 的最短路长度。$k=n-1$ 是因为本题节点编号为 $0$ 到 $n-1$，任意最短路的中间节点编号都 $\le n-1$。
// ### 复杂度分析
//
// - 时间复杂度：$\mathcal{O}(n^3)$。
// - 空间复杂度：$\mathcal{O}(n^2)$。
//
// ## 思考题
//
// 向图中添加一条边，如何维护 $f$ 数组？
//
// 最暴力的做法是，每次添加一条边，就用 $\mathcal{O}(n^3)$ 时间全部重算一遍。有没有更快的做法呢？
//
// 这题是 [2642. 设计可以求最短路径的图类](https://leetcode.cn/problems/design-graph-with-shortest-path-calculator/)
//
// ---
//
// 欢迎关注 [B站@灵茶山艾府](https://space.bilibili.com/206214)
//
// 更多精彩题解，请看 [往期题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)
func FloydDFS(n int, init func(path [][]int)) [][]int {
	path := make([][]int, n)
	for i := range path {
		path[i] = make([]int, n)
		for j := range path[i] {
			path[i][j] = math.MaxInt / 2 // 初始化最大距离， /2 防止加法溢出
		}
		path[i][i] = 0
	}

	// 设置 path[i][j] 两点的边权
	init(path)

	cache := make([][][]*int, n)
	for i := 0; i < n; i++ {
		cache[i] = make([][]*int, n)
		for j := 0; j < n; j++ {
			cache[i][j] = make([]*int, n)
		}
	}

	// dfs(k,i,j)表示 i-j 的最短路且路径中节点的编号不超过 k 的 min
	// dfs(k,i,j) = min(dfs(k-1,i,j), dfs(k-1,i,k)+dfs(k-1,k,j))
	var dfs func(k, i, j int) int
	dfs = func(k, i, j int) (res int) {
		if k == -1 {
			return path[i][j]
		}

		ptr := &cache[i][j][k]
		if *ptr != nil {
			return **ptr
		}
		defer func() { *ptr = &res }()

		return min(
			dfs(k-1, i, j),
			dfs(k-1, i, k)+dfs(k-1, k, j),
		)
	}

	// 迭代更新以 [0, n-1] 为中转节点的最短距离
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if i != j {
				path[i][j] = dfs(n-1, i, j)
			}
		}
	}
	return path
}

type DijkstraEdge struct {
	To     int
	Weight int
}

// Dijkstra
// 单源最短路径算法，返回起点 start 到其他节点的最短路径
// Dijkstra 算法是一种贪心算法，不支持负数边权
//
// 暴力解法求最短路长度最小的节点，适用于稠密图
//
// 复杂度
// - 时间复杂度：O(n^2)。
// - 空间复杂度：O(n^2)。
func Dijkstra(n int, start int, init func(graph [][]DijkstraEdge)) (distance, from []int) {
	const INF = math.MaxInt / 2 // /2 防止加法溢出
	graph := make([][]DijkstraEdge, n)

	// 初始化 graph[i] 的邻接表
	init(graph)

	visit := make([]bool, n)
	from = make([]int, n)
	distance = make([]int, n)
	for i := 0; i < n; i++ {
		distance[i] = INF //  初始化最大距离
		from[i] = -1      // -1 表示无法到达
	}
	distance[start] = 0 // 初始化

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
				from[i] = cur
			}
		}
	}
	return distance, from
}

// DijkstraPriorityQueue
// 单源最短路径算法，返回起点 start 到其他节点的最短路径
// Dijkstra 算法是一种贪心算法，不支持负数边权
//
// 通过「优先队列」维护最短路长度最小的节点，适合稀疏图
//
// 复杂度
// - 时间复杂度：O(mlogm)。
// - 空间复杂度：O(m)。
func DijkstraPriorityQueue(n int, start int, init func(graph [][]DijkstraEdge)) (distance, from []int) {
	const INF = math.MaxInt / 2 // /2 防止加法溢出
	graph := make([][]DijkstraEdge, n)

	// 初始化 graph[i] 的邻接表
	init(graph)

	distance = make([]int, n)
	from = make([]int, n)
	for i := 0; i < n; i++ {
		distance[i] = INF // 初始化最大距离
		from[i] = -1      // -1 表示无法到达
	}
	distance[start] = 0 // 初始化

	pq := &PriorityQueue[int]{} // 优先队列，最大堆
	pq.Insert(start, 0)

	// 计算从起始节点到所有其他节点的最短距离
	for pq.Len() > 0 {
		// 在未访问节点中，找到距离起始节点最近的节点
		head := pq.PopHead()
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
				from[i] = cur
				pq.Insert(i, -d) // 最大堆，取反后就是最小堆了
			}
		}
	}

	return distance, from
}
