package main

import (
	"testing"

	"github.com/niluan304/leetcode/tests"
)

func Test_maximum_alternating_subsequence_sum(t *testing.T) {
	targetCaseNum := 0 // -1

	fs := []interface{}{
		maxAlternatingSum,
	}

	for _, f := range fs {
		err := tests.RunFunc(t, f, samples, targetCaseNum)
		if err != nil {
			t.Error(err)
		}
	}
}

var samples = `
[4,2,5,3]
7

[5,6,7,8]
8

[6,2,1,2,4,5]
10

`
