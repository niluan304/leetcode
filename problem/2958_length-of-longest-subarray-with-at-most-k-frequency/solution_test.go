package main

import (
	"testing"

	"github.com/niluan304/leetcode/tests"
)

func Test_maxSubarrayLength(t *testing.T) {
	targetCaseNum := 0 // -1

	fs := []func(nums []int, k int) int{
		maxSubarrayLength,
		// maxSubarrayLength2,
		// leetcode,
		// endlessCheng,
	}

	for _, f := range fs {
		err := tests.RunFunc(t, f, samples, targetCaseNum)
		if err != nil {
			t.Error(err)
		}
	}
}

var samples = `
[1,2,3,1,2,3,1,2]
2
6

[1,2,1,2,1,2,1,2]
1
2

[5,5,5,5,5,5,5]
4
4


`
