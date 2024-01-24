package main

import (
	"testing"

	"github.com/niluan304/leetcode/tests"
)

func Test_constructProductMatrix(t *testing.T) {
	targetCaseNum := 0 // -1

	fs := []func(grid [][]int) [][]int{
		//bruteForce,
		constructProductMatrix,
		//constructProductMatrix2,
		//leetcode,
		//endlessCheng,
	}

	for _, f := range fs {
		err := tests.RunFunc(t, f, samples, targetCaseNum)
		if err != nil {
			t.Error(err)
		}
	}
}

var samples = `
[[1,2],[3,4]]
[[24,12],[8,6]]

[[12345],[2],[1]]
[[2],[0],[0]]


`
