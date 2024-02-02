package main

import (
	"testing"

	"github.com/niluan304/leetcode/tests"
)

func Test_maximizeTheProfit(t *testing.T) {
	targetCaseNum := 0 // -1

	fs := []func(n int, offers [][]int) int{
		//bruteForce,
		maximizeTheProfit,
		maximizeTheProfit2,
		maximizeTheProfit3,
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
5
[[0,0,1],[0,2,2],[1,3,2]]
3

5
[[0,0,1],[0,2,10],[1,3,2]]
10


`
