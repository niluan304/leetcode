package main

// dfs
// 转化为正常的树，并额外记录边的方向。
// 然后对这棵树进行 dfs，遍历过程中，如果某个父节点无法被它的子节点访问，那么就应该改变边的方向。
// 时间复杂度: O(n)
// 空间复杂度: O(n)
func minReorder(n int, connections [][]int) int {
	var graph = make([][]int, n)
	var pass = make(map[[2]int]bool, n)
	for _, connection := range connections {
		a, b := connection[0], connection[1]
		graph[a] = append(graph[a], b)
		graph[b] = append(graph[b], a)
		pass[[2]int{a, b}] = true
	}

	var ans = 0
	var dfs func(i int, fa int)
	dfs = func(i int, fa int) {
		for _, j := range graph[i] {
			if j == fa {
				continue
			}
			if pass[[2]int{i, j}] {
				ans++
			}
			dfs(j, i)
		}
	}
	dfs(0, -1)
	return ans
}
