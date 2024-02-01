package main

import (
	"testing"

	"github.com/niluan304/leetcode/tests"
)

func Test_word_pattern(t *testing.T) {
	targetCaseNum := 0 // -1

	fs := []interface{}{
		wordPattern,
	}

	for _, f := range fs {
		err := tests.RunFunc(t, f, samples, targetCaseNum)
		if err != nil {
			t.Error(err)
		}
	}
}

var samples = `
"abba"
"dog cat cat dog"
true

"abba"
"dog cat cat fish"
false

"aaaa"
"dog cat cat dog"
false

`
