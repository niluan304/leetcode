package main

import (
	"testing"

	"github.com/niluan304/leetcode/tests"
)

func Test_sum_of_beauty_of_all_substrings(t *testing.T) {
	targetCaseNum := 0 // -1

	fs := []interface{}{
		beautySum,
	}

	for _, f := range fs {
		err := tests.RunFunc(t, f, samples, targetCaseNum)
		if err != nil {
			t.Error(err)
		}
	}
}

var samples = `
"aabcb"
5

"aabcbaa"
17

`
