package main

import (
	"testing"

	"github.com/niluan304/leetcode/tests"
)

func Test_distribute_money_to_maximum_children(t *testing.T) {
	targetCaseNum := 0 // -1

	fs := []interface{}{
		distMoney,
	}

	for _, f := range fs {
		err := tests.RunFunc(t, f, samples, targetCaseNum)
		if err != nil {
			t.Error(err)
		}
	}
}

var samples = `
20
3
1

16
2
2

`
