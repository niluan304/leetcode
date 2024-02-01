package main

import (
	"testing"

	"github.com/niluan304/leetcode/tests"
)

func Test_simplify_path(t *testing.T) {
	targetCaseNum := 0 // -1

	fs := []interface{}{
		simplifyPath,
	}

	for _, f := range fs {
		err := tests.RunFunc(t, f, samples, targetCaseNum)
		if err != nil {
			t.Error(err)
		}
	}
}

var samples = `
"/home/"
"/home"

"/../"
"/"

"/home//foo/"
"/home/foo"

`
