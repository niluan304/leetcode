package main

import (
	"testing"

	"github.com/niluan304/leetcode/tests"
)

func Test_isMatch(t *testing.T) {
	targetCaseNum := 0 // -1

	fs := []func(s string, p string) bool{
		//bruteForce,
		isMatch,
		isMatch2,
		//leetcode,
		//endlessCheng,
	}

	for _, f := range fs {
		err := tests.RunFunc(t, f, samples, targetCaseNum)
		if err != nil {
			t.Error(err)
		}
	}
}

var samples = `
"aa"
"a"
false

"aa"
"a*"
true

"ab"
".*"
true

"ppppp"
"p*pppppp"
false

"xxxxxabcdxxxx"
".*ad.*"
false

"pppp"
"p*pppp"
true

"apppp"
"ap*pppp"
true

"a"
"ab*"
true

"a"
"a.*"
true

"aaaaaaaaaaaaaaaaaaab"
"a*a*a*a*a*a*a*a*a*a*a*a*a*a*a*"
false

"aaaaaaaaaaaaaaaaaaa"
"a*a*a*a*a*a*a*a*a*a*a*a*a*a*b"
false

"abc"
"a***abc"
true

`
