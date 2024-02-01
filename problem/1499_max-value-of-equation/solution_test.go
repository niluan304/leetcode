package main

import (
	"testing"

	"github.com/niluan304/leetcode/tests"
)

func Test_main(t *testing.T) {
	targetCaseNum := 0 // -1

	fs := []func(points [][]int, k int) int{
		findMaxValueOfEquation,
		findMaxValueOfEquation2,
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
[[1,3],[2,0],[5,10],[6,-10]]
1
4

[[0,0],[3,0],[9,2]]
3
3

`
