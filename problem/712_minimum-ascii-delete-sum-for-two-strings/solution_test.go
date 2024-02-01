package main

import (
	"testing"

	"github.com/niluan304/leetcode/tests"
)

func Test_minimum_ascii_delete_sum_for_two_strings(t *testing.T) {
	targetCaseNum := 0 // -1

	fs := []interface{}{
		minimumDeleteSum,
	}

	for _, f := range fs {
		err := tests.RunFunc(t, f, samples, targetCaseNum)
		if err != nil {
			t.Error(err)
		}
	}
}

var samples = `
"sea"
"eat"
231

"delete"
"leet"
403

`
