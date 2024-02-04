package main

import (
	"testing"

	"github.com/niluan304/leetcode/tests"
)

func Test_canWinNim(t *testing.T) {
	targetCaseNum := 0 // -1

	fs := []func(n int) bool{
		//bruteForce,
		canWinNim,
		//canWinNim2,
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
4
false

1
true

2
true


`
