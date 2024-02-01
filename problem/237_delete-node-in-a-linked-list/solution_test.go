package main

import (
	"testing"

	"github.com/niluan304/leetcode/tests"
)

func Test_delete_node_in_a_linked_list(t *testing.T) {
	targetCaseNum := 0 // -1

	fs := []interface{}{
		deleteNode,
	}

	for _, f := range fs {
		err := tests.RunFunc(t, f, samples, targetCaseNum)
		if err != nil {
			t.Error(err)
		}
	}
}

var samples = `
[4,5,1,9]
5
[4,1,9]

[4,5,1,9]
1
[4,5,9]

`
