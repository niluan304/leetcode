package main

import (
	"testing"

	"github.com/niluan304/leetcode/tests"
)

func Test_min_count(t *testing.T) {
	targetCaseNum := 0 // -1

	fs := []interface{}{
		minCount,
	}

	for _, f := range fs {
		err := tests.RunFunc(t, f, samples, targetCaseNum)
		if err != nil {
			t.Error(err)
		}
	}
}

var samples = `
[4,2,1]
4

[2,3,10]
8
`
