package main

import (
	"testing"

	"github.com/niluan304/leetcode/tests"
)

func Test_letterCasePermutation(t *testing.T) {
	targetCaseNum := 0 // -1

	fs := []func(s string) []string{
		//bruteForce,
		letterCasePermutation,
		letterCasePermutation2,
		//leetcode,
		//endlessCheng,
	}

	for _, f := range fs {
		err := tests.RunFunc(t, f, samples, targetCaseNum)
		if err != nil {
			t.Error(err)
		}
	}
}

var samples = `
"a1b2"
["a1b2","a1B2","A1b2","A1B2"]

"3z4"
["3z4","3Z4"]


`
