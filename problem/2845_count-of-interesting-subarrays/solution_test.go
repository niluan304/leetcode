package main

import (
	"testing"

	"github.com/niluan304/leetcode/tests"
)

func Test_countInterestingSubarrays(t *testing.T) {
	targetCaseNum := 0 // -1

	fs := []func(nums []int, modulo int, k int) int64{
		bruteForce,
		countInterestingSubarrays,
		countInterestingSubarrays2,
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
[3,2,4]
2
1
3

[3,1,9,6]
3
0
2

`
