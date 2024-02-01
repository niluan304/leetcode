package main

import (
	"testing"

	"github.com/niluan304/leetcode/tests"
)

func Test_palindrome_number(t *testing.T) {
	targetCaseNum := 0 // -1

	fs := []interface{}{
		isPalindrome,
	}

	for _, f := range fs {
		err := tests.RunFunc(t, f, samples, targetCaseNum)
		if err != nil {
			t.Error(err)
		}
	}
}

var samples = `
121
true

-121
false

10
false

`
