package main

import (
	"testing"

	"github.com/niluan304/leetcode/tests"
)

func Test_distinctDifferenceArray(t *testing.T) {
	targetCaseNum := 0 // -1

	fs := []func(nums []int) []int{
		//bruteForce,
		distinctDifferenceArray,
		//distinctDifferenceArray2,
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
[1,2,3,4,5]
[-3,-1,1,3,5]

[3,2,3,4,2]
[-2,-1,0,2,3]


`
