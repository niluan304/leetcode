package main

import (
	"testing"

	"github.com/niluan304/leetcode/tests"
)

func Test_firstMissingPositive(t *testing.T) {
	targetCaseNum := 0 // -1

	fs := []func(nums []int) int{
		// bruteForce,
		tip1,
		firstMissingPositive,
		// firstMissingPositive2,
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
[1,2,0]
3

[3,4,-1,1]
2

[7,8,9,11,12]
1

[1,1]
2

[1,1,2,2,3]
4

[1,1,3,4]
2

[2147483647]
1

`
