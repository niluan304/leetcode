package main

import (
	"testing"

	"github.com/niluan304/leetcode/tests"
)

func Test_discuss(t *testing.T) {
	targetCaseNum := 0 // -1

	fs := []func(nums []int) int{
		// bruteForce,
		discuss,
		// leetcode,
		// endlessCheng,
	}

	for _, f := range fs {
		err := tests.RunFunc(t, f, discussSamples, targetCaseNum)
		if err != nil {
			t.Error(err)
		}
	}
}

var discussSamples = `
[4,4,4,5,6]
1

[1,1,1,2]
0


`
