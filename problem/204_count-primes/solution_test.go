package main

import (
	"testing"

	"github.com/niluan304/leetcode/tests"
)

func Test_countPrimes(t *testing.T) {
	targetCaseNum := 0 // -1

	fs := []func(n int) int{
		countPrimes,
		countPrimes2,
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
10
4

0
0

1
0

5000000
348513


`
