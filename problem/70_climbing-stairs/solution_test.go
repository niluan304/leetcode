package main

import (
	"testing"

	"github.com/niluan304/leetcode/tests"
)

func Test_climbing_stairs(t *testing.T) {
	targetCaseNum := 0 // -1

	fs := []interface{}{
		climbStairs,
		climbStairs2,
		climbStairs3,
		climbStairs4,
	}

	for _, f := range fs {
		err := tests.RunFunc(t, f, samples, targetCaseNum)
		if err != nil {
			t.Error(err)
		}
	}
}

var samples = `
2
2

3
3

`
