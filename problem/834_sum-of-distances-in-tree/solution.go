package main

// 换根 dp
// - 时间复杂度：O(n)。
// - 空间复杂度：O(n)。
func sumOfDistancesInTree(n int, edges [][]int) []int {
	graph := make([][]int, n)
	for _, edge := range edges {
		x, y := edge[0], edge[1]
		graph[x] = append(graph[x], y)
		graph[y] = append(graph[y], x)
	}

	size := make([]int, n)
	ans := make([]int, n)
	var dfs func(i, fa int, depth int) int
	dfs = func(i, fa int, depth int) int {
		ans[0] += depth
		for _, j := range graph[i] {
			if j == fa {
				continue
			}
			size[i] += dfs(j, i, depth+1)
		}
		size[i]++
		return size[i]
	}
	dfs(0, -1, 0)

	var reroot func(i int, fa int)
	reroot = func(i, fa int) {
		for _, j := range graph[i] { // 遍历 i 的邻居 j
			if j == fa {
				continue
			}
			ans[j] = ans[i] + n - 2*size[j]
			reroot(j, i) // i 是 j 的父节点
		}
	}
	reroot(0, -1) // 0 没有父节点
	return ans
}
