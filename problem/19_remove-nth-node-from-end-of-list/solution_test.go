package main

import (
	"testing"

	"github.com/niluan304/leetcode/tests"
)

func Test_remove_nth_node_from_end_of_list(t *testing.T) {
	targetCaseNum := 0 // -1

	fs := []interface{}{
		removeNthFromEnd,
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
[1,2,3,5]

[1]
1
[]

[1,2]
1
[1]

`
