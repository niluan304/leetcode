package main

import (
	"testing"

	"github.com/niluan304/leetcode/tests"
)

func Test_minimum_time_difference(t *testing.T) {
	targetCaseNum := 0 // -1

	fs := []interface{}{
		findMinDifference,
	}

	for _, f := range fs {
		err := tests.RunFunc(t, f, samples, targetCaseNum)
		if err != nil {
			t.Error(err)
		}
	}
}

var samples = `
["23:59","00:00"]
1

["00:00","23:59","00:00"]
0

`
