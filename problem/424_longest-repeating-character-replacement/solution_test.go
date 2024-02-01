package main

import (
	"testing"

	"github.com/niluan304/leetcode/tests"
)

func Test_longest_repeating_character_replacement(t *testing.T) {
	targetCaseNum := 0 // -1

	fs := []interface{}{
		characterReplacement,
	}

	for _, f := range fs {
		err := tests.RunFunc(t, f, samples, targetCaseNum)
		if err != nil {
			t.Error(err)
		}
	}
}

var samples = `
"ABAB"
2
4

"AABABBA"
1
4

`
