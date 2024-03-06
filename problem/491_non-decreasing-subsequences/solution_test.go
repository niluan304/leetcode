package main

import (
	"testing"

	"github.com/niluan304/leetcode/tests"
)

func Test_non_decreasing_subsequences(t *testing.T) {
	targetCaseNum := 0 // -1

	fs := []func(nums []int) [][]int{
		findSubsequences,
		findSubsequences2,
	}

	for _, f := range fs {
		err := tests.RunFunc(t, f, samples, targetCaseNum)
		if err != nil {
			t.Error(err)
		}
	}
}

var samples = `
[4,6,7,7]
[[4,6],[4,6,7],[4,6,7,7],[4,7],[4,7,7],[6,7],[6,7,7],[7,7]]

[4,4,3,2,1]
[[4,4]]

`
