package main

import (
	"testing"

	"github.com/niluan304/leetcode/tests"
)

func Test_reversePairs(t *testing.T) {
	targetCaseNum := 0 // -1

	fs := []func(record []int) int{
		// bruteForce,
		reversePairs,
		reversePairs2,
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
[9, 7, 5, 4, 6]
8

`
