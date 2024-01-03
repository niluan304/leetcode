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

// Floyd 返回任意两点间的最短距离
// Floyd算法 又称为插点法，一种利用动态规划的思想寻找给定的加权图中多源点之间最短路径的算法，与 Dijkstra算法 类似。
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

func findTheCity2(n int, edges [][]int, distanceThreshold int) int {
	path := Floyd2(n, func(path [][]int) {
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

// Floyd2
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
func Floyd2(n int, initPath func(path [][]int)) [][]int {
	var path = make([][]int, n)
	for i, _ := range path {
		path[i] = make([]int, n)
		for j := 0; j < n; j++ {
			path[i][j] = math.MaxInt32 // 初始化最大距离
		}
		path[i][i] = 0
	}
	// 初始化 path-1，两点直达距离
	initPath(path)

	var cache = make([][][]*int, n)
	for i := 0; i < n; i++ {
		cache[i] = make([][]*int, n)
		for j := 0; j < n; j++ {
			cache[i][j] = make([]*int, n)
		}
	}

	// dfs(k,i,j)表示 i-j 的最短路且路径中节点的编号不超过 k 的 min
	// dfs(k,i,j) = min(dfs(k-1,i,j), dfs(k-1,i,k)+dfs(k-1,k,j))
	var dfs func(k, i, j int) int
	dfs = func(k, i, j int) int {
		if k == -1 {
			return path[i][j]
		}

		ptr := &cache[i][j][k]
		if *ptr != nil {
			return **ptr
		}

		v := min(
			dfs(k-1, i, j),
			dfs(k-1, i, k)+dfs(k-1, k, j),
		)
		*ptr = &v
		return v
	}

	// 迭代更新以 [0, n-1] 为中转节点的最短距离
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if i == j {
				continue
			}
			path[i][j] = dfs(n-1, i, j)
		}
	}
	return path
}
