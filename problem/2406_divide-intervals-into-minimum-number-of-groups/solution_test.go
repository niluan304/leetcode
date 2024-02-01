package main

import (
	"testing"

	"github.com/niluan304/leetcode/tests"
)

func Test_divide_intervals_into_minimum_number_of_groups(t *testing.T) {
	targetCaseNum := 0 // -1

	fs := []interface{}{
		minGroups,
	}

	for _, f := range fs {
		err := tests.RunFunc(t, f, samples, targetCaseNum)
		if err != nil {
			t.Error(err)
		}
	}
}

var samples = `
[[5,10],[6,8],[1,5],[2,3],[1,10]]
3

[[1,3],[5,6],[8,10],[11,13]]
1

`
