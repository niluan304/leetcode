package main

import (
	"testing"

	"github.com/niluan304/leetcode/tests"
)

func Test_reorder_list(t *testing.T) {
	targetCaseNum := 0 // -1

	fs := []interface{}{
		reorderList,
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
[1,4,2,3]

[1,2,3,4,5]
[1,5,2,4,3]

`
