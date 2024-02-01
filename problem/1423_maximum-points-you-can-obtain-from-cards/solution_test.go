package main

import (
	"testing"

	"github.com/niluan304/leetcode/tests"
)

func Test_main(t *testing.T) {
	targetCaseNum := 0 // -1

	fs := []func(cardPoints []int, k int) int{
		maxScore,
		maxScore2,
	}

	for _, f := range fs {
		err := tests.RunFunc(t, f, samples, targetCaseNum)
		if err != nil {
			t.Error(err)
		}
	}
}

var samples = `
[1,2,3,4,5,6,1]
3
12

[2,2,2]
2
4

[9,7,7,9,7,7,9]
7
55

`
