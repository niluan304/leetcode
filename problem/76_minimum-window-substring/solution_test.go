package main

import (
	"testing"

	"github.com/niluan304/leetcode/tests"
)

func Test_minimum_window_substring(t *testing.T) {
	targetCaseNum := 0 // -1

	fs := []interface{}{
		minWindow,
	}

	for _, f := range fs {
		err := tests.RunFunc(t, f, samples, targetCaseNum)
		if err != nil {
			t.Error(err)
		}
	}
}

var samples = `
"ADOBECODEBANC"
"ABC"
"BANC"

"a"
"a"
"a"

"a"
"aa"
""

`
