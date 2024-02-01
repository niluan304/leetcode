package main

import (
	"testing"

	"github.com/niluan304/leetcode/tests"
)

func Test_edit_distance(t *testing.T) {
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
"horse"
"ros"
3

"intention"
"execution"
5

`
