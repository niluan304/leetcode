package main

import (
	"testing"

	"github.com/niluan304/leetcode/tests"
)

func Test_reverse_nodes_in_k_group(t *testing.T) {
	targetCaseNum := 0 // -1

	fs := []interface{}{
		reverseKGroup,
	}

	for _, f := range fs {
		err := tests.RunFunc(t, f, samples, targetCaseNum)
		if err != nil {
			t.Error(err)
		}
	}
}

var samples = `
[1,2,3,4,5]
2
[2,1,4,3,5]

[1,2,3,4,5]
3
[3,2,1,4,5]

`
