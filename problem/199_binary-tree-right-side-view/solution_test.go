package main

import (
	"testing"

	"github.com/niluan304/leetcode/tests"
)

func Test_binary_tree_right_side_view(t *testing.T) {
	targetCaseNum := 0 // -1

	fs := []interface{}{
		rightSideView,
	}

	for _, f := range fs {
		err := tests.RunFunc(t, f, samples, targetCaseNum)
		if err != nil {
			t.Error(err)
		}
	}
}

var samples = `
[1,2,3,null,5,null,4]
[1,3,4]

[1,null,3]
[1,3]

[]
[]

`
