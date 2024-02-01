package main

import (
	"testing"

	"github.com/niluan304/leetcode/tests"
)

func Test_count_the_number_of_vowel_strings_in_range(t *testing.T) {
	targetCaseNum := 0 // -1

	fs := []interface{}{
		vowelStrings,
	}

	for _, f := range fs {
		err := tests.RunFunc(t, f, samples, targetCaseNum)
		if err != nil {
			t.Error(err)
		}
	}
}

var samples = `
["are","amy","u"]
0
2
2

["hey","aeo","mu","ooo","artro"]
1
4
3

`
