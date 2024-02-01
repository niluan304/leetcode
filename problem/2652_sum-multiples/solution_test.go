package main

import (
	"testing"

	"github.com/niluan304/leetcode/tests"
)

func Test_sum_multiples(t *testing.T) {
	targetCaseNum := 0 // -1

	fs := []interface{}{
		sumOfMultiples,
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
21

10
40

9
30

`
