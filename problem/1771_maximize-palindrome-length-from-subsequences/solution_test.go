package main

import (
	"testing"

	"github.com/niluan304/leetcode/tests"
)

func Test_maximize_palindrome_length_from_subsequences(t *testing.T) {
	targetCaseNum := 0 // -1

	fs := []interface{}{
		longestPalindrome,
	}

	for _, f := range fs {
		err := tests.RunFunc(t, f, samples, targetCaseNum)
		if err != nil {
			t.Error(err)
		}
	}
}

var samples = `
"cacb"
"cbba"
5

"ab"
"ab"
3

"aa"
"bb"
0

`
