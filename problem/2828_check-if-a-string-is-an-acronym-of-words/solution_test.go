package main

import (
	"testing"

	"github.com/niluan304/leetcode/tests"
)

func Test_main(t *testing.T) {
	targetCaseNum := 0 // -1

	fs := []func(words []string, s string) bool{
		isAcronym,
	}

	for _, f := range fs {
		err := tests.RunFunc(t, f, samples, targetCaseNum)
		if err != nil {
			t.Error(err)
		}
	}
}

var samples = `
["alice","bob","charlie"]
"abc"
true

["an","apple"]
"a"
false

["never","gonna","give","up","on","you"]
"ngguoy"
true

`
