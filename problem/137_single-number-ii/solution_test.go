package main

import (
	"testing"

	"github.com/niluan304/leetcode/tests"
)

func Test_single_number_ii(t *testing.T) {
	targetCaseNum := 0 // -1

	fs := []interface{}{
		singleNumber,
	}

	for _, f := range fs {
		err := tests.RunFunc(t, f, samples, targetCaseNum)
		if err != nil {
			t.Error(err)
		}
	}
}

var samples = `
[2,2,3,2]
3

[0,1,0,1,0,1,99]
99

`
