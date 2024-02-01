package main

import (
	"testing"

	"github.com/niluan304/leetcode/tests"
)

func Test_main(t *testing.T) {
	targetCaseNum := 0 // -1

	fs := []func(arr []int, mat [][]int) int{
		firstCompleteIndex,
	}

	for _, f := range fs {
		err := tests.RunFunc(t, f, samples, targetCaseNum)
		if err != nil {
			t.Error(err)
		}
	}
}

var samples = `
[1,3,4,2]
[[1,4],[2,3]]
2

[2,8,7,4,1,3,5,6,9]
[[3,2,5],[1,4,6],[8,7,9]]
3

[1,4,5,2,6,3]
[[4,3,5],[1,2,6]]
1
`
