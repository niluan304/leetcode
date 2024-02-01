package main

import (
	"testing"

	"github.com/niluan304/leetcode/tests"
)

func Test_minimum_cost_to_cut_a_stick(t *testing.T) {
	targetCaseNum := 0 // -1

	fs := []interface{}{
		minCost,
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
[1,3,4,5]
16

9
[5,6,1,4,2]
22

`
