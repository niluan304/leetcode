package main

import (
	"testing"

	"github.com/niluan304/leetcode/tests"
)

func Test_palindrome_partitioning(t *testing.T) {
	targetCaseNum := 0 // -1

	fs := []interface{}{
		partition,
	}

	for _, f := range fs {
		err := tests.RunFunc(t, f, samples, targetCaseNum)
		if err != nil {
			t.Error(err)
		}
	}
}

var samples = `
"aab"
[["a","a","b"],["aa","b"]]

"a"
[["a"]]

`
