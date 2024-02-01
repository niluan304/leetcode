package main

import (
	"testing"

	"github.com/niluan304/leetcode/tests"
)

func Test_merge_two_sorted_lists(t *testing.T) {
	targetCaseNum := 0 // -1

	fs := []interface{}{
		mergeTwoLists,
	}

	for _, f := range fs {
		err := tests.RunFunc(t, f, samples, targetCaseNum)
		if err != nil {
			t.Error(err)
		}
	}
}

var samples = `
[1,2,4]
[1,3,4]
[1,1,2,3,4,4]

[]
[]
[]

[]
[0]
[0]

`
