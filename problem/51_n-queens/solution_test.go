package main

import (
	"testing"

	"github.com/niluan304/leetcode/tests"
)

func Test_n_queens(t *testing.T) {
	targetCaseNum := 0 // -1

	fs := []interface{}{
		solveNQueens,
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
