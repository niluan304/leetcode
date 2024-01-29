package main

import (
	"testing"

	"github.com/niluan304/leetcode/tests"
)

func Test_maxScore(t *testing.T) {
	targetCaseNum := 0 // -1

	fs := []func(nums []int, x int) int64{
		//bruteForce,
		maxScore,
		maxScore2,
		maxScore3,
		maxScore4,
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
[2,3,6,1,9,2]
5
13

[2,4,6,8]
3
20

[8,50,65,85,8,73,55,50,29,95,5,68,52,79]
74
470

`
