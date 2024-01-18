package main

import (
	"testing"

	"github.com/niluan304/leetcode/tests"
)

func Test_countSpecialNumbers(t *testing.T) {
	targetCaseNum := 0 // -1

	fs := []func(n int) int{
		countSpecialNumbers,
		// countSpecialNumbers2,
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
20
19

5
5

135
110


`
