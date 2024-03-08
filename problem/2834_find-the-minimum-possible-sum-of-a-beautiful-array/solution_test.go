package main

import (
	"testing"

	"github.com/niluan304/leetcode/tests"
)

func Test_minimumPossibleSum(t *testing.T) {
	targetCaseNum := 0 // -1

	fs := []func(n int, target int) int{
		// bruteForce,
		minimumPossibleSum,
		// minimumPossibleSum2,
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
2
3
4

3
3
8

1
1
1

39636
49035
156198582

`
