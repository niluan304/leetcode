package main

import (
	"testing"

	"github.com/niluan304/leetcode/tests"
)

func Test_minimumEffortPath(t *testing.T) {
	targetCaseNum := 0 // -1

	fs := []func(heights [][]int) int{
		minimumEffortPath,
		minimumEffortPath2,
		leetcode1,
	}

	for _, f := range fs {
		err := tests.RunFunc(t, f, samples, targetCaseNum)
		if err != nil {
			t.Error(err)
		}
	}
}

var samples = `
[[1,2,2],[3,8,2],[5,3,5]]
2

[[1,2,3],[3,8,4],[5,3,5]]
1

[[1,2,1,1,1],[1,2,1,2,1],[1,2,1,2,1],[1,2,1,2,1],[1,1,1,2,1]]
0

[[1,10,6,7,9,10,4,9]]
9

[[4,3,4,10,5,5,9,2],[10,8,2,10,9,7,5,6],[5,8,10,10,10,7,4,2],[5,1,3,1,1,3,1,9],[6,4,10,6,10,9,4,6]]
5

`
