package main

func _max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func _min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func _abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func countPairs(n int, edges [][]int) int64 {
	var graph = make(map[int][]int, n) //邻接表
	for _, edge := range edges {
		e0, e1 := edge[0], edge[1]
		graph[e0] = append(graph[e0], e1)
		graph[e1] = append(graph[e1], e0)
	}

	var ans, m = 0, n
	for i := 0; i < n; i++ {
		if _, ok := graph[i]; !ok {
			continue
		}

		// BFS
		var queen = []int{i}
		circle := map[int]struct{}{}
		for len(queen) > 0 {
			head := queen[0]
			queen = queen[1:]

			for _, e1 := range graph[head] {
				queen = append(queen, e1)
			}
			circle[head] = struct{}{}
			delete(graph, head)
		}

		c := len(circle)
		ans += c * (m - c)
		m -= c
	}
	return int64(ans + m*(m-1)/2)
}
