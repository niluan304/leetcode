package main

import (
	"testing"

	"github.com/niluan304/leetcode/tests"
)

func Test_queens_that_can_attack_the_king(t *testing.T) {
	targetCaseNum := 0 // -1

	fs := []interface{}{
		queensAttacktheKing,
	}

	for _, f := range fs {
		err := tests.RunFunc(t, f, samples, targetCaseNum)
		if err != nil {
			t.Error(err)
		}
	}
}

var samples = `
[[0,1],[1,0],[4,0],[0,4],[3,3],[2,4]]
[0,0]
[[0,1],[1,0],[3,3]]

[[0,0],[1,1],[2,2],[3,4],[3,5],[4,4],[4,5]]
[3,3]
[[2,2],[3,4],[4,4]]

`
