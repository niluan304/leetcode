package main

import (
	"testing"

	"github.com/niluan304/leetcode/tests"
)

func Test_find_minimum_in_rotated_sorted_array_ii(t *testing.T) {
	targetCaseNum := 0 // -1

	fs := []interface{}{
		findMin,
	}

	for _, f := range fs {
		err := tests.RunFunc(t, f, samples, targetCaseNum)
		if err != nil {
			t.Error(err)
		}
	}
}

var samples = `
[1,3,5]
1

[2,2,2,0,1]
0

`
