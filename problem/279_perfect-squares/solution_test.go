package main

import (
	"testing"

	"github.com/niluan304/leetcode/tests"
)

func Test_numSquares(t *testing.T) {
	targetCaseNum := 0 // -1

	fs := []func(n int) int{
		//bruteForce,
		numSquares,
		numSquares2,
		//leetcode,
		//endlessCheng,
	}

	for _, f := range fs {
		err := tests.RunFunc(t, f, samples, targetCaseNum)
		if err != nil {
			t.Error(err)
		}
	}
}

var samples = `
12
3

13
2


`
