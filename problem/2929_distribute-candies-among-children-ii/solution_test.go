package main

import (
	"testing"

	"github.com/niluan304/leetcode/tests"
)

func Test_distributeCandies(t *testing.T) {
	targetCaseNum := 0 // -1

	fs := []func(n int, limit int) int64{
		//bruteForce,
		distributeCandies,
		distributeCandies2,
		distributeCandies3,
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
2
3

3
3
10

99
3
0

`
