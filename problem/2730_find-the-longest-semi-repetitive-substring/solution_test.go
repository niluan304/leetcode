package main

import (
	"testing"

	"github.com/niluan304/leetcode/tests"
)

func Test_main(t *testing.T) {
	targetCaseNum := 0 // -1

	fs := []func(s string) int{
		longestSemiRepetitiveSubstring,
	}

	for _, f := range fs {
		err := tests.RunFunc(t, f, samples, targetCaseNum)
		if err != nil {
			t.Error(err)
		}
	}
}

var samples = `
"52233"
4

"5494"
4

"1111111"
2

"1"
1

"22"
2

"21"
2

"333"
2


`
