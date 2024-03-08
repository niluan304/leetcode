package main

import (
	"testing"

	"github.com/niluan304/leetcode/tests"
)

func Test_numberOfPairs(t *testing.T) {
	targetCaseNum := 0 // -1

	fs := []func(nums1 []int, nums2 []int, diff int) int64{
		// bruteForce,
		numberOfPairs,
		// numberOfPairs2,
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
[3,2,5]
[2,2,1]
1
3

[3,-1]
[-2,2]
-1
0


`
