package main

import (
	"testing"

	"github.com/niluan304/leetcode/tests"
)

func Test_sort_colors(t *testing.T) {
	targetCaseNum := 0 // -1

	fs := []interface{}{
		sortColors,
	}

	for _, f := range fs {
		err := tests.RunFunc(t, f, samples, targetCaseNum)
		if err != nil {
			t.Error(err)
		}
	}
}

var samples = `
[2,0,2,1,1,0]
[0,0,1,1,2,2]

[2,0,1]
[0,1,2]

`
