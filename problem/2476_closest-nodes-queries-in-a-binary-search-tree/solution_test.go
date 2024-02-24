package main

import (
	"testing"

	. "github.com/EndlessCheng/codeforces-go/leetcode/testutil"

	"github.com/niluan304/leetcode/tests"
)

func Test_closestNodes(t *testing.T) {
	targetCaseNum := 0 // -1

	fs := []func(root *TreeNode, queries []int) [][]int{
		//bruteForce,
		closestNodes,
		closestNodes2,
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
[6,2,13,1,4,9,15,null,null,null,null,null,null,14]
[2,5,16]
[[2,2],[4,6],[15,-1]]

[4,null,9]
[3]
[[-1,4]]


`
