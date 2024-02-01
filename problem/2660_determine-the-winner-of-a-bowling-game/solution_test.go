package main

import (
	"testing"

	"github.com/niluan304/leetcode/tests"
)

func Test_main(t *testing.T) {
	targetCaseNum := 0 // -1

	fs := []func(player1 []int, player2 []int) int{
		isWinner,
	}

	for _, f := range fs {
		err := tests.RunFunc(t, f, samples, targetCaseNum)
		if err != nil {
			t.Error(err)
		}
	}
}

var samples = `
[4,10,7,9]
[6,5,2,3]
1

[3,5,7,6]
[8,10,10,2]
2

[2,3]
[4,1]
0

`
