package main

import (
	"testing"

	"github.com/niluan304/leetcode/tests"
)

func Test_h_index_ii(t *testing.T) {
	targetCaseNum := 0 // -1

	fs := []interface{}{
		hIndex,
	}

	for _, f := range fs {
		err := tests.RunFunc(t, f, samples, targetCaseNum)
		if err != nil {
			t.Error(err)
		}
	}
}

var samples = `
[0,1,3,5,6]
3

[1,2,100]
2

`
