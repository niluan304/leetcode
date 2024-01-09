package main

import (
	"testing"

	"github.com/niluan304/leetcode/tests"
)

func Test_hardestWorker(t *testing.T) {
	targetCaseNum := 0 // -1

	fs := []func(n int, logs [][]int) int{
		hardestWorker,
		// hardestWorker2,
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
10
[[0,3],[2,5],[0,9],[1,15]]
1

26
[[1,1],[3,7],[2,12],[7,17]]
3

2
[[0,10],[1,20]]
0


`
