package main

import (
	"testing"

	"github.com/niluan304/leetcode/tests"
)

func Test_minimumBeautifulSubstrings(t *testing.T) {
	targetCaseNum := 0 // -1

	fs := []func(s string) int{
		//bruteForce,
		minimumBeautifulSubstrings,
		minimumBeautifulSubstrings2,
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
"1011"
2

"111"
3

"0"
-1


`
