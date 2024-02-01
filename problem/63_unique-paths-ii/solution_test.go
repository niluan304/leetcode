package main

import (
	"testing"

	"github.com/niluan304/leetcode/tests"
)

func Test_unique_paths_ii(t *testing.T) {
	targetCaseNum := 0 // -1

	fs := []interface{}{
		uniquePathsWithObstacles,
	}

	for _, f := range fs {
		err := tests.RunFunc(t, f, samples, targetCaseNum)
		if err != nil {
			t.Error(err)
		}
	}
}

var samples = `
[[0,0,0],[0,1,0],[0,0,0]]
2

[[0,1],[0,0]]
1

`
