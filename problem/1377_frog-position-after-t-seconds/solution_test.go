package main

import (
	"testing"

	"github.com/niluan304/leetcode/tests"
)

func Test_frog_position_after_t_seconds(t *testing.T) {
	targetCaseNum := 0 // -1

	fs := []func(n int, edges [][]int, t int, target int) float64{
		frogPosition,
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
[[1,2],[1,3],[1,7],[2,4],[2,6],[3,5]]
2
4
0.16666666666666666

7
[[1,2],[1,3],[1,7],[2,4],[2,6],[3,5]]
1
7
0.3333333333333333

`
