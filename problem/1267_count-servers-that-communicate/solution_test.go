package main

import (
	"testing"

	"github.com/niluan304/leetcode/tests"
)

func Test_countServers(t *testing.T) {
	targetCaseNum := 0 // -1

	fs := []func(grid [][]int) int{
		// bruteForce,
		countServers,
		// countServers2,
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
[[1,0],[0,1]]
0

[[1,0],[1,1]]
3

[[1,1,0,0],[0,0,1,0],[0,0,1,0],[0,0,0,1]]
4

[[1,0,0,1,0],[0,0,0,0,0],[0,0,0,1,0]]
3


`
