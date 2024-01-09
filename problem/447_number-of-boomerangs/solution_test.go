package main

import (
	"testing"

	"github.com/niluan304/leetcode/tests"
)

func Test_main(t *testing.T) {
	targetCaseNum := 0 // -1

	fs := []func(points [][]int) int{
		numberOfBoomerangs,
		numberOfBoomerangs2,
	}

	for _, f := range fs {
		err := tests.RunFunc(t, f, samples, targetCaseNum)
		if err != nil {
			t.Error(err)
		}
	}
}

var samples = `
[[0,0],[1,0],[2,0]]
2

[[1,1],[2,2],[3,3]]
2

[[1,1]]
0


`
