package main

import (
	"testing"

	"github.com/niluan304/leetcode/tests"
)

func Test_course_schedule_ii(t *testing.T) {
	targetCaseNum := 0 // -1

	fs := []func(numCourses int, prerequisites [][]int) []int{
		findOrder,
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
[0,1]

4
[[1,0],[2,0],[3,1],[3,2]]
[0,2,1,3]

1
[]
[0]

`
