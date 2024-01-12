package main

import (
	"testing"

	"github.com/niluan304/leetcode/tests"
)

func Test_maximumCostSubstring(t *testing.T) {
	targetCaseNum := 0 // -1

	fs := []func(s string, chars string, vals []int) int{
		maximumCostSubstring,
		maximumCostSubstring2,
		maximumCostSubstring3,
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
"adaa"
"d"
[-1000]
2

"abc"
"abc"
[-1,-1,-1]
0


`
