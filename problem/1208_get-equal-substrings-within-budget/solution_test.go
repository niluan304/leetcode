package main

import (
	"testing"

	"github.com/niluan304/leetcode/tests"
)

func Test_get_equal_substrings_within_budget(t *testing.T) {
	targetCaseNum := 0 // -1

	fs := []interface{}{
		equalSubstring,
	}

	for _, f := range fs {
		err := tests.RunFunc(t, f, samples, targetCaseNum)
		if err != nil {
			t.Error(err)
		}
	}
}

var samples = `
"abcd"
"bcdf"
3
3

"abcd"
"cdef"
3
1

"abcd"
"acde"
0
1

`
