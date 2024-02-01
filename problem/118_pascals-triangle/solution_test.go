package main

import (
	"testing"

	"github.com/niluan304/leetcode/tests"
)

func Test_pascals_triangle(t *testing.T) {
	targetCaseNum := 0 // -1

	fs := []interface{}{
		generate,
	}

	for _, f := range fs {
		err := tests.RunFunc(t, f, samples, targetCaseNum)
		if err != nil {
			t.Error(err)
		}
	}
}

var samples = `
5
[[1],[1,1],[1,2,1],[1,3,3,1],[1,4,6,4,1]]

1
[[1]]

`
