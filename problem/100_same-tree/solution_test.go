package main

import (
	"testing"

	"github.com/niluan304/leetcode/tests"
)

func Test_same_tree(t *testing.T) {
	targetCaseNum := 0 // -1

	fs := []interface{}{
		isSameTree,
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
[1,2,3]
true

[1,2]
[1,null,2]
false

[1,2,1]
[1,1,2]
false

`
