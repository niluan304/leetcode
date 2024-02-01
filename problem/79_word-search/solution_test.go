package main

import (
	"testing"

	"github.com/niluan304/leetcode/tests"
)

func Test_word_search(t *testing.T) {
	targetCaseNum := 0 // -1

	fs := []interface{}{
		exist,
	}

	for _, f := range fs {
		err := tests.RunFunc(t, f, samples, targetCaseNum)
		if err != nil {
			t.Error(err)
		}
	}
}

var samples = `
[["A","B","C","E"],["S","F","C","S"],["A","D","E","E"]]
"ABCCED"
true

[["A","B","C","E"],["S","F","C","S"],["A","D","E","E"]]
"SEE"
true

[["A","B","C","E"],["S","F","C","S"],["A","D","E","E"]]
"ABCB"
false

`
