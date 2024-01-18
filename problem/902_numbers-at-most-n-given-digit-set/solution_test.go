package main

import (
	"testing"

	"github.com/niluan304/leetcode/tests"
)

func Test_atMostNGivenDigitSet(t *testing.T) {
	targetCaseNum := 0 // -1

	fs := []func(digits []string, n int) int{
		atMostNGivenDigitSet,
		atMostNGivenDigitSet2,
		atMostNGivenDigitSet3,
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
["1","3","5","7"]
100
20

["1","4","9"]
1000000000
29523

["7"]
8
1


`
