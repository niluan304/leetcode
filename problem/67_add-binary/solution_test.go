package main

import (
	"testing"

	"github.com/niluan304/leetcode/tests"
)

func Test_add_binary(t *testing.T) {
	targetCaseNum := 0 // -1

	fs := []interface{}{
		addBinary,
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
"1"
"100"

"1010"
"1011"
"10101"

`
