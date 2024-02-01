package main

import (
	"testing"

	"github.com/niluan304/leetcode/tests"
)

func Test_find_the_longest_valid_obstacle_course_at_each_position(t *testing.T) {
	targetCaseNum := 0 // -1

	fs := []interface{}{
		longestObstacleCourseAtEachPosition,
	}

	for _, f := range fs {
		err := tests.RunFunc(t, f, samples, targetCaseNum)
		if err != nil {
			t.Error(err)
		}
	}
}

var samples = `
[1,2,3,2]
[1,2,3,3]

[2,2,1]
[1,2,1]

[3,1,5,6,4,2]
[1,1,2,3,2,2]

`
