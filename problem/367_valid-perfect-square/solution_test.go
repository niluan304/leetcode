package main

import (
	"testing"

	"github.com/niluan304/leetcode/tests"
)

func Test_valid_perfect_square(t *testing.T) {
	targetCaseNum := 0 // -1

	fs := []interface{}{
		isPerfectSquare,
	}

	for _, f := range fs {
		err := tests.RunFunc(t, f, samples, targetCaseNum)
		if err != nil {
			t.Error(err)
		}
	}
}

var samples = `
16
true

14
false

`
