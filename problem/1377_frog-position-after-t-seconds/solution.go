package main

func frogPosition(n int, edges [][]int, t int, target int) float64 {
	var graph = make([][]int, n+1)
	for _, edge := range edges {
		x, y := edge[0], edge[1]
		graph[x] = append(graph[x], y)
		graph[y] = append(graph[y], x)
	}
	var ans *float64
	var dfs func(i int, fa int, t int, p float64)

	dfs = func(i int, fa int, t int, p float64) {
		if ans != nil {
			return
		}
		l := len(graph[i])
		if fa != 0 {
			l--
		}

		if l == 0 && t > 0 {
			t = 0
		}
		if i == target {
			if t == 0 {
				ans = &p
			} else {
				ans = new(float64)
			}
			return
		}
		for _, j := range graph[i] {
			if j == fa {
				continue
			}
			dfs(j, i, t-1, p/float64(l))
		}
	}
	dfs(1, 0, t, 1)
	return *ans
}
