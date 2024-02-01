package main

import (
	"testing"

	"github.com/niluan304/leetcode/tests"
)

func Test_median_of_two_sorted_arrays(t *testing.T) {
	targetCaseNum := 0 // -1

	fs := []interface{}{
		findMedianSortedArrays,
	}

	for _, f := range fs {
		err := tests.RunFunc(t, f, samples, targetCaseNum)
		if err != nil {
			t.Error(err)
		}
	}
}

var samples = `
[1,3]
[2]
2.00000

[1,2]
[3,4]
2.50000

`
