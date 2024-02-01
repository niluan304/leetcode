package main

import (
	"testing"

	. "github.com/EndlessCheng/codeforces-go/leetcode/testutil"

	"github.com/niluan304/leetcode/tests"
)

func Test_main(t *testing.T) {
	targetCaseNum := 0 // -1

	fs := []func(root *TreeNode) *TreeNode{
		reverseOddLevels,
		reverseOddLevels2,
		leetcode2,
	}

	for _, f := range fs {
		err := tests.RunFunc(t, f, samples, targetCaseNum)
		if err != nil {
			t.Error(err)
		}
	}
}

var samples = `
[2,3,5,8,13,21,34]
[2,5,3,8,13,21,34]

[7,13,11]
[7,11,13]

[0,1,2,0,0,0,0,1,1,1,1,2,2,2,2]
[0,2,1,0,0,0,0,2,2,2,2,1,1,1,1]

[]
[]

`
