package main

import (
	"testing"

	"github.com/niluan304/leetcode/tests"
)

func Test_wiggle_subsequence(t *testing.T) {
	targetCaseNum := 0 // -1

	fs := []interface{}{
		wiggleMaxLength,
	}

	for _, f := range fs {
		err := tests.RunFunc(t, f, samples, targetCaseNum)
		if err != nil {
			t.Error(err)
		}
	}
}

var samples = `
[1,7,4,9,2,5]
6

[1,17,5,10,13,15,10,5,16,8]
7

[1,2,3,4,5,6,7,8,9]
2

`
