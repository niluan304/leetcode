package main

import (
	"testing"

	"github.com/niluan304/leetcode/tests"
)

func Test_longest_common_prefix(t *testing.T) {
	targetCaseNum := 0 // -1

	fs := []interface{}{
		longestCommonPrefix,
	}

	for _, f := range fs {
		err := tests.RunFunc(t, f, samples, targetCaseNum)
		if err != nil {
			t.Error(err)
		}
	}
}

var samples = `
["flower","flow","flight"]
"fl"

["dog","racecar","car"]
""

`
