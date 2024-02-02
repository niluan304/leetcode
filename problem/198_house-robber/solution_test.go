package main

import (
	"testing"

	"github.com/niluan304/leetcode/tests"
)

func Test_house_robber(t *testing.T) {
	targetCaseNum := 0 // -1

	fs := []func(nums []int) int{
		rob,
		rob2,
		rob3,
		rob4,
	}

	for _, f := range fs {
		err := tests.RunFunc(t, f, samples, targetCaseNum)
		if err != nil {
			t.Error(err)
		}
	}
}

var samples = `
[1,2,3,1]
4

[2,7,9,3,1]
12

`
