package main

import (
	"testing"

	"github.com/niluan304/leetcode/tests"
)

func Test_last_moment_before_all_ants_fall_out_of_a_plank(t *testing.T) {
	targetCaseNum := 0 // -1

	fs := []interface{}{
		getLastMoment,
	}

	for _, f := range fs {
		err := tests.RunFunc(t, f, samples, targetCaseNum)
		if err != nil {
			t.Error(err)
		}
	}
}

var samples = `
4
[4,3]
[0,1]
4

7
[]
[0,1,2,3,4,5,6,7]
7

7
[0,1,2,3,4,5,6,7]
[]
7

`
