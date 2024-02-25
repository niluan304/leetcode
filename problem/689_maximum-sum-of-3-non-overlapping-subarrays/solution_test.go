package main

import (
	"testing"

	"github.com/niluan304/leetcode/tests"
)

func Test_maxSumOfThreeSubarrays(t *testing.T) {
	targetCaseNum := 0 // -1

	fs := []func(nums []int, k int) []int{
		// bruteForce,
		maxSumOfThreeSubarrays,
		maxSumOfThreeSubarrays2,
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
[1,2,1,2,6,7,5,1]
2
[0,3,5]

[1,2,1,2,1,2,1,2,1]
2
[0,2,4]

[1,2,1,2,1,2]
2
[0,2,4]

[4,3,2,1]
1
[0,1,2]


`
