package main

import (
	"testing"

	"github.com/niluan304/leetcode/tests"
)

func Test_findPrimePairs(t *testing.T) {
	targetCaseNum := 0 // -1

	fs := []func(n int) [][]int{
		//bruteForce,
		findPrimePairs,
		//findPrimePairs2,
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
10
[[3,7],[5,5]]

2
[]


`
