package main

import (
	"testing"

	"github.com/niluan304/leetcode/tests"
)

func Test_swap_for_longest_repeated_character_substring(t *testing.T) {
	targetCaseNum := 0 // -1

	fs := []interface{}{
		maxRepOpt1,
	}

	for _, f := range fs {
		err := tests.RunFunc(t, f, samples, targetCaseNum)
		if err != nil {
			t.Error(err)
		}
	}
}

var samples = `
"ababa"
3

"aaabaaa"
6

"aaaaa"
5

`
