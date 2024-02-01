package main

import (
	"testing"

	. "github.com/EndlessCheng/codeforces-go/leetcode/testutil"

	"github.com/niluan304/leetcode/tests"
)

func Test_populating_next_right_pointers_in_each_node_ii(t *testing.T) {
	targetCaseNum := 0 // -1

	fs := []func(root *Node) *Node{
		connect,
		connect2,
	}

	for _, f := range fs {
		err := tests.RunFunc(t, f, samples, targetCaseNum)
		if err != nil {
			t.Error(err)
		}
	}
}

var samples = `
[1,2,3,4,5,null,7]
[1,#,2,3,#,4,5,7,#]

[]
[]

`
