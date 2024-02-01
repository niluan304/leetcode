package main

import (
	"testing"

	"github.com/niluan304/leetcode/tests"
)

func Test_palindrome_partitioning_ii(t *testing.T) {
	targetCaseNum := 0 // -1

	fs := []interface{}{
		minCut,
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
1

"a"
0

"ab"
1

`
