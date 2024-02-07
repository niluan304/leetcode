package main

import (
	"math"
	"sort"
)

func minCost(n int, cuts []int) int {
	sort.Ints(cuts)
	type Key struct{ i, j int }
	var cache = make(map[Key]int, n)

	var dfs func(i, j int) int
	dfs = func(i, j int) int {
		key := Key{i: i, j: j}
		if v, ok := cache[key]; ok {
			return v
		}

		x, y := sort.SearchInts(cuts, i+1), sort.SearchInts(cuts, j)
		if x == y {
			cache[key] = 0
			return 0
		}

		cost := math.MaxInt32
		for _, cut := range cuts[x:y] {
			v1 := dfs(i, cut)
			v2 := dfs(cut, j)
			cost = min(cost, j-i+v1+v2)
		}

		cache[key] = cost
		return cost
	}
	return dfs(0, n)
}

func minCost2(n int, cuts []int) int {
	cuts = append(cuts, 0, n)
	sort.Ints(cuts)

	m := len(cuts)
	var cache = make([][]*int, m)
	for i, _ := range cache {
		cache[i] = make([]*int, m)
	}
	for i := 0; i < m-1; i++ {
		cache[i][i+1] = new(int)
	}

	var dfs func(i, j int) int
	dfs = func(i, j int) int {
		ptr := &cache[i][j]
		if *ptr != nil {
			return **ptr
		}

		cost := math.MaxInt32
		for k := i + 1; k < j; k++ {
			cost = min(
				cost,
				cuts[j]-cuts[i]+dfs(i, k)+dfs(k, j),
			)
		}
		*ptr = &cost
		return cost
	}
	return dfs(0, m-1)
}
