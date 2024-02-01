package main

import (
	"testing"

	"github.com/niluan304/leetcode/tests"
)

func Test_set_matrix_zeroes(t *testing.T) {
	targetCaseNum := 0 // -1

	fs := []interface{}{
		setZeroes,
	}

	for _, f := range fs {
		err := tests.RunFunc(t, f, samples, targetCaseNum)
		if err != nil {
			t.Error(err)
		}
	}
}

var samples = `
[[1,1,1],[1,0,1],[1,1,1]]
[[1,0,1],[0,0,0],[1,0,1]]

[[0,1,2,0],[3,4,5,2],[1,3,1,5]]
[[0,0,0,0],[0,4,5,0],[0,3,1,0]]

`
