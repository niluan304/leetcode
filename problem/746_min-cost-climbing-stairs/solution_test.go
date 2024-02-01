package main

import (
	"testing"

	"github.com/niluan304/leetcode/tests"
)

func Test_min_cost_climbing_stairs(t *testing.T) {
	targetCaseNum := 0 // -1

	fs := []interface{}{
		minCostClimbingStairs,
		minCostClimbingStairs2,
		minCostClimbingStairs3,
	}

	for _, f := range fs {
		err := tests.RunFunc(t, f, samples, targetCaseNum)
		if err != nil {
			t.Error(err)
		}
	}
}

var samples = `
[10,15,20]
15

[1,100,1,1,1,100,1,1,100,1]
6

`
