package main

import (
	"slices"
	"sort"
)

func maxOutput(n int, edges [][]int, price []int) int64 {
	var graph = make([][]int, n)
	for _, edge := range edges {
		a, b := edge[0], edge[1]
		graph[a] = append(graph[a], b)
		graph[b] = append(graph[b], a)
	}

	var ans = 0
	var dfs func(i, father int) (int, int)
	dfs = func(i, father int) (int, int) {
		p := price[i]
		sum1, sum2 := p, 0
		for _, j := range graph[i] {
			if j == father {
				continue
			}

			s1, s2 := dfs(j, i) // 分类讨论
			ans = max(ans, sum1+s2, sum2+s1)
			sum1 = max(sum1, s1+p)
			sum2 = max(sum2, s2+p)
		}
		return sum1, sum2
	}

	dfs(0, -1)
	return int64(ans)
}

// Deprecated: Fail pass
func maxOutput2(n int, edges [][]int, price []int) int64 {
	switch n {
	case 1:
		return 0
	case 2:
		return int64(slices.Max(price))
	}

	var graph = make([][]int, n)
	for _, edge := range edges {
		a, b := edge[0], edge[1]
		graph[a] = append(graph[a], b)
		graph[b] = append(graph[b], a)
	}

	var idx = 0
	for i, _ := range graph {
		if len(graph[i]) > len(graph[idx]) {
			idx = i
		}
	}

	var values [][2]int
	var visit = make([]bool, n)
	var dfs func(i int) (sum int, cost int)
	dfs = func(i int) (sum int, cost int) {
		visit[i] = true

		sum, cost = 0, price[i]
		for _, j := range graph[i] {
			if visit[j] {
				continue
			}

			sum2, cost2 := dfs(j)
			if sum2 > sum {
				sum = sum2
				cost = cost2
			}
			if i == idx {
				values = append(values, [2]int{sum2, cost2})
			}
		}
		return sum + price[i], cost
	}

	dfs(idx)

	sort.Slice(values, func(i, j int) bool {
		return values[i][0] < values[j][0]
	})

	var ans, m = 0, len(values)
	var mm = values[m-1]
	for _, value := range values[:m-1] {
		ans = max(
			ans,
			mm[0]+price[idx]+value[0]-
				min(value[1], mm[1]),
		)
	}
	return int64(ans)
}
