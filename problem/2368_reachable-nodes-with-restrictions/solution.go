package main

// 树形 dfs
//
// - 时间复杂度：O(n)。
// - 空间复杂度：O(n)。
func reachableNodes(n int, edges [][]int, restricted []int) int {
	graph := make([][]int, n)
	for _, edge := range edges {
		x, y := edge[0], edge[1]
		graph[x] = append(graph[x], y)
		graph[y] = append(graph[y], x)
	}

	forbidden := make(map[int]bool, len(restricted))
	for _, i := range restricted {
		forbidden[i] = true
	}

	ans := 0
	var dfs func(i, fa int)
	dfs = func(i, fa int) {
		ans++
		for _, j := range graph[i] {
			if j == fa || forbidden[j] {
				continue
			}
			dfs(j, i)
		}
	}
	dfs(0, -1)
	return ans
}

// 树形 dfs
//
// reachableNodes 的优化方法
//
// - 时间复杂度：O(n)。
// - 空间复杂度：O(n)。
func reachableNodes2(n int, edges [][]int, restricted []int) int {
	forbidden := make(map[int]bool, len(restricted))
	for _, i := range restricted {
		forbidden[i] = true
	}

	graph := make([][]int, n)
	for _, edge := range edges {
		x, y := edge[0], edge[1]
		if forbidden[x] || forbidden[y] {
			continue
		}
		graph[x] = append(graph[x], y)
		graph[y] = append(graph[y], x)
	}

	ans := 0
	var dfs func(i, fa int)
	dfs = func(i, fa int) {
		ans++
		for _, j := range graph[i] {
			if j == fa {
				continue
			}
			dfs(j, i)
		}
	}
	dfs(0, -1)
	return ans
}
