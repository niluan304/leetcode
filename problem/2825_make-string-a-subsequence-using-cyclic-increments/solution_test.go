package main

import (
	"testing"

	"github.com/niluan304/leetcode/tests"
)

func Test_canMakeSubsequence(t *testing.T) {
	targetCaseNum := 0 // -1

	fs := []func(str1 string, str2 string) bool{
		canMakeSubsequence,
		// canMakeSubsequence2,
		// leetcode,
		// endlessCheng,
	}

	for _, f := range fs {
		err := tests.RunFunc(t, f, samples, targetCaseNum)
		if err != nil {
			t.Error(err)
		}
	}
}

var samples = `
"abc"
"ad"
true

"zc"
"ad"
true

"ab"
"d"
false


`
