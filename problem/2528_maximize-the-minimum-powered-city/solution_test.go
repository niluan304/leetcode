package main

import (
	"testing"

	"github.com/niluan304/leetcode/tests"
)

func Test_maxPower(t *testing.T) {
	targetCaseNum := 0 // -1

	fs := []func(stations []int, r int, k int) int64{
		maxPower,
		maxPower2,
		// leetcode,
		// endlessCheng,
	}

	for _, f := range fs {
		err := tests.RunFunc(t, f, samples, targetCaseNum)
		if err != nil {
			t.Error(err)
		}
	}
}

var samples = `
[1,2,4,5,0]
1
2
5

[4,4,4,4]
0
3
4


`
