package main

import (
	"testing"

	"github.com/niluan304/leetcode/tests"
)

func Test_minimumVisitedCells(t *testing.T) {
	targetCaseNum := 0 // -1

	fs := []func(grid [][]int) int{
		bruteForce,
		minimumVisitedCells,
		minimumVisitedCells2,
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
[[3,4,2,1],[4,2,3,1],[2,1,0,0],[2,4,0,0]]
4

[[3,4,2,1],[4,2,1,1],[2,1,1,0],[3,4,1,0]]
3

[[2,1,0],[1,0,0]]
-1

[[7,0,2,4],[0,2,3,0]]
3


`
