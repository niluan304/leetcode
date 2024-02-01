package main

import (
	"testing"

	"github.com/niluan304/leetcode/tests"
)

func Test_main(t *testing.T) {
	targetCaseNum := 0 // -1

	fs := []func(prices []int, money int) int{
		buyChoco,
		buyChoco2,
	}

	for _, f := range fs {
		err := tests.RunFunc(t, f, samples, targetCaseNum)
		if err != nil {
			t.Error(err)
		}
	}
}

var samples = `
[1,2,2]
3
0

[3,2,3]
3
3

[98,54,6,34,66,63,52,39]
62
22

[69,91,78,19,40,13]
94
62

`
