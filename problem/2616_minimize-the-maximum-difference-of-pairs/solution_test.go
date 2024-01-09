package main

import (
	"testing"

	"github.com/niluan304/leetcode/tests"
)

func Test_minimizeMax(t *testing.T) {
	targetCaseNum := 0 // -1

	fs := []func(nums []int, p int) int{
		minimizeMax,
		// minimizeMax2,
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
[10,1,2,7,1,3]
2
1

[4,2,1,2]
1
0


`
