package main

import (
	"testing"

	"github.com/niluan304/leetcode/tests"
)

func Test_take_gifts_from_the_richest_pile(t *testing.T) {
	targetCaseNum := 0 // -1

	fs := []func(gifts []int, k int) int64{
		pickGifts,
	}

	for _, f := range fs {
		err := tests.RunFunc(t, f, samples, targetCaseNum)
		if err != nil {
			t.Error(err)
		}
	}
}

var samples = `
[25,64,9,4,100]
4
29

[1,1,1,1]
4
4

`
