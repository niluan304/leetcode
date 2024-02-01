package main

import (
	"testing"

	"github.com/niluan304/leetcode/tests"
)

func Test_longest_subarray_of_1s_after_deleting_one_element(t *testing.T) {
	targetCaseNum := 0 // -1

	fs := []interface{}{
		longestSubarray,
	}

	for _, f := range fs {
		err := tests.RunFunc(t, f, samples, targetCaseNum)
		if err != nil {
			t.Error(err)
		}
	}
}

var samples = `
[1,1,0,1]
3

[0,1,1,1,0,1,1,0,1]
5

[1,1,1]
2

`
