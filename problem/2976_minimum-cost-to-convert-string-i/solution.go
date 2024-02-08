package main

import (
	"math"

	. "github.com/niluan304/leetcode/copypasta/graph"
)

func minimumCost(source string, target string, original []byte, changed []byte, cost []int) int64 {
	distance := Floyd(26, func(distance [][]int) {
		for x, _ := range original {
			i, j := original[x]-'a', changed[x]-'a'
			distance[i][j] = min(distance[i][j], cost[x]) // i -> j 的最低代价
		}
	})

	var ans = 0
	for x, _ := range source {
		i, j := source[x]-'a', target[x]-'a'
		dis := distance[i][j]
		if dis >= math.MaxInt/2 {
			return -1
		}
		ans += dis
	}
	return int64(ans)
}
