package main

import (
	"testing"

	"github.com/niluan304/leetcode/tests"
)

func Test_couples_holding_hands(t *testing.T) {
	targetCaseNum := 0 // -1

	fs := []interface{}{
		minSwapsCouples,
	}

	for _, f := range fs {
		err := tests.RunFunc(t, f, samples, targetCaseNum)
		if err != nil {
			t.Error(err)
		}
	}
}

var samples = `
[0,2,1,3]
1

[3,2,0,1]
0

`
