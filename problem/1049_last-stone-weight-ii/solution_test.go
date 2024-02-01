package main

import (
	"testing"

	"github.com/niluan304/leetcode/tests"
)

func Test_last_stone_weight_ii(t *testing.T) {
	targetCaseNum := 0 // -1

	fs := []interface{}{
		lastStoneWeightII,
		lastStoneWeightII2,
		lastStoneWeightII3,
	}

	for _, f := range fs {
		err := tests.RunFunc(t, f, samples, targetCaseNum)
		if err != nil {
			t.Error(err)
		}
	}
}

var samples = `
[2,7,4,1,8,1]
1

[31,26,33,21,40]
5


`
