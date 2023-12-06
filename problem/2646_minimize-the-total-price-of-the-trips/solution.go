package main

// 回溯 + dp
// 时间复杂度: O(mn)
// 空间复杂度: O(n)
// 对 trips 进行回溯，找到 trips[i] = [starti, endi] 经过的点，经过的点 weight[i] + 1
// 然后 对所有点进行 打家劫舍III 计算
func minimumTotalPrice(n int, edges [][]int, price []int, trips [][]int) int {
	var graph = make([][]int, n)
	for _, edge := range edges {
		x, y := edge[0], edge[1]
		graph[x] = append(graph[x], y)
		graph[y] = append(graph[y], x)
	}

	var weight = make([]int, n)
	var tripPath func(start, end, fa int) (find bool)
	tripPath = func(start, end, fa int) (find bool) {
		weight[start]++
		if start == end {
			return true
		}
		for _, i := range graph[start] {
			if i == fa {
				continue
			}
			if tripPath(i, end, start) {
				return true
			}
		}
		weight[start]--
		return false
	}

	// 遍历 trips 时间复杂度：O(mn)
	for _, trip := range trips {
		start, end := trip[0], trip[1]
		tripPath(start, end, -1)
	}

	// 树形 dp 时间复杂度：O(n)
	var dfs func(i int, fa int) (full, half int)
	dfs = func(i int, fa int) (full, half int) {
		full = weight[i] * price[i]
		half = full / 2
		for _, j := range graph[i] {
			if j == fa {
				continue
			}
			fullJ, halfJ := dfs(j, i)
			full += min(fullJ, halfJ)
			half += fullJ
		}
		return full, half
	}
	return min(dfs(0, -1))
}
