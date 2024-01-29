package main

import (
	"testing"

	"github.com/niluan304/leetcode/tests"
)

func Test_maxNonDecreasingLength(t *testing.T) {
	targetCaseNum := 0 // -1

	fs := []func(nums1 []int, nums2 []int) int{
		//bruteForce,
		maxNonDecreasingLength,
		//maxNonDecreasingLength2,
		//leetcode,
		//endlessCheng,
	}

	for _, f := range fs {
		err := tests.RunFunc(t, f, samples, targetCaseNum)
		if err != nil {
			t.Error(err)
		}
	}
}

var samples = `
[2,3,1]
[1,2,1]
2

[1,3,2,1]
[2,2,3,4]
4

[1,1]
[2,2]
2

[1]
[2]
1

`
