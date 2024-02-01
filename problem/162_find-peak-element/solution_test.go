package main

import (
	"testing"

	"github.com/niluan304/leetcode/tests"
)

func Test_find_peak_element(t *testing.T) {
	targetCaseNum := 0 // -1

	fs := []func(nums []int) int{
		findPeakElement,
		findPeakElement2,
	}

	for _, f := range fs {
		err := tests.RunFunc(t, f, samples, targetCaseNum)
		if err != nil {
			t.Error(err)
		}
	}
}

var samples = `
[1,2,3,1]
2

[1,2,1,3,5,6,4]
5

`
