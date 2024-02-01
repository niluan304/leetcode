package main

import (
	"testing"

	"github.com/niluan304/leetcode/tests"
)

func Test_main(t *testing.T) {
	targetCaseNum := 0 // -1

	fs := []func(chargeTimes []int, runningCosts []int, budget int64) int{
		maximumRobots,
	}

	for _, f := range fs {
		err := tests.RunFunc(t, f, samples, targetCaseNum)
		if err != nil {
			t.Error(err)
		}
	}
}

var samples = `
[3,6,1,3,4]
[2,1,3,4,5]
25
3

[11,12,19]
[10,8,7]
19
0

`
