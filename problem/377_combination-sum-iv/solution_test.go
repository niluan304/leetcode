package main

import (
	"testing"

	"github.com/niluan304/leetcode/tests"
)

func Test_combination_sum_iv(t *testing.T) {
	targetCaseNum := 0 // -1

	fs := []func(nums []int, target int) int{
		combinationSum4,
		combinationSum4II,
		combinationSum4III,
		combinationSum4IV,
	}

	for _, f := range fs {
		err := tests.RunFunc(t, f, samples, targetCaseNum)
		if err != nil {
			t.Error(err)
		}
	}
}

var samples = `
[1,2,3]
4
7

[9]
3
0

`
