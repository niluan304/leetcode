package main

import (
	"testing"

	"github.com/niluan304/leetcode/tests"
)

func Test_lowest_common_ancestor_of_a_binary_search_tree(t *testing.T) {
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
[6,2,8,0,4,7,9,null,null,3,5]
2
8
6

[6,2,8,0,4,7,9,null,null,3,5]
2
4
2

[2,1]
2
1
2

`
