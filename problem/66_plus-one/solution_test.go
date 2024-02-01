package main

import (
	"testing"

	"github.com/niluan304/leetcode/tests"
)

func Test_plus_one(t *testing.T) {
	targetCaseNum := 0 // -1

	fs := []interface{}{
		plusOne,
	}

	for _, f := range fs {
		err := tests.RunFunc(t, f, samples, targetCaseNum)
		if err != nil {
			t.Error(err)
		}
	}
}

var samples = `
[1,2,3]
[1,2,4]

[4,3,2,1]
[4,3,2,2]

[9]
[1,0]

`
