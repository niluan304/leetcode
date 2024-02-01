package main

import (
	"testing"

	"github.com/niluan304/leetcode/tests"
)

func Test_take_k_of_each_character_from_left_and_right(t *testing.T) {
	targetCaseNum := 0 // -1

	fs := []interface{}{
		takeCharacters,
	}

	for _, f := range fs {
		err := tests.RunFunc(t, f, samples, targetCaseNum)
		if err != nil {
			t.Error(err)
		}
	}
}

var samples = `
"aabaaaacaabc"
2
8

"a"
1
-1

`
