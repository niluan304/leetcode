package main

import (
	"testing"

	"github.com/niluan304/leetcode/tests"
)

func Test_lengthOfLongestSubsequence(t *testing.T) {
	targetCaseNum := 0 // -1

	fs := []func(nums []int, target int) int{
		lengthOfLongestSubsequence,
		lengthOfLongestSubsequence2,
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
[1,2,3,4,5]
9
3

[4,1,3,2,1,5]
7
4

[1,1,5,4,5]
3
-1


`
