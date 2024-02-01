package main

import (
	"testing"

	"github.com/niluan304/leetcode/tests"
)

func Test_binary_tree_maximum_path_sum(t *testing.T) {
	targetCaseNum := 0 // -1

	fs := []interface{}{
		maxPathSum,
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
6

[-10,9,20,null,null,15,7]
42

`
