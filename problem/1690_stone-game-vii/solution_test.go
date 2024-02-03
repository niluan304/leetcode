package main

import (
	"testing"

	"github.com/niluan304/leetcode/tests"
)

func Test_stoneGameVII(t *testing.T) {
	targetCaseNum := 0 // -1

	fs := []func(stones []int) int{
		//bruteForce,
		stoneGameVII,
		stoneGameVII2,
		stoneGameVII3,
		stoneGameVII4,
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
[5,3,1,4,2]
6

[7,90,5,1,100,10,10,2]
122


`
