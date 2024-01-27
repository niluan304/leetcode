package main

import (
	"testing"

	"github.com/niluan304/leetcode/tests"
)

func Test_checkSubarraySum(t *testing.T) {
	targetCaseNum := 0 // -1

	fs := []func(nums []int, k int) bool{
		//bruteForce,
		checkSubarraySum,
		checkSubarraySum2,
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
[23,2,4,6,7]
6
true

[23,2,6,4,7]
6
true

[23,2,6,4,7]
13
false

[0,0,0]
56757
true


`
