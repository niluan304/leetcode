package main

import (
	"testing"

	"github.com/niluan304/leetcode/tests"
)

func Test_splitWordsBySeparator(t *testing.T) {
	targetCaseNum := 0 // -1

	fs := []func(words []string, separator byte) []string{
		splitWordsBySeparator,
		// splitWordsBySeparator2,
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
["one.two.three","four.five","six"]
"."
["one","two","three","four","five","six"]

["$easy$","$problem$"]
"$"
["easy","problem"]

["|||"]
"|"
[]


`
