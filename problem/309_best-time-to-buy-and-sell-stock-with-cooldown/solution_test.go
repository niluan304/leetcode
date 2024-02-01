package main

import (
	"testing"

	"github.com/niluan304/leetcode/tests"
)

func Test_best_time_to_buy_and_sell_stock_with_cooldown(t *testing.T) {
	targetCaseNum := 0 // -1

	fs := []interface{}{
		maxProfit,
	}

	for _, f := range fs {
		err := tests.RunFunc(t, f, samples, targetCaseNum)
		if err != nil {
			t.Error(err)
		}
	}
}

var samples = `
[1,2,3,0,2]
3

[1]
0

`
