package main

import (
	"testing"

	. "github.com/EndlessCheng/codeforces-go/leetcode/testutil"

	"github.com/niluan304/leetcode/tests"
)

func Test_constructFromPrePost(t *testing.T) {
	targetCaseNum := 0 // -1

	fs := []func(preorder []int, postorder []int) *TreeNode{
		// bruteForce,
		constructFromPrePost,
		constructFromPrePost2,
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
[3,9,5,8,7,20,15,10]
[8,5,7,9,15,10,20,3]
[3,9,20,5,7,15,10,null,8]

[1,2,4,5,3,6,7]
[4,5,2,6,7,3,1]
[1,2,3,4,5,6,7]

[1]
[1]
[1]

[2,1]
[1,2]
[2,null,1]


`
