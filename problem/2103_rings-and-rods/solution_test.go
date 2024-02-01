package main

import (
	"testing"

	"github.com/niluan304/leetcode/tests"
)

func Test_rings_and_rods(t *testing.T) {
	targetCaseNum := 0 // -1

	fs := []interface{}{
		countPoints,
	}

	for _, f := range fs {
		err := tests.RunFunc(t, f, samples, targetCaseNum)
		if err != nil {
			t.Error(err)
		}
	}
}

var samples = `
"B0B6G0R6R0R6G9"
1

"B0R0G0R9R0B0G0"
1

"G4"
0

`
