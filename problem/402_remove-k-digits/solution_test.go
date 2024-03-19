package main

import (
	"testing"

	"github.com/niluan304/leetcode/tests"
)

func Test_removeKdigits(t *testing.T) {
	targetCaseNum := 0 // -1

	fs := []func(num string, k int) string{
		// bruteForce,
		removeKdigits,
		// removeKdigits2,
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
"1432219"
3
"1219"

"10200"
1
"200"

"10"
2
"0"

"9"
1
"0"

`
