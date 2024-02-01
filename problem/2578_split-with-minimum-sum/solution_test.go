package main

import (
	"testing"

	"github.com/niluan304/leetcode/tests"
)

func Test_split_with_minimum_sum(t *testing.T) {
	targetCaseNum := 0 // -1

	fs := []interface{}{
		splitNum,
	}

	for _, f := range fs {
		err := tests.RunFunc(t, f, samples, targetCaseNum)
		if err != nil {
			t.Error(err)
		}
	}
}

var samples = `
4325
59

687
75

`
