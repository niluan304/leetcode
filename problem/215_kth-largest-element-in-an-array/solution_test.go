package main

import (
	"testing"

	"github.com/niluan304/leetcode/tests"
)

func Test_kth_largest_element_in_an_array(t *testing.T) {
	targetCaseNum := 0 // -1

	fs := []interface{}{
		findKthLargest,
	}

	for _, f := range fs {
		err := tests.RunFunc(t, f, samples, targetCaseNum)
		if err != nil {
			t.Error(err)
		}
	}
}

var samples = `
[3,2,1,5,6,4]
2
5

[3,2,3,1,2,4,5,5,6]
4
4

`
