package main

import (
	"testing"

	"github.com/niluan304/leetcode/tests"
)

func Test_decode_string(t *testing.T) {
	targetCaseNum := 0 // -1

	fs := []interface{}{
		decodeString,
	}

	for _, f := range fs {
		err := tests.RunFunc(t, f, samples, targetCaseNum)
		if err != nil {
			t.Error(err)
		}
	}
}

var samples = `
"3[a]2[bc]"
"aaabcbc"

"3[a2[c]]"
"accaccacc"

"2[abc]3[cd]ef"
"abcabccdcdcdef"

`
