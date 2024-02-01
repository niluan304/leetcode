package main

import (
	"testing"

	"github.com/niluan304/leetcode/tests"
)

func Test_find_first_and_last_position_of_element_in_sorted_array(t *testing.T) {
	targetCaseNum := 0 // -1

	fs := []interface{}{
		searchRange,
	}

	for _, f := range fs {
		err := tests.RunFunc(t, f, samples, targetCaseNum)
		if err != nil {
			t.Error(err)
		}
	}
}

var samples = `
[5,7,7,8,8,10]
8
[3,4]

[5,7,7,8,8,10]
6
[-1,-1]

[]
0
[-1,-1]

`
