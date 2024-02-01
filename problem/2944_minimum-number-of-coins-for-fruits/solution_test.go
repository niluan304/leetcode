package main

import (
	"testing"

	"github.com/niluan304/leetcode/tests"
)

func Test_main(t *testing.T) {
	targetCaseNum := 0 // -1

	fs := []func(prices []int) int{
		minimumCoins,
		minimumCoins2,
		minimumCoins3,
		minimumCoins4,
		minimumCoins5,
	}

	for _, f := range fs {
		err := tests.RunFunc(t, f, samples, targetCaseNum)
		if err != nil {
			t.Error(err)
		}
	}
}

var samples = `
[3,1,2]
4

[1,10,1,1]
2

`
