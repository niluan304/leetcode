package main

import (
	"testing"

	"github.com/niluan304/leetcode/tests"
)

func Test_max_consecutive_ones_iii(t *testing.T) {
	targetCaseNum := 0 // -1

	fs := []interface{}{
		longestOnes,
	}

	for _, f := range fs {
		err := tests.RunFunc(t, f, samples, targetCaseNum)
		if err != nil {
			t.Error(err)
		}
	}
}

var samples = `
[1,1,1,0,0,0,1,1,1,1,0]
2
6

[0,0,1,1,0,0,1,1,1,0,1,1,0,0,0,1,1,1,1]
3
10

`
