package main

import "math"

func hardestWorker(n int, logs [][]int) int {
	var cost, id = math.MinInt, 0

	for i, log := range logs {
		x := log[1]
		if i > 0 {
			x -= logs[i-1][1]
		}
		if x > cost {
			cost = x
			id = log[0]
		}
		if x == cost {
			id = min(id, log[0])
		}
	}
	return id
}
