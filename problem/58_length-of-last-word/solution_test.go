package main

import (
	"testing"

	"github.com/niluan304/leetcode/tests"
)

func Test_length_of_last_word(t *testing.T) {
	targetCaseNum := 0 // -1

	fs := []interface{}{
		lengthOfLastWord,
	}

	for _, f := range fs {
		err := tests.RunFunc(t, f, samples, targetCaseNum)
		if err != nil {
			t.Error(err)
		}
	}
}

var samples = `
"Hello World"
5

"   fly me   to   the moon  "
4

"luffy is still joyboy"
6

`
