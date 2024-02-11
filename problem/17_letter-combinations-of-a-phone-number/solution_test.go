package main

import (
	"testing"

	"github.com/niluan304/leetcode/tests"
)

func Test_letterCombinations(t *testing.T) {
	targetCaseNum := 0 // -1

	fs := []func(digits string) []string{
		//bruteForce,
		letterCombinations,
		//letterCombinations2,
		//leetcode,
		//endlessCheng,
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
