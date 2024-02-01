package main

import (
	"testing"

	"github.com/niluan304/leetcode/tests"
)

func Test_letter_combinations_of_a_phone_number(t *testing.T) {
	targetCaseNum := 0 // -1

	fs := []interface{}{
		letterCombinations,
	}

	for _, f := range fs {
		err := tests.RunFunc(t, f, samples, targetCaseNum)
		if err != nil {
			t.Error(err)
		}
	}
}

var samples = `
"23"
["ad","ae","af","bd","be","bf","cd","ce","cf"]

""
[]

"2"
["a","b","c"]

`
