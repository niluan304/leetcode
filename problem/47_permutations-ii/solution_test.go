package main

import (
	"testing"

	"github.com/niluan304/leetcode/tests"
)

func Test_permutations_ii(t *testing.T) {
	targetCaseNum := 0 // -1

	fs := []interface{}{
		permuteUnique,
	}

	for _, f := range fs {
		err := tests.RunFunc(t, f, samples, targetCaseNum)
		if err != nil {
			t.Error(err)
		}
	}
}

var samples = `
[1,1,2]
[[1,1,2],[1,2,1],[2,1,1]]

[1,2,3]
[[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]]

`
