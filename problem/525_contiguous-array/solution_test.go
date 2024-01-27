package main

import (
	"testing"

	"github.com/niluan304/leetcode/tests"
)

func Test_findMaxLength(t *testing.T) {
	targetCaseNum := 0 // -1

	fs := []func(nums []int) int{
		//bruteForce,
		findMaxLength,
		findMaxLength2,
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
[0,1]
2

[0,1,0]
2


`
