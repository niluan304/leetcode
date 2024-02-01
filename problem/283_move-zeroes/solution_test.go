package main

import (
	"testing"

	"github.com/niluan304/leetcode/tests"
)

func Test_move_zeroes(t *testing.T) {
	targetCaseNum := 0 // -1

	fs := []interface{}{
		moveZeroes,
	}

	for _, f := range fs {
		err := tests.RunFunc(t, f, samples, targetCaseNum)
		if err != nil {
			t.Error(err)
		}
	}
}

var samples = `
[0,1,0,3,12]
[1,3,12,0,0]

[0]
[0]

`
