package main

import (
	"testing"

	"github.com/niluan304/leetcode/tests"
)

func Test_course_schedule_iii(t *testing.T) {
	targetCaseNum := 0 // -1

	fs := []func(courses [][]int) int{
		scheduleCourse,
		scheduleCourse2,
		scheduleCourse3,
		scheduleCourse4,
	}

	for _, f := range fs {
		err := tests.RunFunc(t, f, samples, targetCaseNum)
		if err != nil {
			t.Error(err)
		}
	}
}

var samples = `
[[100,200],[200,1300],[1000,1250],[2000,3200]]
3

[[1,2]]
1

[[3,2],[4,3]]
0

[[7,16],[2,3],[3,12],[3,14],[10,19],[10,16],[6,8],[6,11],[3,13],[6,16]]
4

`
