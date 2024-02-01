package main

import (
	"testing"

	"github.com/niluan304/leetcode/tests"
)

func Test_3sum(t *testing.T) {
	targetCaseNum := 0 // -1

	fs := []interface{}{
		threeSum,
	}

	for _, f := range fs {
		err := tests.RunFunc(t, f, samples, targetCaseNum)
		if err != nil {
			t.Error(err)
		}
	}
}

var samples = `
[-1,0,1,2,-1,-4]
[[-1,-1,2],[-1,0,1]]

[0,1,1]
[]

[0,0,0]
[[0,0,0]]

`
