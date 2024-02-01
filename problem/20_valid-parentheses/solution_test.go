package main

import (
	"testing"

	"github.com/niluan304/leetcode/tests"
)

func Test_valid_parentheses(t *testing.T) {
	targetCaseNum := 0 // -1

	fs := []interface{}{
		isValid,
	}

	for _, f := range fs {
		err := tests.RunFunc(t, f, samples, targetCaseNum)
		if err != nil {
			t.Error(err)
		}
	}
}

var samples = `
"()"
true

"()[]{}"
true

"(]"
false

`
