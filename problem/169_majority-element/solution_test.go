package main

import (
	"testing"

	"github.com/niluan304/leetcode/tests"
)

func Test_majority_element(t *testing.T) {
	targetCaseNum := 0 // -1

	fs := []interface{}{
		majorityElement,
	}

	for _, f := range fs {
		err := tests.RunFunc(t, f, samples, targetCaseNum)
		if err != nil {
			t.Error(err)
		}
	}
}

var samples = `
[3,2,3]
3

[2,2,1,1,1,2,2]
2

`
