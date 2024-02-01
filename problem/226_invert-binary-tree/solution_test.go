package main

import (
	"testing"

	"github.com/niluan304/leetcode/tests"
)

func Test_invert_binary_tree(t *testing.T) {
	targetCaseNum := 0 // -1

	fs := []interface{}{
		invertTree,
	}

	for _, f := range fs {
		err := tests.RunFunc(t, f, samples, targetCaseNum)
		if err != nil {
			t.Error(err)
		}
	}
}

var samples = `
[4,2,7,1,3,6,9]
[4,7,2,9,6,3,1]

[2,1,3]
[2,3,1]

[]
[]

`
