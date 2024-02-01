package main

import (
	"testing"

	"github.com/niluan304/leetcode/tests"
)

func Test_find_the_longest_balanced_substring_of_a_binary_string(t *testing.T) {
	targetCaseNum := 0 // -1

	fs := []interface{}{
		findTheLongestBalancedSubstring,
	}

	for _, f := range fs {
		err := tests.RunFunc(t, f, samples, targetCaseNum)
		if err != nil {
			t.Error(err)
		}
	}
}

var samples = `
"01000111"
6

"00111"
4

"111"
0

`
