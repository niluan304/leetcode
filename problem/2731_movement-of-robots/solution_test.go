package main

import (
	"testing"

	"github.com/niluan304/leetcode/tests"
)

func Test_movement_of_robots(t *testing.T) {
	targetCaseNum := 0 // -1

	fs := []interface{}{
		sumDistance,
	}

	for _, f := range fs {
		err := tests.RunFunc(t, f, samples, targetCaseNum)
		if err != nil {
			t.Error(err)
		}
	}
}

var samples = `
[-2,0,2]
"RLL"
3
8

[1,0]
"RL"
2
5

`
