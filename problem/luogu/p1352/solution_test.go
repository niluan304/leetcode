package main

import (
	"testing"

	"github.com/niluan304/leetcode/tests"
)

func Test_string_transformation(t *testing.T) {
	targetCaseNum := 0 // -1

	fs := []interface{}{
		withoutLeader,
	}

	for _, f := range fs {
		err := tests.RunFunc(t, f, samples, targetCaseNum)
		if err != nil {
			t.Error(err)
		}
	}
}

var samples = `
7
[1,1,1,1,1,1,1]
[[1,3],[2,3],[6,4],[7,4],[4,5],[3,5]]
5
`
