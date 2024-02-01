package main

import (
	"testing"

	"github.com/niluan304/leetcode/tests"
)

func Test_partition_equal_subset_sum(t *testing.T) {
	targetCaseNum := 0 // -1

	fs := []interface{}{
		canPartition,
	}

	for _, f := range fs {
		err := tests.RunFunc(t, f, samples, targetCaseNum)
		if err != nil {
			t.Error(err)
		}
	}
}

var samples = `
[1,5,11,5]
true

[1,2,3,5]
false

`
