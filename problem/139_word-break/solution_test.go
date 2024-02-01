package main

import (
	"testing"

	"github.com/niluan304/leetcode/tests"
)

func Test_word_break(t *testing.T) {
	targetCaseNum := 0 // -1

	fs := []interface{}{
		wordBreak,
	}

	for _, f := range fs {
		err := tests.RunFunc(t, f, samples, targetCaseNum)
		if err != nil {
			t.Error(err)
		}
	}
}

var samples = `
"leetcode"
["leet","code"]
true

"applepenapple"
["apple","pen"]
true

"catsandog"
["cats","dog","sand","and","cat"]
false

`
