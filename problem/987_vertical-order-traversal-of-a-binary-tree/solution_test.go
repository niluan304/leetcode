package main

import (
	"testing"

	. "github.com/EndlessCheng/codeforces-go/leetcode/testutil"

	"github.com/niluan304/leetcode/tests"
)

func Test_verticalTraversal(t *testing.T) {
	targetCaseNum := 0 // -1

	fs := []func(root *TreeNode) [][]int{
		//bruteForce,
		verticalTraversal,
		verticalTraversal2,
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
[3,9,20,null,null,15,7]
[[9],[3,15],[20],[7]]

[1,2,3,4,5,6,7]
[[4],[2],[1,5,6],[3],[7]]

[1,2,3,4,6,5,7]
[[4],[2],[1,5,6],[3],[7]]


`
