package main

import (
	"testing"

	"github.com/niluan304/leetcode/tests"
)

func Test_validate_binary_search_tree(t *testing.T) {
	targetCaseNum := 0 // -1

	fs := []interface{}{
		isValidBST,
	}

	for _, f := range fs {
		err := tests.RunFunc(t, f, samples, targetCaseNum)
		if err != nil {
			t.Error(err)
		}
	}
}

var samples = `
[2,1,3]
true

[5,1,4,null,null,3,6]
false

`
