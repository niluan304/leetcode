package main

import (
	"testing"

	"github.com/niluan304/leetcode/tests"
)

func Test_longest_palindrome(t *testing.T) {
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
"abccccdd"
7

"a"
1

`
