package main

import (
	"testing"

	"github.com/niluan304/leetcode/tests"
)

func Test_binary_tree_level_order_traversal(t *testing.T) {
	targetCaseNum := 0 // -1

	fs := []interface{}{
		levelOrder,
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
[[3],[9,20],[15,7]]

[1]
[[1]]

[]
[]

`
