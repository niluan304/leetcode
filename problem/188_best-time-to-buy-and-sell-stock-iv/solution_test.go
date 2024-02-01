package main

import (
	"testing"

	"github.com/niluan304/leetcode/tests"
)

func Test_best_time_to_buy_and_sell_stock_iv(t *testing.T) {
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
2
[2,4,1]
2

2
[3,2,6,5,0,3]
7

`
