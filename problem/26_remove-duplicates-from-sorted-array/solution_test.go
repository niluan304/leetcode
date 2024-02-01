package main

import (
	"testing"

	"github.com/niluan304/leetcode/tests"
)

func Test_remove_duplicates_from_sorted_array(t *testing.T) {
	targetCaseNum := 0 // -1

	fs := []interface{}{
		removeDuplicates,
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
2

[0,0,1,1,1,2,2,3,3,4]
5

`
