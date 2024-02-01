package main

import (
	"testing"

	"github.com/niluan304/leetcode/tests"
)

func Test_minimum_number_of_removals_to_make_mountain_array(t *testing.T) {
	targetCaseNum := 0 // -1

	fs := []interface{}{
		minimumMountainRemovals,
	}

	for _, f := range fs {
		err := tests.RunFunc(t, f, samples, targetCaseNum)
		if err != nil {
			t.Error(err)
		}
	}
}

var samples = `
[1,3,1]
0

[2,1,1,5,6,2,3,1]
3

`
