package main

import (
	"testing"

	"github.com/niluan304/leetcode/tests"
)

func Test_difference_between_maximum_and_minimum_price_sum(t *testing.T) {
	targetCaseNum := 0 // -1

	fs := []interface{}{
		maxOutput,
	}

	for _, f := range fs {
		err := tests.RunFunc(t, f, samples, targetCaseNum)
		if err != nil {
			t.Error(err)
		}
	}
}

var samples = `
6
[[0,1],[1,2],[1,3],[3,4],[3,5]]
[9,8,7,6,10,5]
24

3
[[0,1],[1,2]]
[1,1,1]
2

`
