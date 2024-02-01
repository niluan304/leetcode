package main

import (
	"testing"

	"github.com/niluan304/leetcode/tests"
)

func Test_count_unreachable_pairs_of_nodes_in_an_undirected_graph(t *testing.T) {
	targetCaseNum := 0 // -1

	fs := []interface{}{
		countPairs,
	}

	for _, f := range fs {
		err := tests.RunFunc(t, f, samples, targetCaseNum)
		if err != nil {
			t.Error(err)
		}
	}
}

var samples = `
3
[[0,1],[0,2],[1,2]]
0

7
[[0,2],[0,5],[2,4],[1,6],[5,4]]
14

`
