package main

import (
	"testing"

	. "github.com/EndlessCheng/codeforces-go/leetcode/testutil"

	"github.com/niluan304/leetcode/tests"
)

func Test_rangeSumBST(t *testing.T) {
	targetCaseNum := 0 // -1

	fs := []func(root *TreeNode, low int, high int) int{
		//bruteForce,
		rangeSumBST,
		rangeSumBST2,
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
[10,5,15,3,7,null,18]
7
15
32

[10,5,15,3,7,13,18,1,null,6]
6
10
23


`
