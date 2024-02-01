package main

import (
	"testing"

	"github.com/niluan304/leetcode/tests"
)

func Test_main(t *testing.T) {
	targetCaseNum := 0 // -1

	fs := []func(n int, connections [][]int) int{
		minReorder,
	}

	for _, f := range fs {
		err := tests.RunFunc(t, f, samples, targetCaseNum)
		if err != nil {
			t.Error(err)
		}
	}
}

var samples = `
6
[[0,1],[1,3],[2,3],[4,0],[4,5]]
3

5
[[1,0],[1,2],[3,2],[3,4]]
2

3
[[1,0],[2,0]]
0

`
