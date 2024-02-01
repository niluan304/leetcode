package main

import (
	"testing"

	"github.com/niluan304/leetcode/tests"
)

func Test_reducing_dishes(t *testing.T) {
	targetCaseNum := 0 // -1

	fs := []interface{}{
		maxSatisfaction,
	}

	for _, f := range fs {
		err := tests.RunFunc(t, f, samples, targetCaseNum)
		if err != nil {
			t.Error(err)
		}
	}
}

var samples = `
[-1,-8,0,5,-9]
14

[4,3,2]
20

[-1,-4,-5]
0

`
