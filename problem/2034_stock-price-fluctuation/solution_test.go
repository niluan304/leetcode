package main

import (
	"testing"

	"github.com/niluan304/leetcode/tests"
)

func Test_stock_price_fluctuation(t *testing.T) {
	targetCaseNum := 0 // -1

	fs := []interface{}{
		Constructor,
		Constructor2,
	}

	for _, f := range fs {
		err := tests.RunClass(t, f, samples, targetCaseNum)
		if err != nil {
			t.Error(err)
		}
	}
}

var samples = `
["StockPrice", "update", "update", "current", "maximum", "update", "maximum", "update", "minimum"]
[[], [1, 10], [2, 5], [], [], [1, 3], [], [4, 2], []]
[null, null, null, 5, 10, null, 5, null, 2]

`
