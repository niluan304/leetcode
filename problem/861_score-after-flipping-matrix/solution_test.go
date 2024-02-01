package main

import (
	"testing"

	"github.com/niluan304/leetcode/tests"
)

func Test_score_after_flipping_matrix(t *testing.T) {
	targetCaseNum := 0 // -1

	fs := []interface{}{
		matrixScore,
	}

	for _, f := range fs {
		err := tests.RunFunc(t, f, samples, targetCaseNum)
		if err != nil {
			t.Error(err)
		}
	}
}

var samples = `
[[0,0,1,1],[1,0,1,0],[1,1,0,0]]
39

[[0]]
1

`
