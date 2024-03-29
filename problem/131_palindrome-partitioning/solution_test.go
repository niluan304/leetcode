package main

import (
	"testing"

	"github.com/niluan304/leetcode/tests"
)

func Test_partition(t *testing.T) {
	targetCaseNum := 0 // -1

	fs := []func(s string) [][]string{
		//bruteForce,
		partition,
		partition2,
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
"aab"
[["a","a","b"],["aa","b"]]

"a"
[["a"]]


`
