package main

import (
	"testing"

	"github.com/niluan304/leetcode/tests"
)

func Test_find_bottom_left_tree_value(t *testing.T) {
	targetCaseNum := 0 // -1

	fs := []interface{}{
		findBottomLeftValue,
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
1

[1,2,3,4,null,5,6,null,null,7]
7

`
