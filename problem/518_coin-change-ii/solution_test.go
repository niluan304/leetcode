package main

import (
	"testing"

	"github.com/niluan304/leetcode/tests"
)

func Test_coin_change_ii(t *testing.T) {
	targetCaseNum := 0 // -1

	fs := []func(amount int, coins []int) int{
		//bruteForce,
		change,
		change2,
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
[1,2,5]
4

3
[2]
0

10
[10]
1

100
[1,2,5]
541

`
