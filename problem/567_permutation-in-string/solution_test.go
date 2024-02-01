package main

import (
	"testing"

	"github.com/niluan304/leetcode/tests"
)

func Test_permutation_in_string(t *testing.T) {
	targetCaseNum := 0 // -1

	fs := []interface{}{
		checkInclusion,
	}

	for _, f := range fs {
		err := tests.RunFunc(t, f, samples, targetCaseNum)
		if err != nil {
			t.Error(err)
		}
	}
}

var samples = `
"ab"
"eidbaooo"
true

"ab"
"eidboaoo"
false

`
