package main

import (
	"testing"

	"github.com/niluan304/leetcode/tests"
)

func Test_lowest_common_ancestor_of_a_binary_tree(t *testing.T) {
	targetCaseNum := 0 // -1

	fs := []interface{}{
		lowestCommonAncestor,
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
5
1
3

[3,5,1,6,2,0,8,null,null,7,4]
5
4
5

[1,2]
1
2
1

`
