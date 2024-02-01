package main

import (
	"testing"

	"github.com/niluan304/leetcode/tests"
)

func Test_combination_sum_ii(t *testing.T) {
	targetCaseNum := 0 // -1

	fs := []interface{}{
		combinationSum2,
	}

	for _, f := range fs {
		err := tests.RunFunc(t, f, samples, targetCaseNum)
		if err != nil {
			t.Error(err)
		}
	}
}

var samples = `
[10,1,2,7,6,1,5]
8
[[1,1,6],[1,2,5],[1,7],[2,6]]

[2,5,2,1,2]
5
[[1,2,2],[5]]

`
