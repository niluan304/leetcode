package main

import (
	"testing"

	"github.com/niluan304/leetcode/tests"
)

func Test_minimum_cost_to_merge_stones(t *testing.T) {
	targetCaseNum := 0 // -1

	fs := []interface{}{
		mergeStones,
	}

	for _, f := range fs {
		err := tests.RunFunc(t, f, samples, targetCaseNum)
		if err != nil {
			t.Error(err)
		}
	}
}

var samples = `
[3,2,4,1]
2
20

[3,2,4,1]
3
-1

[3,5,1,2,6]
3
25

`
