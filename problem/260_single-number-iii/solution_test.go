package main

import (
	"testing"

	"github.com/niluan304/leetcode/tests"
)

func Test_single_number_iii(t *testing.T) {
	targetCaseNum := 0 // -1

	fs := []func(nums []int) []int{
		singleNumber,
	}

	for _, f := range fs {
		err := tests.RunFunc(t, f, samples, targetCaseNum)
		if err != nil {
			t.Error(err)
		}
	}
}

var samples = `
[1,2,1,3,2,5]
[3,5]

[-1,0]
[-1,0]

[0,1]
[1,0]

`
