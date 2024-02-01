package main

import (
	"testing"

	"github.com/niluan304/leetcode/tests"
)

func Test_symmetric_tree(t *testing.T) {
	targetCaseNum := 0 // -1

	fs := []interface{}{
		isSymmetric,
	}

	for _, f := range fs {
		err := tests.RunFunc(t, f, samples, targetCaseNum)
		if err != nil {
			t.Error(err)
		}
	}
}

var samples = `
[1,2,2,3,4,4,3]
true

[1,2,2,null,3,null,3]
false

`
