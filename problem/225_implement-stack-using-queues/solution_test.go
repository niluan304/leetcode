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
["MyStack", "push", "push", "top", "pop", "empty"]
[[], [1], [2], [], [], []]
[null, null, null, 2, 2, false]


`
