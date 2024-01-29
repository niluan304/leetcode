package main

import (
	"testing"

	"github.com/niluan304/leetcode/tests"
)

func Test_maximumLength(t *testing.T) {
	targetCaseNum := 0 // -1

	fs := []func(nums []int) int{
		//bruteForce,
		maximumLength,
		//maximumLength2,
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
[5,4,1,2,2]
3

[1,3,2,4]
1

[1,1,1,1,1,1,1,1,1]
9

[2,2,2,2,4,4,16,16,256,256]
7

[48841,358801,28561,18974736,4356,221,358801,599,13,4356,66,48841,28561,815730721,13,815730721,18974736,66,169,599,169,221]
7

`
