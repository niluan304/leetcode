package main

import (
	"testing"

	"github.com/niluan304/leetcode/tests"
)

func Test_sellingWood(t *testing.T) {
	targetCaseNum := 0 // -1

	fs := []func(m int, n int, prices [][]int) int64{
		// bruteForce,
		sellingWood,
		sellingWood2,
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
3
5
[[1,4,2],[2,2,7],[2,1,3]]
19

4
6
[[3,2,10],[1,4,2],[4,1,3]]
32


`
