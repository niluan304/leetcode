package main

import (
	"testing"

	"github.com/niluan304/leetcode/tests"
)

func Test_minimumMountainRemovals(t *testing.T) {
	targetCaseNum := 0 // -1

	fs := []func(nums []int) int{
		//bruteForce,
		minimumMountainRemovals,
		//minimumMountainRemovals2,
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
[1,3,1]
0

[2,1,1,5,6,2,3,1]
3

[9,8,1,7,6,5,4,3,2,1]
2

`
