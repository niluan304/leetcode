package main

import (
	"testing"

	"github.com/niluan304/leetcode/tests"
)

func Test_longest_substring_without_repeating_characters(t *testing.T) {
	targetCaseNum := 0 // -1

	fs := []interface{}{
		lengthOfLongestSubstring,
	}

	for _, f := range fs {
		err := tests.RunFunc(t, f, samples, targetCaseNum)
		if err != nil {
			t.Error(err)
		}
	}
}

var samples = `
"abcabcbb"
3

"bbbbb"
1

"pwwkew"
3

`
