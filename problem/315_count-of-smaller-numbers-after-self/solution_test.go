package main

import (
	"testing"

	"github.com/niluan304/leetcode/tests"
)

func Test_countSmaller(t *testing.T) {
	targetCaseNum := 0 // -1

	fs := []func(nums []int) []int{
		// bruteForce,
		countSmaller,
		countSmaller2,
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
[5,2,6,1]
[2,1,1,0]

[-1]
[0]

[-1,-1]
[0,0]


`
