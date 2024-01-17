package main

import (
	"testing"

	"github.com/niluan304/leetcode/tests"
)

func Test_maximumNumberOfStringPairs(t *testing.T) {
	targetCaseNum := 0 // -1

	fs := []func(words []string) int{
		maximumNumberOfStringPairs,
		// maximumNumberOfStringPairs2,
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
["cd","ac","dc","ca","zz"]
2

["ab","ba","cc"]
1

["aa","ab"]
0


`
