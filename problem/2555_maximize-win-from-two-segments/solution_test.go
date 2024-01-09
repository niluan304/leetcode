package main

import (
	"testing"

	"github.com/niluan304/leetcode/tests"
)

func Test_main(t *testing.T) {
	targetCaseNum := 0 // -1

	fs := []func(prizePositions []int, k int) int{
		maximizeWin,
		maximizeWin2,
		endlessCheng,
	}

	for _, f := range fs {
		err := tests.RunFunc(t, f, samples, targetCaseNum)
		if err != nil {
			t.Error(err)
		}
	}
}

var samples = `
[1,1,2,2,3,3,5]
2
7

[1,2,3,4]
0
2

[11,9999999,99999999,999999999]
19999999
3


`
