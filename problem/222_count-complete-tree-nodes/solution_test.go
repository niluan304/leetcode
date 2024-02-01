package main

import (
	"testing"

	. "github.com/EndlessCheng/codeforces-go/leetcode/testutil"

	"github.com/niluan304/leetcode/tests"
)

func Test_count_complete_tree_nodes(t *testing.T) {
	targetCaseNum := 0 // -1

	fs := []func(root *TreeNode) int{
		countNodes,
		countNodes2,
		leetcode1,
	}

	for _, f := range fs {
		err := tests.RunFunc(t, f, samples, targetCaseNum)
		if err != nil {
			t.Error(err)
		}
	}
}

var samples = `
[1,2,3,4,5,6]
6

[]
0

[1]
1

`
