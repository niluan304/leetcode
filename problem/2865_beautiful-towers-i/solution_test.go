package main

import (
	"testing"

	"github.com/niluan304/leetcode/tests"
)

func Test_maximumSumOfHeights(t *testing.T) {
	targetCaseNum := 0 // -1

	fs := []func(maxHeights []int) int64{
		maximumSumOfHeights,
		//maximumSumOfHeights2,
		//leetcode,
		//endlessCheng,
	}

	for _, f := range fs {
		err := tests.RunFunc(t, f, samples, targetCaseNum)
		if err != nil {
			t.Error(err)
		}
	}
}

var samples = `
[5,3,4,1,1]
13

[6,5,3,9,2,7]
22

[3,2,5,5,2,3]
18


`
