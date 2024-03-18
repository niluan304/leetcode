package main

import (
	"testing"

	"github.com/niluan304/leetcode/tests"
)

func Test_Constructor(t *testing.T) {
	targetCaseNum := 0 // -1

	fs := []interface{}{
		// bruteForce,
		Constructor,
		// Constructor2,
		// leetcode,
		// endlessCheng,
	}

	for _, f := range fs {
		err := tests.RunClass(t, f, samples, targetCaseNum)
		if err != nil {
			t.Error(err)
		}
	}
}

var samples = `
["NumArray", "sumRange", "sumRange", "sumRange"]
[[[-2, 0, 3, -5, 2, -1]], [0, 2], [2, 5], [0, 5]]
[null, 1, -1, -3]


`
