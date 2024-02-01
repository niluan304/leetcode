package main

import (
	"testing"

	"github.com/niluan304/leetcode/tests"
)

func Test_generate_parentheses(t *testing.T) {
	targetCaseNum := 0 // -1

	fs := []func(n int) []string{
		generateParenthesis,
	}

	for _, f := range fs {
		err := tests.RunFunc(t, f, samples, targetCaseNum)
		if err != nil {
			t.Error(err)
		}
	}
}

var samples = `
3
["((()))","(()())","(())()","()(())","()()()"]

1
["()"]

`
