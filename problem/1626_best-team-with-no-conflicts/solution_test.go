package main

import (
	"testing"

	"github.com/niluan304/leetcode/tests"
)

func Test_bestTeamScore(t *testing.T) {
	targetCaseNum := 0 // -1

	fs := []func(scores []int, ages []int) int{
		//bruteForce,
		bestTeamScore,
		//bestTeamScore2,
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
[1,3,5,10,15]
[1,2,3,4,5]
34

[4,5,6,5]
[2,1,2,1]
16

[1,2,3,5]
[8,9,10,1]
6


`
