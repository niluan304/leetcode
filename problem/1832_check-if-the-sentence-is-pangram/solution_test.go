package main

import (
	"testing"

	"github.com/niluan304/leetcode/tests"
)

func Test_check_if_the_sentence_is_pangram(t *testing.T) {
	targetCaseNum := 0 // -1

	fs := []interface{}{
		checkIfPangram,
	}

	for _, f := range fs {
		err := tests.RunFunc(t, f, samples, targetCaseNum)
		if err != nil {
			t.Error(err)
		}
	}
}

var samples = `
"thequickbrownfoxjumpsoverthelazydog"
true

"leetcode"
false

`
