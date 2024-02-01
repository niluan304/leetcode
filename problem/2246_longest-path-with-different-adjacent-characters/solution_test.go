package main

import (
	"testing"

	"github.com/niluan304/leetcode/tests"
)

func Test_longest_path_with_different_adjacent_characters(t *testing.T) {
	targetCaseNum := 0 // -1

	fs := []interface{}{
		longestPath,
	}

	for _, f := range fs {
		err := tests.RunFunc(t, f, samples, targetCaseNum)
		if err != nil {
			t.Error(err)
		}
	}
}

var samples = `
[-1,0,0,1,1,2]
"abacbe"
3

[-1,0,0,0]
"aabc"
3

`
