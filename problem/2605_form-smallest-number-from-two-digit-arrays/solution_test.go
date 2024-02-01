package main

import (
	"testing"

	"github.com/niluan304/leetcode/tests"
)

func Test_form_smallest_number_from_two_digit_arrays(t *testing.T) {
	targetCaseNum := 0 // -1

	fs := []interface{}{
		minNumber,
	}

	for _, f := range fs {
		err := tests.RunFunc(t, f, samples, targetCaseNum)
		if err != nil {
			t.Error(err)
		}
	}
}

var samples = `
[4,1,3]
[5,7]
15

[3,5,2,6]
[3,1,7]
3

`
