package main

import (
	"testing"

	"github.com/niluan304/leetcode/tests"
)

func Test_main(t *testing.T) {
	targetCaseNum := 0 // -1

	fs := []func(matrix [][]int) int{
		countSquares,
	}

	for _, f := range fs {
		err := tests.RunFunc(t, f, samples, targetCaseNum)
		if err != nil {
			t.Error(err)
		}
	}
}

var samples = `
[[0,1,1,1],[1,1,1,1],[0,1,1,1]]
15

[[1,0,1],[1,1,0],[1,1,0]]
7


`
