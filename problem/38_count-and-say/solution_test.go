package main

import (
	"testing"

	"github.com/niluan304/leetcode/tests"
)

func Test_count_and_say(t *testing.T) {
	targetCaseNum := 0 // -1

	fs := []interface{}{
		countAndSay,
	}

	for _, f := range fs {
		err := tests.RunFunc(t, f, samples, targetCaseNum)
		if err != nil {
			t.Error(err)
		}
	}
}

var samples = `
1
"1"

4
"1211"

`
