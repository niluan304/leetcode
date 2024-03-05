package main

import (
	"testing"

	"github.com/niluan304/leetcode/tests"
)

func Test_countPaths(t *testing.T) {
	targetCaseNum := 0 // -1

	fs := []func(n int, roads [][]int) int{
		// bruteForce,
		countPaths,
		countPaths2,
		countPaths3,
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
[[0,6,7],[0,1,2],[1,2,3],[1,3,3],[6,3,3],[3,5,1],[6,5,1],[2,5,1],[0,4,5],[4,6,2]]
4

2
[[1,0,10]]
1


`
