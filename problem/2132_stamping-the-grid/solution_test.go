package main

import (
	"testing"

	"github.com/niluan304/leetcode/tests"
)

func Test_main(t *testing.T) {
	targetCaseNum := 0 // -1

	fs := []func(grid [][]int, stampHeight int, stampWidth int) bool{
		possibleToStamp,
	}

	for _, f := range fs {
		err := tests.RunFunc(t, f, samples, targetCaseNum)
		if err != nil {
			t.Error(err)
		}
	}
}

var samples = `
[[1,0,0,0],[1,0,0,0],[1,0,0,0],[1,0,0,0],[1,0,0,0]]
4
3
true

[[1,0,0,0],[0,1,0,0],[0,0,1,0],[0,0,0,1]]
2
2
false

`
