package main

import (
	"testing"

	"github.com/niluan304/leetcode/tests"
)

func Test_integer_break(t *testing.T) {
	targetCaseNum := 0 // -1

	fs := []interface{}{
		integerBreak,
	}

	for _, f := range fs {
		err := tests.RunFunc(t, f, samples, targetCaseNum)
		if err != nil {
			t.Error(err)
		}
	}
}

var samples = `
2
1

10
36

`
