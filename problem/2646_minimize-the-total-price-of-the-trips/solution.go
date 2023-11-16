package main

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
	for _, trip := range trips {
		start, end := trip[0], trip[1]
		tripPath(start, end, -1)
	}

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
