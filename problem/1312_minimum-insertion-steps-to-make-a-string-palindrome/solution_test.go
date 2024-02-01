package main

import (
	"testing"

	"github.com/niluan304/leetcode/tests"
)

func Test_minimum_insertion_steps_to_make_a_string_palindrome(t *testing.T) {
	targetCaseNum := 0 // -1

	fs := []interface{}{
		minInsertions,
	}

	for _, f := range fs {
		err := tests.RunFunc(t, f, samples, targetCaseNum)
		if err != nil {
			t.Error(err)
		}
	}
}

var samples = `
"zzazz"
0

"mbadm"
2

"leetcode"
5

`
