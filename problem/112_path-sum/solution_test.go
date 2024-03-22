package main

import (
	"testing"

	. "github.com/EndlessCheng/codeforces-go/leetcode/testutil"

	"github.com/niluan304/leetcode/tests"
)

func Test_hasPathSum(t *testing.T) {
	targetCaseNum := 0 // -1

	fs := []func(root *TreeNode, targetSum int) bool{
		// bruteForce,
		hasPathSum,
		// hasPathSum2,
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
[5,4,8,11,null,13,4,7,2,null,null,null,1]
22
true

[1,2,3]
5
false

[]
0
false


`
