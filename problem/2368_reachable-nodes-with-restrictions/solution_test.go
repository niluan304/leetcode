package main

import (
	"testing"

	"github.com/niluan304/leetcode/tests"
)

func Test_reachableNodes(t *testing.T) {
	targetCaseNum := 0 // -1

	fs := []func(n int, edges [][]int, restricted []int) int{
		// bruteForce,
		reachableNodes,
		reachableNodes2,
		// leetcode,
		// endlessCheng,
	}

	for _, f := range fs {
		err := tests.RunFunc(t, f, samples, targetCaseNum)
		if err != nil {
			t.Error(err)
		}
	}
}

var samples = `
7
[[0,1],[1,2],[3,1],[4,0],[0,5],[5,6]]
[4,5]
4

7
[[0,1],[0,2],[0,5],[0,4],[3,2],[6,5]]
[4,2,1]
3


`
