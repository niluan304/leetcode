package main

import (
	"testing"

	"github.com/niluan304/leetcode/tests"
)

func Test_main(t *testing.T) {
	targetCaseNum := 0 // -1

	fs := []func(nums []int, lower int, upper int) int64{
		countFairPairs,
	}

	for _, f := range fs {
		err := tests.RunFunc(t, f, samples, targetCaseNum)
		if err != nil {
			t.Error(err)
		}
	}
}

var samples = `
[0,1,7,4,4,5]
3
6
6

[1,7,9,2,5]
11
11
1

[0,0,0,0,0,0]
0
0
15

`
