package main

import (
	"testing"

	"github.com/niluan304/leetcode/tests"
)

func Test_longest_palindromic_subsequence(t *testing.T) {
	targetCaseNum := 0 // -1

	fs := []interface{}{
		longestPalindromeSubseq,
	}

	for _, f := range fs {
		err := tests.RunFunc(t, f, samples, targetCaseNum)
		if err != nil {
			t.Error(err)
		}
	}
}

var samples = `
"bbbab"
4

"cbbd"
2

`
