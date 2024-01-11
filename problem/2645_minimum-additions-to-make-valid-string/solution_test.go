package main

import (
	"testing"

	"github.com/niluan304/leetcode/tests"
)

func Test_addMinimum(t *testing.T) {
	targetCaseNum := 0 // -1

	fs := []func(word string) int{
		addMinimum,
		addMinimum2,
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
"b"
2

"aaa"
6

"abc"
0


`
