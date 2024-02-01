package main

import (
	"testing"

	"github.com/niluan304/leetcode/tests"
)

func Test_keyboard_row(t *testing.T) {
	targetCaseNum := 0 // -1

	fs := []interface{}{
		findWords,
	}

	for _, f := range fs {
		err := tests.RunFunc(t, f, samples, targetCaseNum)
		if err != nil {
			t.Error(err)
		}
	}
}

var samples = `
["Hello","Alaska","Dad","Peace"]
["Alaska","Dad"]

["omk"]
[]

["adsdf","sfd"]
["adsdf","sfd"]

`
