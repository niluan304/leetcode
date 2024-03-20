package main

import (
	"testing"

	"github.com/niluan304/leetcode/tests"
)

func Test_minNonZeroProduct(t *testing.T) {
	targetCaseNum := 0 // -1

	fs := []func(p int) int{
		// bruteForce,
		minNonZeroProduct,
		// minNonZeroProduct2,
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
1
1

2
6

3
1512

4
581202553
`
