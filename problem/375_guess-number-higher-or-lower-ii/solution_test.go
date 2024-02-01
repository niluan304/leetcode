package main

import (
	"testing"

	"github.com/niluan304/leetcode/tests"
)

func Test_guess_number_higher_or_lower_ii(t *testing.T) {
	targetCaseNum := 0 // -1

	fs := []interface{}{
		getMoneyAmount,
	}

	for _, f := range fs {
		err := tests.RunFunc(t, f, samples, targetCaseNum)
		if err != nil {
			t.Error(err)
		}
	}
}

var samples = `
10
16

1
0

2
1

`
