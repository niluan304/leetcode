package main

import (
	"testing"

	"github.com/niluan304/leetcode/tests"
)

func Test_maximum_difference_between_node_and_ancestor(t *testing.T) {
	targetCaseNum := 0 // -1

	fs := []interface{}{
		maxAncestorDiff,
	}

	for _, f := range fs {
		err := tests.RunFunc(t, f, samples, targetCaseNum)
		if err != nil {
			t.Error(err)
		}
	}
}

var samples = `
[8,3,10,1,6,null,14,null,null,4,7,13]
7

[1,null,2,null,0,3]
3

`
