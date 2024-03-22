package main

import (
	"testing"

	. "github.com/EndlessCheng/codeforces-go/leetcode/testutil"

	"github.com/niluan304/leetcode/tests"
)

func Test_minDepth(t *testing.T) {
	targetCaseNum := 0 // -1

	fs := []func(root *TreeNode) int{
		// bruteForce,
		minDepth,
		// minDepth2,
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
[3,9,20,null,null,15,7]
2

[2,null,3,null,4,null,5,null,6]
5


`
