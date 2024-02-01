package main

import (
	"testing"

	"github.com/niluan304/leetcode/tests"
)

func Test_successful_pairs_of_spells_and_potions(t *testing.T) {
	targetCaseNum := 0 // -1

	fs := []interface{}{
		successfulPairs,
	}

	for _, f := range fs {
		err := tests.RunFunc(t, f, samples, targetCaseNum)
		if err != nil {
			t.Error(err)
		}
	}
}

var samples = `
[5,1,3]
[1,2,3,4,5]
7
[4,0,3]

[3,1,2]
[8,5,8]
16
[2,0,2]

`
