package main

import (
	"testing"

	"github.com/niluan304/leetcode/tests"
)

func Test_continuousSubarrays(t *testing.T) {
	targetCaseNum := 0 // -1

	fs := []func(nums []int) int64{
		continuousSubarrays,
		continuousSubarrays2,
		continuousSubarrays3,
		continuousSubarrays4,
		leetcode2,
		endlessCheng,
	}

	for _, f := range fs {
		err := tests.RunFunc(t, f, samples, targetCaseNum)
		if err != nil {
			t.Error(err)
		}
	}
}

var samples = `
[5,4,2,4]
8

[1,2,3]
6

[3,2,1]
6

[1,2,3,2,3,2,1,2,3,3,1,4,2,1]
71

[1,1,1,1,1,1,1,1,1,1,1]
66

`
