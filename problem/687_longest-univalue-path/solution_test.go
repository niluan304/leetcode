package main

import (
	"testing"

	"github.com/niluan304/leetcode/tests"
)

func Test_longest_univalue_path(t *testing.T) {
	targetCaseNum := 0 // -1

	fs := []interface{}{
		longestUnivaluePath,
	}

	for _, f := range fs {
		err := tests.RunFunc(t, f, samples, targetCaseNum)
		if err != nil {
			t.Error(err)
		}
	}
}

var samples = `
[5,4,5,1,1,null,5]
2

[1,4,5,4,4,null,5]
2

`
