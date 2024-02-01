package main

import (
	"testing"

	"github.com/niluan304/leetcode/tests"
)

func Test_maximum_enemy_forts_that_can_be_captured(t *testing.T) {
	targetCaseNum := 0 // -1

	fs := []interface{}{
		captureForts,
	}

	for _, f := range fs {
		err := tests.RunFunc(t, f, samples, targetCaseNum)
		if err != nil {
			t.Error(err)
		}
	}
}

var samples = `
[1,0,0,-1,0,0,0,0,1]
4

[0,0,1,-1]
0

`
