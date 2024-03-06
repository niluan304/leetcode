package main

import (
	"testing"

	"github.com/niluan304/leetcode/tests"
)

func Test_divisibilityArray(t *testing.T) {
	targetCaseNum := 0 // -1

	fs := []func(word string, m int) []int{
		// bruteForce,
		divisibilityArray,
		// divisibilityArray2,
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
"998244353"
3
[1,1,0,0,0,1,1,0,0]

"1010"
10
[0,1,0,1]


`
