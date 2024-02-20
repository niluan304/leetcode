package main

import (
	"testing"

	. "github.com/EndlessCheng/codeforces-go/leetcode/testutil"

	"github.com/niluan304/leetcode/tests"
)

func Test_buildTree(t *testing.T) {
	targetCaseNum := 0 // -1

	fs := []func(preorder []int, inorder []int) *TreeNode{
		// bruteForce,
		buildTree,
		buildTree2,
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
[3,9,20,15,7]
[9,3,15,20,7]
[3,9,20,null,null,15,7]

[-1]
[-1]
[-1]

[3,9,5,8,7,20,15,10]
[8,5,9,7,3,15,20,10]
[3,9,20,5,7,15,10,8]

`
