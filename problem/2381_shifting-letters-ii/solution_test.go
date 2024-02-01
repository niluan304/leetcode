package main

import (
	"testing"

	"github.com/niluan304/leetcode/tests"
)

func Test_shifting_letters_ii(t *testing.T) {
	targetCaseNum := 0 // -1

	fs := []interface{}{
		shiftingLetters,
	}

	for _, f := range fs {
		err := tests.RunFunc(t, f, samples, targetCaseNum)
		if err != nil {
			t.Error(err)
		}
	}
}

var samples = `
"abc"
[[0,1,0],[1,2,1],[0,2,1]]
"ace"

"dztz"
[[0,0,0],[1,1,1]]
"catz"

`
