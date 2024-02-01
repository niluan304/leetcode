package main

import (
	"testing"

	"github.com/niluan304/leetcode/tests"
)

func Test_reverse_string(t *testing.T) {
	targetCaseNum := 0 // -1

	fs := []interface{}{
		reverseString,
	}

	for _, f := range fs {
		err := tests.RunFunc(t, f, samples, targetCaseNum)
		if err != nil {
			t.Error(err)
		}
	}
}

var samples = `
["h","e","l","l","o"]
["o","l","l","e","h"]

["H","a","n","n","a","h"]
["h","a","n","n","a","H"]

`
