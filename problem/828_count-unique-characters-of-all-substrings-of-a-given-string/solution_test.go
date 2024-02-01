package main

import (
	"testing"

	"github.com/niluan304/leetcode/tests"
)

func Test_main(t *testing.T) {
	targetCaseNum := 0 // -1

	fs := []interface{}{
		uniqueLetterString,
		uniqueLetterString2,
	}

	for _, f := range fs {
		err := tests.RunFunc(t, f, samples, targetCaseNum)
		if err != nil {
			t.Error(err)
		}
	}
}

var samples = `
"ABC"
10

"ABA"
8

"LEETCODE"
92

"LEXEFE"
38

"LEXEFEEE"
58

"LEXEFEEEF"
67
`
