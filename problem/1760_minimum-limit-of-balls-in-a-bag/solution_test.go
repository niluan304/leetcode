package main

import (
	"testing"

	"github.com/niluan304/leetcode/tests"
)

func Test_minimum_limit_of_balls_in_a_bag(t *testing.T) {
	targetCaseNum := 0 // -1

	fs := []interface{}{
		minimumSize,
	}

	for _, f := range fs {
		err := tests.RunFunc(t, f, samples, targetCaseNum)
		if err != nil {
			t.Error(err)
		}
	}
}

var samples = `
[9]
2
3

[2,4,8,2]
4
2

`
