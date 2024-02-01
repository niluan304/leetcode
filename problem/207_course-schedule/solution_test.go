package main

import (
	"testing"

	"github.com/niluan304/leetcode/tests"
)

func Test_course_schedule(t *testing.T) {
	targetCaseNum := 0 // -1

	fs := []interface{}{
		canFinish,
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
[[1,0]]
true

2
[[1,0],[0,1]]
false

`
