package main

import (
	"testing"

	"github.com/niluan304/leetcode/tests"
)

func Test_reverse(t *testing.T) {
	targetCaseNum := 0 // -1

	fs := []func(x int) int{
		reverse,
		// reverse2,
		// leetcode,
		// endlessCheng,
	}

	for _, f := range fs {
		err := tests.RunFunc(t, f, samples, targetCaseNum)
		if err != nil {
			t.Error(err)
		}
	}
}

var samples = `
123
321

-123
-321

120
21

1534236469
0

`
