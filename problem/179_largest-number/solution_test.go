package main

import (
	"testing"

	"github.com/niluan304/leetcode/tests"
)

func Test_largest_number(t *testing.T) {
	targetCaseNum := 0 // -1

	fs := []interface{}{
		largestNumber,
	}

	for _, f := range fs {
		err := tests.RunFunc(t, f, samples, targetCaseNum)
		if err != nil {
			t.Error(err)
		}
	}
}

var samples = `
[10,2]
"210"

[3,30,34,5,9]
"9534330"

`
