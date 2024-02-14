package main

import (
	"testing"

	"github.com/niluan304/leetcode/tests"
)

func Test_minPathSum(t *testing.T) {
	targetCaseNum := 0 // -1

	fs := []func(grid [][]int) int{
		//bruteForce,
		minPathSum,
		minPathSum2,
		//leetcode,
		//endlessCheng,
	}

	for _, f := range fs {
		err := tests.RunFunc(t, f, samples, targetCaseNum)
		if err != nil {
			t.Error(err)
		}
	}
}

var samples = `
[[1,3,1],[1,5,1],[4,2,1]]
7

[[1,2,3],[4,5,6]]
12


`
