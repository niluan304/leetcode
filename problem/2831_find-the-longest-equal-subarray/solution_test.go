package main

import (
	"testing"

	"github.com/niluan304/leetcode/tests"
)

func Test_longestEqualSubarray(t *testing.T) {
	targetCaseNum := 0 // -1

	fs := []func(nums []int, k int) int{
		bruteForce,
		longestEqualSubarray,
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
[1,3,2,3,1,3]
3
3

[1,1,2,2,1,1]
2
4

[6,4,7,6,4,8,6,6]
1
2


`
