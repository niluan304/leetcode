package main

import (
	"testing"

	"github.com/niluan304/leetcode/tests"
)

func Test_k_items_with_the_maximum_sum(t *testing.T) {
	targetCaseNum := 0 // -1

	fs := []interface{}{
		kItemsWithMaximumSum,
	}

	for _, f := range fs {
		err := tests.RunFunc(t, f, samples, targetCaseNum)
		if err != nil {
			t.Error(err)
		}
	}
}

var samples = `
3
2
0
2
2

3
2
0
4
3

`
