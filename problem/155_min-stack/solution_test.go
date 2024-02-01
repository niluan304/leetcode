package main

import (
	"testing"

	"github.com/niluan304/leetcode/tests"
)

func Test_min_stack(t *testing.T) {
	targetCaseNum := 0 // -1

	fs := []interface{}{
		Constructor,
	}

	for _, f := range fs {
		err := tests.RunClass(t, f, samples, targetCaseNum)
		if err != nil {
			t.Error(err)
		}
	}
}

var samples = `
["MinStack","push","push","push","getMin","pop","top","getMin"]
[[],[-2],[0],[-3],[],[],[],[]]
[null,null,null,null,-3,null,0,-2]

`
