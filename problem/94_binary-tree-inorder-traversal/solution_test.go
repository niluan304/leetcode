package main

import (
	"testing"

	"github.com/niluan304/leetcode/tests"
)

func Test_binary_tree_inorder_traversal(t *testing.T) {
	targetCaseNum := 0 // -1

	fs := []interface{}{
		inorderTraversal,
	}

	for _, f := range fs {
		err := tests.RunFunc(t, f, samples, targetCaseNum)
		if err != nil {
			t.Error(err)
		}
	}
}

var samples = `
[1,null,2,3]
[1,3,2]

[]
[]

[1]
[1]

`
