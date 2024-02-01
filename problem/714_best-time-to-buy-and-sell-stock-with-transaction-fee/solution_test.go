package main

import (
	"testing"

	"github.com/niluan304/leetcode/tests"
)

func Test_best_time_to_buy_and_sell_stock_with_transaction_fee(t *testing.T) {
	targetCaseNum := 0 // -1

	fs := []interface{}{
		maxProfit,
		maxProfit2,
		maxProfit3,
	}

	for _, f := range fs {
		err := tests.RunFunc(t, f, samples, targetCaseNum)
		if err != nil {
			t.Error(err)
		}
	}
}

var samples = `
[1,3,2,8,4,9]
2
8

[1,3,7,5,10,3]
3
6

`
