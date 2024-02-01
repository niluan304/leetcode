package main

import (
	"testing"

	"github.com/niluan304/leetcode/tests"
)

func Test_subsets_ii(t *testing.T) {
	targetCaseNum := 0 // -1

	fs := []func(nums []int) [][]int{
		subsetsWithDup,
		subsetsWithDup2,
		subsetsWithDup3,
	}

	for _, f := range fs {
		err := tests.RunFunc(t, f, samples, targetCaseNum)
		if err != nil {
			t.Error(err)
		}
	}
}

var samples = `
[1,2,2]
[[],[1],[1,2],[1,2,2],[2],[2,2]]

[0]
[[],[0]]

`
