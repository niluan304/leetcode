package main

import (
	"testing"

	"github.com/niluan304/leetcode/tests"
)

func Test_add_digits(t *testing.T) {
	targetCaseNum := 0 // -1

	fs := []interface{}{
		addDigits,
	}

	for _, f := range fs {
		err := tests.RunFunc(t, f, samples, targetCaseNum)
		if err != nil {
			t.Error(err)
		}
	}
}

var samples = `
38
2

0
0

`
