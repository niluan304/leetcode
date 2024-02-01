package main

import (
	"testing"

	"github.com/niluan304/leetcode/tests"
)

func Test_maximum_area_of_a_piece_of_cake_after_horizontal_and_vertical_cuts(t *testing.T) {
	targetCaseNum := 0 // -1

	fs := []interface{}{
		maxArea,
	}

	for _, f := range fs {
		err := tests.RunFunc(t, f, samples, targetCaseNum)
		if err != nil {
			t.Error(err)
		}
	}
}

var samples = `
5
4
[1,2,4]
[1,3]
4

5
4
[3,1]
[1]
6

5
4
[3]
[3]
9

`
