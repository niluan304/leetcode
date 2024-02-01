package main

import (
	"testing"

	"github.com/niluan304/leetcode/tests"
)

func Test_check_knight_tour_configuration(t *testing.T) {
	targetCaseNum := 0 // -1

	fs := []interface{}{
		checkValidGrid,
	}

	for _, f := range fs {
		err := tests.RunFunc(t, f, samples, targetCaseNum)
		if err != nil {
			t.Error(err)
		}
	}
}

var samples = `
[[0,11,16,5,20],[17,4,19,10,15],[12,1,8,21,6],[3,18,23,14,9],[24,13,2,7,22]]
true

[[0,3,6],[5,8,1],[2,7,4]]
false

`
