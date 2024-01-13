package main

import (
	"testing"

	"github.com/niluan304/leetcode/tests"
)

func Test_minimumSum(t *testing.T) {
	targetCaseNum := 0 // -1

	fs := []func(nums []int) int{
		minimumSum,
		// minimumSum2,
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
[8,6,1,5,3]
9

[5,4,8,7,10,2]
13

[6,5,4,3,4,5]
-1


`
