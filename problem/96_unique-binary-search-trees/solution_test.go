package main

import (
	"testing"

	"github.com/niluan304/leetcode/tests"
)

func Test_unique_binary_search_trees(t *testing.T) {
	targetCaseNum := 0 // -1

	fs := []interface{}{
		numTrees,
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
5

1
1

`
