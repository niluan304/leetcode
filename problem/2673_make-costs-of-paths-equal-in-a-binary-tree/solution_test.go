package main

import (
	"testing"

	"github.com/niluan304/leetcode/tests"
)

func Test_minIncrements(t *testing.T) {
	targetCaseNum := 0 // -1

	fs := []func(n int, cost []int) int{
		// bruteForce,
		minIncrements,
		minIncrements2,
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
7
[1,5,2,2,3,3,1]
6

3
[5,3,3]
0


`
