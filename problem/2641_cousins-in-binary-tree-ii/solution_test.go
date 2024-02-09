package main

import (
	"testing"

	. "github.com/EndlessCheng/codeforces-go/leetcode/testutil"

	"github.com/niluan304/leetcode/tests"
)

func Test_replaceValueInTree(t *testing.T) {
	targetCaseNum := 0 // -1

	fs := []func(root *TreeNode) *TreeNode{
		//bruteForce,
		replaceValueInTree,
		//replaceValueInTree2,
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
[5,4,9,1,10,null,7]
[0,0,0,7,7,null,11]

[3,1,2]
[0,0,0]


`
