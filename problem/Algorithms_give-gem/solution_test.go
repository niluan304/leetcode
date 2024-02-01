package main

import (
	"testing"

	"github.com/niluan304/leetcode/tests"
)

func Test_give_gem(t *testing.T) {
	targetCaseNum := 0 // -1

	fs := []interface{}{
		giveGem,
	}

	for _, f := range fs {
		err := tests.RunFunc(t, f, samples, targetCaseNum)
		if err != nil {
			t.Error(err)
		}
	}
}

var samples = `
[3,1,2]
[[0,2],[2,1],[2,0]]
2

[100,0,50,100]
[[0,2],[0,1],[3,0],[3,0]]
75

[0,0,0,0]
[[1,2],[3,1],[1,2]]
0
`
