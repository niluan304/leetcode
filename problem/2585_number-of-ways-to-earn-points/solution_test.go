package main

import (
	"testing"

	"github.com/niluan304/leetcode/tests"
)

func Test_waysToReachTarget(t *testing.T) {
	targetCaseNum := 0 // -1

	fs := []func(target int, types [][]int) int{
		//bruteForce,
		waysToReachTarget,
		waysToReachTarget2,
		waysToReachTarget3,
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
6
[[6,1],[3,2],[2,3]]
7

5
[[50,1],[50,2],[50,5]]
4

18
[[6,1],[3,2],[2,3]]
1


`
