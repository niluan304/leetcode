package main

import (
	"testing"

	"github.com/niluan304/leetcode/tests"
)

func Test_squares_of_a_sorted_array(t *testing.T) {
	targetCaseNum := 0 // -1

	fs := []interface{}{
		sortedSquares,
	}

	for _, f := range fs {
		err := tests.RunFunc(t, f, samples, targetCaseNum)
		if err != nil {
			t.Error(err)
		}
	}
}

var samples = `
[-4,-1,0,3,10]
[0,1,9,16,100]

[-7,-3,2,3,11]
[4,9,9,49,121]

`
