package main

import (
	"testing"

	"github.com/niluan304/leetcode/tests"
)

func Test_swap_nodes_in_pairs(t *testing.T) {
	targetCaseNum := 0 // -1

	fs := []interface{}{
		swapPairs,
	}

	for _, f := range fs {
		err := tests.RunFunc(t, f, samples, targetCaseNum)
		if err != nil {
			t.Error(err)
		}
	}
}

var samples = `
[1,2,3,4]
[2,1,4,3]

[]
[]

[1]
[1]

`
