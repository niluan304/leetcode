package main

import (
	"testing"

	"github.com/niluan304/leetcode/tests"
)

func Test_max_dot_product_of_two_subsequences(t *testing.T) {
	targetCaseNum := 0 // -1

	fs := []interface{}{
		maxDotProduct,
	}

	for _, f := range fs {
		err := tests.RunFunc(t, f, samples, targetCaseNum)
		if err != nil {
			t.Error(err)
		}
	}
}

var samples = `
[2,1,-2,5]
[3,0,-6]
18

[3,-2]
[2,-6,7]
21

[-1,-1]
[1,1]
-1

`
