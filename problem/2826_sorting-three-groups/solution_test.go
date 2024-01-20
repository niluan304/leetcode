package main

import (
	"testing"

	"github.com/niluan304/leetcode/tests"
)

func Test_minimumOperations(t *testing.T) {
	targetCaseNum := 0 // -1

	fs := []func(nums []int) int{
		minimumOperations,
		minimumOperations2,
		minimumOperations3,
		minimumOperations4,
		minimumOperations5,
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
[2,1,3,2,1]
3

[1,3,2,1,3,3]
2

[2,2,2,2,3,3]
0

[1]
0

`
