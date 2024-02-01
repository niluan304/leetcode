package main

import (
	"testing"

	"github.com/niluan304/leetcode/tests"
)

func Test_string_to_integer_atoi(t *testing.T) {
	targetCaseNum := 0 // -1

	fs := []interface{}{
		myAtoi,
	}

	for _, f := range fs {
		err := tests.RunFunc(t, f, samples, targetCaseNum)
		if err != nil {
			t.Error(err)
		}
	}
}

var samples = `
"42"
42

"   -42"
-42

"4193 with words"
4193

`
