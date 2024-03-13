package main

import (
	"testing"

	"github.com/niluan304/leetcode/tests"
)

func Test_maximumOddBinaryNumber(t *testing.T) {
	targetCaseNum := 0 // -1

	fs := []func(s string) string{
		// bruteForce,
		maximumOddBinaryNumber,
		// maximumOddBinaryNumber2,
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
"010"
"001"

"0101"
"1001"


`
