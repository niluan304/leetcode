package main

import (
	"testing"

	"github.com/niluan304/leetcode/tests"
)

func Test_add_strings(t *testing.T) {
	targetCaseNum := 0 // -1

	fs := []interface{}{
		addStrings,
	}

	for _, f := range fs {
		err := tests.RunFunc(t, f, samples, targetCaseNum)
		if err != nil {
			t.Error(err)
		}
	}
}

var samples = `
"11"
"123"
"134"

"456"
"77"
"533"

"0"
"0"
"0"

`
