package main

import (
	"testing"

	"github.com/niluan304/leetcode/tests"
)

func Test_discuss(t *testing.T) {
	targetCaseNum := 0 // -1

	fs := []func(n int, cost []int) int{
		discuss,
	}

	for _, f := range fs {
		err := tests.RunFunc(t, f, discussSamples, targetCaseNum)
		if err != nil {
			t.Error(err)
		}
	}
}

var discussSamples = `
7
[1,5,0,2,4,3,4]
6

3
[5,3,3]
0

7
[1,5,2,2,3,3,1]
5

`
