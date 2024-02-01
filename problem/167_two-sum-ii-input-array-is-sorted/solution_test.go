package main

import (
	"testing"

	"github.com/niluan304/leetcode/tests"
)

func Test_two_sum_ii_input_array_is_sorted(t *testing.T) {
	targetCaseNum := 0 // -1

	fs := []interface{}{
		twoSum,
	}

	for _, f := range fs {
		err := tests.RunFunc(t, f, samples, targetCaseNum)
		if err != nil {
			t.Error(err)
		}
	}
}

var samples = `
[2,7,11,15]
9
[1,2]

[2,3,4]
6
[1,3]

[-1,0]
-1
[1,2]

`
