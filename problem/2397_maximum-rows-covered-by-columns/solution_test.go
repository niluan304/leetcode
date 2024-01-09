package main

import (
	"testing"

	"github.com/niluan304/leetcode/tests"
)

func Test_main(t *testing.T) {
	targetCaseNum := 0 // -1

	fs := []func(matrix [][]int, numSelect int) int{
		maximumRows,
		maximumRows2,
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
[[0,0,0],[1,0,1],[0,1,1],[0,0,1]]
2
3

[[1],[0]]
1
2

[[1,1,0,0,0,0,0,0],[1,1,0,0,0,0,0,0],[1,1,0,0,0,0,0,0],[1,1,0,0,0,0,0,0],[0,0,0,0,1,0,0,0]]
2
4


[[1,1,1,0,1,0,0,0,1,0,0,0,0],[0,1,1,1,1,0,0,0,1,0,0,0,0],[1,0,0,1,1,0,0,1,1,0,0,0,0],[1,1,0,1,1,0,0,0,0,0,0,0,0],[1,1,1,1,1,0,0,1,1,0,0,0,0],[1,1,1,1,1,0,0,1,1,0,0,0,0]]
6
3

`
