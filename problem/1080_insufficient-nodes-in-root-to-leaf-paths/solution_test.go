package main

import (
	"testing"

	"github.com/niluan304/leetcode/tests"
)

func Test_insufficient_nodes_in_root_to_leaf_paths(t *testing.T) {
	targetCaseNum := 0 // -1

	fs := []interface{}{
		sufficientSubset,
	}

	for _, f := range fs {
		err := tests.RunFunc(t, f, samples, targetCaseNum)
		if err != nil {
			t.Error(err)
		}
	}
}

var samples = `
[1,2,3,4,-99,-99,7,8,9,-99,-99,12,13,-99,14]
1
[1,2,3,4,null,null,7,8,9,null,14]

[5,4,8,11,null,17,4,7,1,null,null,5,3]
22
[5,4,8,11,null,17,4,7,null,null,null,5]

[1,2,-3,-5,null,4,null]
-1
[1,null,-3,4]

`
