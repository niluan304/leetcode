package main

func maximumScoreAfterOperations(edges [][]int, values []int) int64 {
	type Value struct {
		sum  int
		cost int
	}

	var n = len(values)
	var cache = make([]*Value, n)

	var graph = make([][]int, n)
	for _, edge := range edges {
		a, b := edge[0], edge[1]
		graph[a] = append(graph[a], b)
		graph[b] = append(graph[b], a)
	}

	type Visit struct {
		i, j int
	}
	var visit = map[Visit]bool{}

	var dfs func(i int) (value Value)
	dfs = func(i int) (value Value) {
		ptr := &cache[i]
		if *ptr != nil {
			return **ptr
		}
		defer func() { *ptr = &value }()

		var sum, cost = values[i], 0
		for _, path := range graph[i] {
			key := Visit{i: min(path, i), j: max(path, i)}
			if visit[key] {
				continue
			}

			visit[key] = true

			v := dfs(path)
			sum += v.sum
			cost += v.cost
		}

		if cost == 0 {
			cost = values[i]
		}
		return Value{
			sum:  sum,
			cost: min(cost, values[i]),
		}
	}

	ans := dfs(0)
	return int64(ans.sum - ans.cost)
}
