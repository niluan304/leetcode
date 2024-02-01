package main

import (
	"testing"

	"github.com/niluan304/leetcode/tests"
)

func Test_minimum_size_subarray_sum(t *testing.T) {
	targetCaseNum := 0 // -1

	fs := []func(target int, nums []int) int{
		minSubArrayLen,
		minSubArrayLen2,
	}

	for _, f := range fs {
		err := tests.RunFunc(t, f, samples, targetCaseNum)
		if err != nil {
			t.Error(err)
		}
	}
}

var samples = `
7
[2,3,1,2,4,3]
2

4
[1,4,4]
1

11
[1,1,1,1,1,1,1,1]
0

`
