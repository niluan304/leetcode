package main

import (
	"testing"

	"github.com/niluan304/leetcode/tests"
)

func Test_maximumScoreAfterOperations(t *testing.T) {
	targetCaseNum := 0 // -1

	fs := []func(edges [][]int, values []int) int64{
		//bruteForce,
		maximumScoreAfterOperations,
		//maximumScoreAfterOperations2,
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
[[0,1],[0,2],[0,3],[2,4],[4,5]]
[5,2,5,2,1,1]
11

[[0,1],[0,2],[1,3],[1,4],[2,5],[2,6]]
[20,10,9,7,4,3,5]
40

[[0,1]]
[2,1]
2

`
