package main

// #### 方法一：贪心 + 深度优先搜索
//
// **思路与算法**
//
// 题目等价于给出了一棵以节点 $0$ 为根结点的树，并且初始树上的每一个节点上都有一个人，现在所有人都需要通过「车子」向结点 $0$ 移动。
//
// 对于某一个节点 $x$，$x \neq 0$，其父节点为 $y$。因为以节点 $x$ 为根结点的子树上的人都需要通过边 $x \rightarrow y$ 向节点 $0$ 移动，所以为了使这条边上的「车子」利用率最高，我们贪心的让 $x$ 的全部子节点上的人到了节点 $x$ 后再一起坐车向上移动，我们不妨设以节点 $x$ 为根节点的子树大小为 $\text{cnt}_x$，那么我
// 们至少需要「车子」的数量为 $\lceil \frac{\textit{cnt}_x}{\textit{seats}} \rceil$，其中 $\textit{seats}$ 为一辆车的给定座位数。
//
// 那么我们可以通过从根结点 $0$ 往下进行「深度优先搜索」，每一条边上「车子」的数目即为该条边上汽油的开销，统计全部边上汽油的开销即为最终答案。
//
// **代码**
// *复杂度分析**
//
// - 时间复杂度：$O(n)$，其中 $n$ 为数组 $\textit{roads}$ 的长度。
// - 空间复杂度：$O(n)$，其中 $n$ 为数组 $\textit{roads}$ 的长度，主要为递归所需要的空间开销。
func leetcode(roads [][]int, seats int) int64 {
	graph := make([][]int, len(roads)+1)
	for _, e := range roads {
		a, b := e[0], e[1]
		graph[a] = append(graph[a], b)
		graph[b] = append(graph[b], a)
	}
	var res = 0
	var dfs func(i, fa int) int
	dfs = func(i, fa int) int {
		people := 1
		for _, j := range graph[i] {
			if j == fa {
				continue
			}
			p := dfs(j, i)
			res += (p + seats - 1) / seats
			people += p
		}
		return people
	}
	dfs(0, -1)
	return int64(res)
}
