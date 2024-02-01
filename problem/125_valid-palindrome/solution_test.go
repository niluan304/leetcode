package main

import (
	"testing"

	"github.com/niluan304/leetcode/tests"
)

func Test_valid_palindrome(t *testing.T) {
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
"A man, a plan, a canal: Panama"
true

"race a car"
false

" "
true

`
