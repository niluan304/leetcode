package main

import (
	"testing"

	"github.com/niluan304/leetcode/tests"
)

func Test_target_sum(t *testing.T) {
	targetCaseNum := 0 // -1

	fs := []func(nums []int, target int) int{
		findTargetSumWays,
		findTargetSumWays2,
		findTargetSumWays3,
		findTargetSumWays4,
	}

	for _, f := range fs {
		err := tests.RunFunc(t, f, samples, targetCaseNum)
		if err != nil {
			t.Error(err)
		}
	}
}

var samples = `
[1,1,1,1,1]
3
5

[1]
1
1



`
