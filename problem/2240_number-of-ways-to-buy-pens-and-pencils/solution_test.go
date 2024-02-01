package main

import (
	"testing"

	"github.com/niluan304/leetcode/tests"
)

func Test_number_of_ways_to_buy_pens_and_pencils(t *testing.T) {
	targetCaseNum := 0 // -1

	fs := []interface{}{
		waysToBuyPensPencils,
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
10
5
9

5
10
10
1

`
