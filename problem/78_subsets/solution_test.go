package main

import (
	"testing"

	"github.com/niluan304/leetcode/tests"
)

func Test_subsets(t *testing.T) {
	targetCaseNum := 0 // -1

	fs := []interface{}{
		subsets,
	}

	for _, f := range fs {
		err := tests.RunFunc(t, f, samples, targetCaseNum)
		if err != nil {
			t.Error(err)
		}
	}
}

var samples = `
[1,2,3]
[[],[1],[2],[1,2],[3],[1,3],[2,3],[1,2,3]]

[0]
[[],[0]]

`
