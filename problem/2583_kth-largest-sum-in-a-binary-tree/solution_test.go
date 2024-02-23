package main

import (
	"testing"

	. "github.com/EndlessCheng/codeforces-go/leetcode/testutil"

	"github.com/niluan304/leetcode/tests"
)

func Test_kthLargestLevelSum(t *testing.T) {
	targetCaseNum := 0 // -1

	fs := []func(root *TreeNode, k int) int64{
		//bruteForce,
		kthLargestLevelSum,
		kthLargestLevelSum2,
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
[5,8,9,2,1,3,7,4,6]
2
13

[1,2,null,3]
1
3


`
