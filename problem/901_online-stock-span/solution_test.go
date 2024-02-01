package main

import (
	"testing"

	"github.com/niluan304/leetcode/tests"
)

func Test_online_stock_span(t *testing.T) {
	targetCaseNum := 0 // -1

	fs := []interface{}{
		Constructor,
		Constructor2,
		Constructor3,
		Constructor4,
	}

	for _, f := range fs {
		err := tests.RunClass(t, f, samples, targetCaseNum)
		if err != nil {
			t.Error(err)
		}
	}
}

var samples = `
["StockSpanner", "next", "next", "next", "next", "next", "next", "next"]
[[], [100], [80], [60], [70], [60], [75], [85]]
[null, 1, 1, 1, 2, 1, 4, 6]

`
