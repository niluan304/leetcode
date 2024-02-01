package main

import (
	"testing"

	"github.com/niluan304/leetcode/tests"
)

func Test_number_of_dice_rolls_with_target_sum(t *testing.T) {
	targetCaseNum := 0 // -1

	fs := []interface{}{
		numRollsToTarget,
	}

	for _, f := range fs {
		err := tests.RunFunc(t, f, samples, targetCaseNum)
		if err != nil {
			t.Error(err)
		}
	}
}

var samples = `
1
6
3
1

2
6
7
6

30
30
500
222616187

`
