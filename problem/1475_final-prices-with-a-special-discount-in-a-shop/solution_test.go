package main

import (
	"testing"

	"github.com/niluan304/leetcode/tests"
)

func Test_main(t *testing.T) {
	targetCaseNum := 0 // -1

	fs := []interface{}{
		finalPrices,
		finalPrices2,
	}

	for _, f := range fs {
		err := tests.RunFunc(t, f, samples, targetCaseNum)
		if err != nil {
			t.Error(err)
		}
	}
}

var samples = `
[8,4,6,2,3]
[4,2,4,2,3]

[1,2,3,4,5]
[1,2,3,4,5]

[10,1,1,6]
[9,0,1,6]

`
