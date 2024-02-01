package main

import (
	"testing"

	"github.com/niluan304/leetcode/tests"
)

func Test_merge_intervals(t *testing.T) {
	targetCaseNum := 0 // -1

	fs := []interface{}{
		merge,
	}

	for _, f := range fs {
		err := tests.RunFunc(t, f, samples, targetCaseNum)
		if err != nil {
			t.Error(err)
		}
	}
}

var samples = `
[[1,3],[2,6],[8,10],[15,18]]
[[1,6],[8,10],[15,18]]

[[1,4],[4,5]]
[[1,5]]

`
