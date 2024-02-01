package main

import (
	"testing"

	"github.com/niluan304/leetcode/tests"
)

func Test_main(t *testing.T) {
	targetCaseNum := 0 // -1

	fs := []func(roads [][]int, seats int) int64{
		minimumFuelCost,
		minimumFuelCost2,
		leetcode,
	}

	for _, f := range fs {
		err := tests.RunFunc(t, f, samples, targetCaseNum)
		if err != nil {
			t.Error(err)
		}
	}
}

var samples = `
[[0,1],[0,2],[0,3]]
5
3

[[3,1],[3,2],[1,0],[0,4],[0,5],[4,6]]
2
7

[]
1
0

[[3,1],[3,2],[1,0],[0,4],[0,5],[4,6],[4,7]]
2
9

[[3,1],[3,2],[1,0],[0,4],[0,5],[4,6],[4,7],[3,8]]
2
11

[[0,1],[0,2],[1,3],[1,4]]
5
4
`
