package main

import (
	"testing"

	"github.com/niluan304/leetcode/tests"
)

func Test_non_overlapping_intervals(t *testing.T) {
	targetCaseNum := 0 // -1

	fs := []interface{}{
		eraseOverlapIntervals,
	}

	for _, f := range fs {
		err := tests.RunFunc(t, f, samples, targetCaseNum)
		if err != nil {
			t.Error(err)
		}
	}
}

var samples = `
[[1,2],[2,3],[3,4],[1,3]]
1

[[1,2],[1,2],[1,2]]
2

[[1,2],[2,3]]
0

`
