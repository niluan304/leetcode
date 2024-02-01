package main

import (
	"testing"

	"github.com/niluan304/leetcode/tests"
)

func Test_4sum(t *testing.T) {
	targetCaseNum := 0 // -1

	fs := []interface{}{
		fourSum,
	}

	for _, f := range fs {
		err := tests.RunFunc(t, f, samples, targetCaseNum)
		if err != nil {
			t.Error(err)
		}
	}
}

var samples = `
[1,0,-1,0,-2,2]
0
[[-2,-1,1,2],[-2,0,0,2],[-1,0,0,1]]

[2,2,2,2,2]
8
[[2,2,2,2]]

`
