package main

import (
	"testing"

	"github.com/niluan304/leetcode/tests"
)

func Test_lowest_common_ancestor_of_deepest_leaves(t *testing.T) {
	targetCaseNum := 0 // -1

	fs := []interface{}{
		lcaDeepestLeaves,
	}

	for _, f := range fs {
		err := tests.RunFunc(t, f, samples, targetCaseNum)
		if err != nil {
			t.Error(err)
		}
	}
}

var samples = `
[3,5,1,6,2,0,8,null,null,7,4]
[2,7,4]

[1]
[1]

[0,1,3,null,2]
[2]

`
