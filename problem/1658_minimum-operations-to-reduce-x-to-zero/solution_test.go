package main

import (
	"testing"

	"github.com/niluan304/leetcode/tests"
)

func Test_minimum_operations_to_reduce_x_to_zero(t *testing.T) {
	targetCaseNum := 0 // -1

	fs := []interface{}{
		minOperations,
	}

	for _, f := range fs {
		err := tests.RunFunc(t, f, samples, targetCaseNum)
		if err != nil {
			t.Error(err)
		}
	}
}

var samples = `
[1,1,4,2,3]
5
2

[5,6,7,8,9]
4
-1

[3,2,20,1,1,3]
10
5

`
