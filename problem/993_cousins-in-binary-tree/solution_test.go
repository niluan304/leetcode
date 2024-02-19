package main

import (
	"testing"

	. "github.com/EndlessCheng/codeforces-go/leetcode/testutil"

	"github.com/niluan304/leetcode/tests"
)

func Test_isCousins(t *testing.T) {
	targetCaseNum := 0 // -1

	fs := []func(root *TreeNode, x int, y int) bool{
		//bruteForce,
		isCousins,
		isCousins2,
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
[1,2,3,4]
4
3
false

[1,2,3,null,4,null,5]
5
4
true

[1,2,3,null,4]
2
3
false


`
