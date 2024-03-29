package main

import (
	"testing"

	"github.com/niluan304/leetcode/tests"
)

func Test_solveNQueens(t *testing.T) {
	targetCaseNum := 0 // -1

	fs := []func(n int) [][]string{
		//bruteForce,
		solveNQueens,
		solveNQueens2,
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
4
[[".Q..","...Q","Q...","..Q."],["..Q.","Q...","...Q",".Q.."]]

1
[["Q"]]


`
