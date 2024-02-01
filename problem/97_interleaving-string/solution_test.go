package main

import (
	"testing"

	"github.com/niluan304/leetcode/tests"
)

func Test_interleaving_string(t *testing.T) {
	targetCaseNum := 0 // -1

	fs := []interface{}{
		isInterleave,
	}

	for _, f := range fs {
		err := tests.RunFunc(t, f, samples, targetCaseNum)
		if err != nil {
			t.Error(err)
		}
	}
}

var samples = `
"aabcc"
"dbbca"
"aadbbcbcac"
true

"aabcc"
"dbbca"
"aadbbbaccc"
false

""
""
""
true

`
