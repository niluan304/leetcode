package main

import (
	"testing"

	"github.com/niluan304/leetcode/tests"
)

func Test_combination_sum(t *testing.T) {
	targetCaseNum := 0 // -1

	fs := []interface{}{
		combinationSum,
	}

	for _, f := range fs {
		err := tests.RunFunc(t, f, samples, targetCaseNum)
		if err != nil {
			t.Error(err)
		}
	}
}

var samples = `
[2,3,6,7]
7
[[2,2,3],[7]]

[2,3,5]
8
[[2,2,2,2],[2,3,3],[3,5]]

[2]
1
[]

`
