package main

import (
	"testing"

	"github.com/niluan304/leetcode/tests"
)

func Test_ones_and_zeroes(t *testing.T) {
	targetCaseNum := 0 // -1

	fs := []interface{}{
		findMaxForm,
	}

	for _, f := range fs {
		err := tests.RunFunc(t, f, samples, targetCaseNum)
		if err != nil {
			t.Error(err)
		}
	}
}

var samples = `
["10","0001","111001","1","0"]
5
3
4

["10","0","1"]
1
1
2

`
