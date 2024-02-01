package main

import (
	"testing"

	"github.com/niluan304/leetcode/tests"
)

func Test_combination_sum_iii(t *testing.T) {
	targetCaseNum := 0 // -1

	fs := []interface{}{
		combinationSum3,
	}

	for _, f := range fs {
		err := tests.RunFunc(t, f, samples, targetCaseNum)
		if err != nil {
			t.Error(err)
		}
	}
}

var samples = `
3
7
[[1,2,4]]

3
9
[[1,2,6],[1,3,5],[2,3,4]]

4
1
[]

`
