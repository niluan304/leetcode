package main

import (
	"testing"

	"github.com/niluan304/leetcode/tests"
)

func Test_string_transformation(t *testing.T) {
	targetCaseNum := 0 // -1

	fs := []interface{}{
		minCost2,
	}

	for _, f := range fs {
		err := tests.RunFunc(t, f, samples, targetCaseNum)
		if err != nil {
			t.Error(err)
		}
	}
}

func minCost2(n int, workers [][]int) int {
	var graph = make([][]int, n+1)
	var costs = make([]int, n+1)
	for _, worker := range workers {
		i := worker[0]
		v := worker[1]
		graph[i] = worker[3:]
		costs[i] = v
	}
	return minCost(1, graph, costs)
}

var samples = `
6
[[1,30,3,2,3,4],[2,16,2,5,6],[3,5,0],[4,4,0],[5,11,0],[6,5,0]]
25

6
[[1,30,3,2,3,4],[2,16,2,5,6],[3,15,0],[4,24,0],[5,11,0],[6,5,0]]
46

`
