package main

import (
	"testing"

	"github.com/niluan304/leetcode/tests"
)

func Test_repeatLimitedString(t *testing.T) {
	targetCaseNum := 0 // -1

	fs := []func(s string, repeatLimit int) string{
		repeatLimitedString,
		// repeatLimitedString2,
		leetcode,
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
"cczazcc"
3
"zzcccac"

"aababab"
2
"bbabaa"

"zzzzzzzzzzzzzzzzzzzzzzzzaa"
2
"zzazzazz"

"zzzzzzzzzzzzzzzzzzzzzzzza"
1
"zaz"


`
