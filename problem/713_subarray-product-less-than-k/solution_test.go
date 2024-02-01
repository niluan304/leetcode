package main

import (
	"testing"

	"github.com/niluan304/leetcode/tests"
)

func Test_subarray_product_less_than_k(t *testing.T) {
	targetCaseNum := 0 // -1

	fs := []interface{}{
		numSubarrayProductLessThanK,
	}

	for _, f := range fs {
		err := tests.RunFunc(t, f, samples, targetCaseNum)
		if err != nil {
			t.Error(err)
		}
	}
}

var samples = `
[10,5,2,6]
100
8

[1,2,3]
0
0

`
