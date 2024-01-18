package main

import (
	"testing"

	"github.com/niluan304/leetcode/tests"
)

func Test_minimumRemoval(t *testing.T) {
	targetCaseNum := 0 // -1

	fs := []func(beans []int) int64{
		minimumRemoval,
		minimumRemoval2,
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
[4,1,6,5]
4

[2,10,3,2]
7


`
