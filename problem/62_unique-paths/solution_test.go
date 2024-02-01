package main

import (
	"testing"

	"github.com/niluan304/leetcode/tests"
)

func Test_unique_paths(t *testing.T) {
	targetCaseNum := 0 // -1

	fs := []interface{}{
		uniquePaths,
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
28

3
2
3

`
