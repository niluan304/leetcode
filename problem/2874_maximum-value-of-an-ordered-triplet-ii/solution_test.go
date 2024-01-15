package main

import (
	"testing"

	"github.com/niluan304/leetcode/tests"
)

func Test_maximumTripletValue(t *testing.T) {
	targetCaseNum := 0 // -1

	fs := []func(nums []int) int64{
		maximumTripletValue,
		maximumTripletValue2,
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
[12,6,1,2,7]
77

[1,10,3,4,19]
133

[1,2,3]
0


`
