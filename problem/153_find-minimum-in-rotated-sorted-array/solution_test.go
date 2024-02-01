package main

import (
	"testing"

	"github.com/niluan304/leetcode/tests"
)

func Test_find_minimum_in_rotated_sorted_array(t *testing.T) {
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
[3,4,5,1,2]
1

[4,5,6,7,0,1,2]
0

[11,13,15,17]
11

`
