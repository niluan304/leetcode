package main

import (
	"testing"

	"github.com/niluan304/leetcode/tests"
)

func Test_delete_operation_for_two_strings(t *testing.T) {
	targetCaseNum := 0 // -1

	fs := []interface{}{
		minDistance,
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
2

"leetcode"
"etco"
4

`
