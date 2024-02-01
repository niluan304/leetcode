package main

import (
	"testing"

	"github.com/niluan304/leetcode/tests"
)

func Test_find_all_anagrams_in_a_string(t *testing.T) {
	targetCaseNum := 0 // -1

	fs := []interface{}{
		findAnagrams,
	}

	for _, f := range fs {
		err := tests.RunFunc(t, f, samples, targetCaseNum)
		if err != nil {
			t.Error(err)
		}
	}
}

var samples = `
"cbaebabacd"
"abc"
[0,6]

"abab"
"ab"
[0,1,2]

`
