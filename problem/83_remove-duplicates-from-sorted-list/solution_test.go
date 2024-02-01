package main

import (
	"testing"

	"github.com/niluan304/leetcode/tests"
)

func Test_remove_duplicates_from_sorted_list(t *testing.T) {
	targetCaseNum := 0 // -1

	fs := []interface{}{
		deleteDuplicates,
	}

	for _, f := range fs {
		err := tests.RunFunc(t, f, samples, targetCaseNum)
		if err != nil {
			t.Error(err)
		}
	}
}

var samples = `
[1,1,2]
[1,2]

[1,1,2,3,3]
[1,2,3]

`
