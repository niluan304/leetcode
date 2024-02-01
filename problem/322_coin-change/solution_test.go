package main

import (
	"testing"

	"github.com/niluan304/leetcode/tests"
)

func Test_coin_change(t *testing.T) {
	targetCaseNum := 0 // -1

	fs := []interface{}{
		coinChange,
	}

	for _, f := range fs {
		err := tests.RunFunc(t, f, samples, targetCaseNum)
		if err != nil {
			t.Error(err)
		}
	}
}

var samples = `
[1,2,5]
11
3

[2]
3
-1

[1]
0
0

`
