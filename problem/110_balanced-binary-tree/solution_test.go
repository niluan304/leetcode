package main

import (
	"testing"

	"github.com/niluan304/leetcode/tests"
)

func Test_balanced_binary_tree(t *testing.T) {
	targetCaseNum := 0 // -1

	fs := []interface{}{
		isBalanced,
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
true

[1,2,2,3,3,null,null,4,4]
false

[]
true

`
