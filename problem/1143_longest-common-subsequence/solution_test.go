package main

import (
	"testing"

	"github.com/niluan304/leetcode/tests"
)

func Test_longest_common_subsequence(t *testing.T) {
	targetCaseNum := 0 // -1

	fs := []interface{}{
		longestCommonSubsequence,
	}

	for _, f := range fs {
		err := tests.RunFunc(t, f, samples, targetCaseNum)
		if err != nil {
			t.Error(err)
		}
	}
}

var samples = `
"abcde"
"ace"
3

"abc"
"abc"
3

"abc"
"def"
0

`
