package main

import (
	"testing"

	"github.com/niluan304/leetcode/tests"
)

func Test_valid_triangle_number(t *testing.T) {
	targetCaseNum := 0 // -1

	fs := []interface{}{
		triangleNumber,
	}

	for _, f := range fs {
		err := tests.RunFunc(t, f, samples, targetCaseNum)
		if err != nil {
			t.Error(err)
		}
	}
}

var samples = `
[2,2,3,4]
3

[4,2,3,4]
4

`
