package main

import (
	"testing"

	"github.com/niluan304/leetcode/tests"
)

func Test_number_of_longest_increasing_subsequence(t *testing.T) {
	targetCaseNum := 0 // -1

	fs := []interface{}{
		findNumberOfLIS,
	}

	for _, f := range fs {
		err := tests.RunFunc(t, f, samples, targetCaseNum)
		if err != nil {
			t.Error(err)
		}
	}
}

var samples = `
[1,3,5,4,7]
2

[2,2,2,2,2]
5

`
