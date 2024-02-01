package main

import (
	"testing"

	"github.com/niluan304/leetcode/tests"
)

func Test_main(t *testing.T) {
	targetCaseNum := 0 // -1

	fs := []func(hours []int) int{
		longestWPI,
		longestWPI2,
		longestWPI3,
	}

	for _, f := range fs {
		err := tests.RunFunc(t, f, samples, targetCaseNum)
		if err != nil {
			t.Error(err)
		}
	}
}

var samples = `
[9,9,6,0,6,6,9]
3

[6,6,6]
0

[9,6,9,0,9,6,9,0]
7

[9,9,6,0,6,9,9]
7

[6,9,9]
3

[6,6,6,6,6,6,9]
1

[6,6,6,6,6,6,9,6]
1

[9,9,9,9,9,6,9,9,9,9]
10

[6,6,9]
1
`
