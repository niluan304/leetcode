package main

import (
	"testing"

	"github.com/niluan304/leetcode/tests"
)

func Test_maximizeSquareHoleArea(t *testing.T) {
	targetCaseNum := 0 // -1

	fs := []func(n int, m int, hBars []int, vBars []int) int{
		maximizeSquareHoleArea,
		// maximizeSquareHoleArea2,
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
2
1
[2,3]
[2]
4

1
1
[2]
[2]
4

2
3
[2,3]
[2,3,4]
9

4
40
[5,3,2,4]
[36,41,6,34,33]
9

`
