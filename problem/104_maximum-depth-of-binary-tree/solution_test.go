package main

import (
	"testing"

	"github.com/niluan304/leetcode/tests"
)

func Test_maximum_depth_of_binary_tree(t *testing.T) {
	targetCaseNum := 0 // -1

	fs := []interface{}{
		maxDepth,
	}

	for _, f := range fs {
		err := tests.RunFunc(t, f, samples, targetCaseNum)
		if err != nil {
			t.Error(err)
		}
	}
}

var samples = `
[3,9,20,null,null,15,7]
3

[1,null,2]
2

`
