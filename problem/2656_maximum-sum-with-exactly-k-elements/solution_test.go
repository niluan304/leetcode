package main

import (
	"testing"

	"github.com/niluan304/leetcode/tests"
)

func Test_maximum_sum_with_exactly_k_elements(t *testing.T) {
	targetCaseNum := 0 // -1

	fs := []interface{}{
		maximizeSum,
	}

	for _, f := range fs {
		err := tests.RunFunc(t, f, samples, targetCaseNum)
		if err != nil {
			t.Error(err)
		}
	}
}

var samples = `
[1,2,3,4,5]
3
18

[5,5,5]
2
11

`
