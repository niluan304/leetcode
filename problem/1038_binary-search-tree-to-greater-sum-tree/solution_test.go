package main

import (
	"testing"

	. "github.com/EndlessCheng/codeforces-go/leetcode/testutil"

	"github.com/niluan304/leetcode/tests"
)

func Test_main(t *testing.T) {
	targetCaseNum := 0 // -1

	fs := []func(root *TreeNode) *TreeNode{
		bstToGst,
		bstToGst2,
		bstToGst3,
	}

	for _, f := range fs {
		err := tests.RunFunc(t, f, samples, targetCaseNum)
		if err != nil {
			t.Error(err)
		}
	}
}

var samples = `
[4,1,6,0,2,5,7,null,null,null,3,null,null,null,8]
[30,36,21,36,35,26,15,null,null,null,33,null,null,null,8]

[0,null,1]
[1,null,1]

`
