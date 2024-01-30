package main

import (
	"testing"

	"github.com/niluan304/leetcode/tests"
)

func Test_minimumSeconds(t *testing.T) {
	targetCaseNum := 0 // -1

	fs := []func(nums []int) int{
		//bruteForce,
		minimumSeconds,
		//minimumSeconds2,
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
[1,2,1,2]
1

[2,1,3,3,2]
2

[5,5,5,5]
0

[1,2,1,2,1,2,3,3,3]
2

[1,2,3,4,5,6]
3

[1,2,3,1,2,3,1,2,3,1,2,3,1,2,3]
1

[1,2,3,4,1,2,3,4,1,2,3,4,1,2,3,4]
2


`
