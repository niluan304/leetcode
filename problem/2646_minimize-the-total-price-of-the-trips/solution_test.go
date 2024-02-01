package main

import (
	"testing"

	"github.com/niluan304/leetcode/tests"
)

func Test_minimize_the_total_price_of_the_trips(t *testing.T) {
	targetCaseNum := 0 // -1

	fs := []interface{}{
		minimumTotalPrice,
	}

	for _, f := range fs {
		err := tests.RunFunc(t, f, samples, targetCaseNum)
		if err != nil {
			t.Error(err)
		}
	}
}

var samples = `
4
[[0,1],[1,2],[1,3]]
[2,2,10,6]
[[0,3],[2,1],[2,3]]
23

2
[[0,1]]
[2,2]
[[0,0]]
1

`
